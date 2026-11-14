package builder_test

// A path key written straight into the payload becomes a literal JSON key: Data["meta/profile"] produces
// {"meta/profile": …}, which is valid JSON and not FHIR. Nested elements must go through setPath.

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestNoLiteralPathKeysInPayload(t *testing.T) {
	pola := regexp.MustCompile(`b\.Data\["[^"]*/[^"]*"\]`)
	berkas, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatalf("glob: %v", err)
	}

	var temuan []string
	for _, nama := range berkas {
		if strings.HasSuffix(nama, "_test.go") {
			continue
		}
		isi, err := os.ReadFile(nama)
		if err != nil {
			t.Fatalf("baca %s: %v", nama, err)
		}
		if pola.Match(isi) {
			temuan = append(temuan, nama)
		}
	}

	if len(temuan) > 0 {
		t.Fatalf("kunci jalur ditulis sebagai kunci literal di: %s", strings.Join(temuan, ", "))
	}
}
