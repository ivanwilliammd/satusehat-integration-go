package builder_test

// BundleBuilder must produce the same payload as PayloadBuilderBundle in the PHP SDK, including the
// transaction entry with its conditional-create request block.
//
// The PHP constructor stamps timestamp with the current time; the test pins it with SetTimestamp so
// the comparison is about the payload shape, not about the clock.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
)

func TestBundleMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "bundle_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	b := builder.NewBundleBuilder("transaction")
	b.SetId("bundle-1")
	b.SetTimestamp("2026-01-05T09:00:00+07:00")
	b.SetTotal(2)
	b.AddTransactionEntry(
		map[string]interface{}{"resourceType": "Patient", "id": "1000000001"},
		"urn:uuid:1",
		"POST",
		"Patient",
		"", "", "identifier=http://sys-ids.kemkes.go.id/patient/1000000001|1000000001",
	)
	b.AddGetEntry("urn:uuid:2", "Patient/1000000001", "")
	b.AddLink("self", "https://api.satusehat.kemkes.go.id/fhir-r4/v1")

	if err := b.Err(); err != nil {
		t.Fatalf("tidak diharapkan error: %v", err)
	}

	mentah, err := b.BuildJSON()
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var dapat map[string]interface{}
	if err := json.Unmarshal(mentah, &dapat); err != nil {
		t.Fatalf("hasil bukan JSON sah: %v", err)
	}

	if !reflect.DeepEqual(harapan, dapat) {
		h, _ := json.MarshalIndent(harapan, "", "  ")
		d, _ := json.MarshalIndent(dapat, "", "  ")
		t.Fatalf("payload Go berbeda dari PHP\n--- PHP ---\n%s\n--- Go ---\n%s", h, d)
	}
}

// An invalid type is rejected by PHP with an InvalidArgumentException; Go records it in Err() and
// must not write the element.
func TestBundleRejectsInvalidType(t *testing.T) {
	b := builder.NewBundleBuilder("not-a-bundle-type")
	if b.Err() == nil {
		t.Fatal("tipe bundle tidak sah seharusnya menghasilkan error")
	}
	if _, ada := b.Build()["type"]; ada {
		t.Fatal("tipe tidak sah tidak boleh ditulis ke payload")
	}
}
