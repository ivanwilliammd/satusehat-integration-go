package builder_test

// HealthcareServiceBuilder must produce the same payload as PayloadBuilderHealthcareService in the
// PHP SDK for every setter PHP can actually run.
//
// addIdentifier is excluded from the comparison on purpose: the PHP method calls $this->add(),
// which Builder does not define, so it dies with "Call to undefined method" instead of building a
// payload. The Go implementation writes the identifier, and TestHealthcareServiceIdentifierWorks
// asserts that separately rather than pretending parity exists where the canonical code is broken.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
)

func TestHealthcareServiceMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "healthcare_service_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	b := builder.NewHealthcareServiceBuilder()
	b.SetId("hs-1")
	b.SetStatus("active")
	b.SetName("Poli Umum")
	b.SetType("http://terminology.hl7.org/CodeSystem/service-type", "124", "General Practice")
	b.SetType("http://terminology.hl7.org/CodeSystem/service-type", "125", "General Practice 2")
	b.SetLocation("Location/1000000003", str("Poli Umum Lt. 1"))
	b.SetProvidedBy("Organization/1000000001", nil)

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

func TestHealthcareServiceIdentifierWorks(t *testing.T) {
	b := builder.NewHealthcareServiceBuilder()
	b.AddIdentifier("http://sys-ids.kemkes.go.id/healthcareservice/1", "1", "official", "PRN", "")

	ids, ok := b.Build()["identifier"].([]interface{})
	if !ok || len(ids) != 1 {
		t.Fatalf("identifier tidak tertulis: %v", b.Build()["identifier"])
	}
	ident := ids[0].(map[string]interface{})
	if ident["system"] != "http://sys-ids.kemkes.go.id/healthcareservice/1" || ident["value"] != "1" {
		t.Fatalf("identifier salah: %v", ident)
	}
	if ident["use"] != "official" {
		t.Fatalf("use hilang: %v", ident)
	}
	tipe := ident["type"].(map[string]interface{})["coding"].([]interface{})[0].(map[string]interface{})
	if tipe["code"] != "PRN" || tipe["display"] != "PRN" {
		t.Fatalf("type coding salah (display harus default ke code): %v", tipe)
	}
}
