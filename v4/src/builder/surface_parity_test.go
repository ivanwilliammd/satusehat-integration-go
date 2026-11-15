package builder

// Uji RATCHET permukaan method kanon PHP vs Go.
//
// Uji golden menjaga perilaku resource yang sudah diporting. Uji ini menjaga PERMUKAAN:
// berapa method publik di PayloadBuilder*.php yang belum punya pasangan di paket ini.
// Angkanya hanya boleh turun.
//
// Daftar method kanon dibaca dari contracts/php_method_surface.json yang ikut ter-commit,
// sehingga uji ini berjalan di CI tanpa checkout repo PHP (privat) maupun token.
// Bila checkout PHP tersedia, uji kedua memastikan kontrak itu belum basi.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

const batasPermukaan = 0

// Method kanon yang memang tidak diporting, beserta alasannya.
var diperbolehkan = map[string]bool{
	"Bundle.json":     true, // mengembalikan string JSON; Go memakai BuildJSON
	"Bundle.validate": true, // validasi ada di Validate() pada paket validate
}

func jalurKontrak() string {
	if v := os.Getenv("PHP_METHOD_CONTRACT"); v != "" {
		return v
	}
	return filepath.Join("..", "..", "contracts", "php_method_surface.json")
}

func dirBuilderPHP() string {
	if v := os.Getenv("PHP_BUILDER_DIR"); v != "" {
		return v
	}
	return "/home/openclaw/projects/satusehat-integration/src/Builder"
}

var (
	reResourceType = regexp.MustCompile(`resourceType\s*=\s*'([^']+)'`)
	reMethodPHP    = regexp.MustCompile(`public function (\w+)\(`)
	reResourceGo   = regexp.MustCompile(`ResourceType:\s*"([^"]+)"`)
	reAltResource  = regexp.MustCompile(`"resourceType":\s*"([^"]+)"`)
)

func muatKontrak(t *testing.T) map[string][]string {
	t.Helper()
	isi, err := os.ReadFile(jalurKontrak())
	if err != nil {
		t.Fatalf("baca kontrak: %v", err)
	}
	var kontrak map[string][]string
	if err := json.Unmarshal(isi, &kontrak); err != nil {
		t.Fatalf("parse kontrak: %v", err)
	}
	return kontrak
}

// petaResourceGo memetakan resourceType ke isi berkas Go yang mendefinisikannya.
func petaResourceGo(t *testing.T) map[string]string {
	berkas, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatalf("glob: %v", err)
	}
	peta := map[string]string{}
	for _, nama := range berkas {
		if strings.HasSuffix(nama, "_test.go") {
			continue
		}
		isi, err := os.ReadFile(nama)
		if err != nil {
			t.Fatalf("baca %s: %v", nama, err)
		}
		if m := reResourceGo.FindSubmatch(isi); m != nil {
			peta[string(m[1])] = string(isi)
		} else if m := reAltResource.FindSubmatch(isi); m != nil {
			peta[string(m[1])] = string(isi)
		}
	}
	return peta
}

func TestPermukaanMethodKanon(t *testing.T) {
	kontrak := muatKontrak(t)
	peta := petaResourceGo(t)

	var kurang []string
	var tanpaBuilder []string

	for resource, methods := range kontrak {
		kodeGo, ada := peta[resource]
		if !ada {
			tanpaBuilder = append(tanpaBuilder, resource)
			continue
		}
		for _, method := range methods {
			if diperbolehkan[resource+"."+method] {
				continue
			}
			// Go memakai PascalCase untuk method terekspor.
			pascal := strings.ToUpper(method[:1]) + method[1:]
			if !strings.Contains(kodeGo, ") "+pascal+"(") && !strings.Contains(kodeGo, ") "+method+"(") {
				kurang = append(kurang, resource+"."+method)
			}
		}
	}

	if len(tanpaBuilder) > 0 {
		t.Errorf("resource tanpa builder Go (%d): %s", len(tanpaBuilder), strings.Join(tanpaBuilder, ", "))
	}
	if len(kurang) > batasPermukaan {
		t.Errorf("method kanon belum diporting: %d (batas %d): %s", len(kurang), batasPermukaan, strings.Join(kurang, ", "))
	}
}

// TestKontrakTidakBasi hanya berjalan bila checkout kanon PHP tersedia.
func TestKontrakTidakBasi(t *testing.T) {
	dir := dirBuilderPHP()
	if _, err := os.Stat(dir); err != nil {
		t.Skipf("checkout kanon PHP tidak ada di %s (set PHP_BUILDER_DIR)", dir)
	}
	phpFiles, err := filepath.Glob(filepath.Join(dir, "PayloadBuilder*.php"))
	if err != nil || len(phpFiles) == 0 {
		t.Skipf("tidak ada PayloadBuilder*.php di %s", dir)
	}

	nyata := map[string][]string{}
	for _, berkas := range phpFiles {
		isi, err := os.ReadFile(berkas)
		if err != nil {
			t.Fatalf("baca %s: %v", berkas, err)
		}
		m := reResourceType.FindSubmatch(isi)
		if m == nil {
			continue
		}
		metode := []string{}
		for _, mm := range reMethodPHP.FindAllSubmatch(isi, -1) {
			nama := string(mm[1])
			if nama == "__construct" || nama == "build" {
				continue
			}
			metode = append(metode, nama)
		}
		nyata[string(m[1])] = metode
	}

	kontrak := muatKontrak(t)
	if len(nyata) != len(kontrak) {
		t.Fatalf("kontrak basi: kanon %d resource, kontrak %d resource", len(nyata), len(kontrak))
	}
	for resource, metode := range nyata {
		if len(metode) != len(kontrak[resource]) {
			t.Errorf("kontrak basi untuk %s: kanon %d method, kontrak %d", resource, len(metode), len(kontrak[resource]))
		}
	}
}
