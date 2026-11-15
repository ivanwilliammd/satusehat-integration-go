package builder

// Kontrak lintas-SDK: jalankan golden/builder/*.input.json terhadap builder Go.
//
// Kasus dan hasil harapan tinggal di satusehat-conformance, bukan di sini — supaya lima SDK
// diuji terhadap satu kontrak, bukan terhadap asumsinya masing-masing.
//
// Spec memakai camelCase (nama method kanon PHP). Go memakai PascalCase, jadi nama method
// diubah huruf awalnya saja; argumen scalar diberikan apa adanya lewat reflection.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func dirKonformansi() string {
	if v := os.Getenv("CONFORMANCE_DIR"); v != "" {
		return v
	}
	return "/home/openclaw/projects/satusehat-conformance"
}

// registri membatasi resource yang sudah punya kasus kontrak.
var registri = map[string]func() interface{}{
	"Patient": func() interface{} { return NewPatientBuilder() },
}

type panggilanSpek struct {
	Method string            `json:"method"`
	Args   []json.RawMessage `json:"args"`
}

type kasusSpek struct {
	Resource string          `json:"resource"`
	Calls    []panggilanSpek `json:"calls"`
}

func TestKontrakKonformansi(t *testing.T) {
	dir := filepath.Join(dirKonformansi(), "golden", "builder")
	berkas, err := filepath.Glob(filepath.Join(dir, "*.input.json"))
	if err != nil || len(berkas) == 0 {
		t.Skipf("kasus kontrak tidak ada di %s", dir)
	}

	for _, path := range berkas {
		nama := strings.TrimSuffix(filepath.Base(path), ".input.json")
		t.Run(nama, func(t *testing.T) {
			isi, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("baca %s: %v", path, err)
			}
			var spec kasusSpek
			if err := json.Unmarshal(isi, &spec); err != nil {
				t.Fatalf("parse %s: %v", path, err)
			}

			buat, ada := registri[spec.Resource]
			if !ada {
				t.Skipf("resource %s belum terdaftar di registri uji kontrak", spec.Resource)
			}
			builder := buat()

			for _, panggilan := range spec.Calls {
				namaMethod := strings.ToUpper(panggilan.Method[:1]) + panggilan.Method[1:]
				m := reflect.ValueOf(builder).MethodByName(namaMethod)
				if !m.IsValid() {
					t.Fatalf("method %s tidak ada di builder %s", namaMethod, spec.Resource)
				}
				args := make([]reflect.Value, 0, len(panggilan.Args))
				for _, raw := range panggilan.Args {
					args = append(args, nilaiArg(t, m.Type().In(len(args)), raw))
				}
				m.Call(args)
			}

			hasil := builder.(interface{ Build() map[string]interface{} }).Build()
			harapanRaw, err := os.ReadFile(filepath.Join(dir, nama+".json"))
			if err != nil {
				t.Fatalf("baca golden %s.json: %v", nama, err)
			}
			var harapan map[string]interface{}
			if err := json.Unmarshal(harapanRaw, &harapan); err != nil {
				t.Fatalf("parse golden %s.json: %v", nama, err)
			}

			a, _ := json.Marshal(normalisasi(hasil))
			b, _ := json.Marshal(normalisasi(harapan))
			if string(a) != string(b) {
				t.Errorf("payload berbeda dari kontrak\nGo:      %s\nKontrak: %s", a, b)
			}
		})
	}
}

// nilaiArg mengubah JSON mentah menjadi nilai dengan tipe parameter yang diminta.
func nilaiArg(t *testing.T, tipe reflect.Type, raw json.RawMessage) reflect.Value {
	t.Helper()
	switch tipe.Kind() {
	case reflect.String:
		var s string
		_ = json.Unmarshal(raw, &s)
		return reflect.ValueOf(s).Convert(tipe)
	case reflect.Bool:
		var b bool
		_ = json.Unmarshal(raw, &b)
		return reflect.ValueOf(b).Convert(tipe)
	case reflect.Int, reflect.Int64:
		var n int64
		_ = json.Unmarshal(raw, &n)
		return reflect.ValueOf(n).Convert(tipe)
	case reflect.Float64:
		var f float64
		_ = json.Unmarshal(raw, &f)
		return reflect.ValueOf(f).Convert(tipe)
	case reflect.Ptr:
		// Materialiser datatype: struct datatype punya json tag, jadi JSON dari spec bisa
		// di-unmarshal langsung ke tipe parameter sebenarnya. Kunci "$type" diabaikan.
		if tipe.Elem().Kind() == reflect.Struct {
			ptr := reflect.New(tipe.Elem())
			if err := json.Unmarshal(raw, ptr.Interface()); err != nil {
				t.Fatalf("materialise %s: %v", tipe, err)
			}
			return ptr
		}
		t.Fatalf("tipe argumen %s (pointer non-struct) belum didukung uji kontrak", tipe)
		return reflect.Value{}
	default:
		t.Fatalf("tipe argumen %s belum didukung uji kontrak (butuh materialiser datatype)", tipe)
		return reflect.Value{}
	}
}

// normalisasi membuang null dan memaksa angka menjadi float64 agar setara dengan JSON.
func normalisasi(v interface{}) interface{} {
	switch x := v.(type) {
	case map[string]interface{}:
		hasil := map[string]interface{}{}
		for k, isi := range x {
			if isi == nil {
				continue
			}
			hasil[k] = normalisasi(isi)
		}
		return hasil
	case []interface{}:
		out := make([]interface{}, 0, len(x))
		for _, isi := range x {
			out = append(out, normalisasi(isi))
		}
		return out
	case int:
		return float64(x)
	case int64:
		return float64(x)
	default:
		return v
	}
}
