package builder_test

// Guard against setters that accept an argument and drop it. Such a builder compiles, its arguments
// look right at the call site, and the payload silently loses the value — the failure only shows up
// when the server rejects the resource. 84 builders shipped with exactly that bug: SetSubject(reference
// string) wrote datatype.Reference{}.ToArray(), an empty reference.
//
// Read the package source instead of a hand-written list, so a new builder with the same defect fails
// this test without anyone remembering to add a case.

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestNoSetterWritesAnEmptyStructAndDropsItsArgument(t *testing.T) {
	pola := regexp.MustCompile(`b\.Data\["[^"]+"\] = datatype\.\w+\{\}\.ToArray\(\)`)
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
		t.Fatalf("setter menulis struct kosong dan membuang argumennya di: %s", strings.Join(temuan, ", "))
	}
}
