package builder

// Uji RATCHET permukaan method kanon PHP vs Go.
//
// Uji golden menjaga perilaku resource yang sudah diporting. Uji ini menjaga PERMUKAAN:
// berapa method publik di PayloadBuilder*.php yang belum punya pasangan di paket ini.
// Angkanya hanya boleh turun.
//
// Berkas PHP dibaca dari PHP_BUILDER_DIR (default: jalur repo kanon di mesin pengembang).
// Bila tidak ada — mis. CI tanpa checkout PHP — uji dilewati, bukan gagal, supaya uji ini
// tidak pernah menjadi hijau palsu: ia hanya bisa lulus bila benar-benar membaca kanon.

import (
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
	dir := dirBuilderPHP()
	if _, err := os.Stat(dir); err != nil {
		t.Skipf("berkas kanon PHP tidak ada di %s (set PHP_BUILDER_DIR)", dir)
	}

	phpFiles, err := filepath.Glob(filepath.Join(dir, "PayloadBuilder*.php"))
	if err != nil || len(phpFiles) == 0 {
		t.Skipf("tidak ada PayloadBuilder*.php di %s", dir)
	}

	peta := petaResourceGo(t)
	var kurang []string
	var tanpaBuilder []string

	for _, berkas := range phpFiles {
		isi, err := os.ReadFile(berkas)
		if err != nil {
			t.Fatalf("baca %s: %v", berkas, err)
		}
		m := reResourceType.FindSubmatch(isi)
		if m == nil {
			continue
		}
		resource := string(m[1])
		kodeGo, ada := peta[resource]
		if !ada {
			tanpaBuilder = append(tanpaBuilder, resource)
			continue
		}
		for _, mm := range reMethodPHP.FindAllSubmatch(isi, -1) {
			method := string(mm[1])
			if method == "__construct" || method == "build" {
				continue
			}
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
