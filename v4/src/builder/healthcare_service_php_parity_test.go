package builder_test

// HealthcareServiceBuilder must produce the same payload as PayloadBuilderHealthcareService in the
// PHP SDK for every setter PHP can actually run.
//
// addIdentifier is included: the PHP method used to call $this->add(), which Builder does not define,
// but that was fixed in the canonical SDK (88 builders affected), so identifier parity is now real.

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
	b.AddIdentifier("http://sys-ids.kemkes.go.id/healthcareservice/1", "1", "official", "PRN", "")
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
