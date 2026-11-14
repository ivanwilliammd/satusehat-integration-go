package builder_test

// CareTeamBuilder must produce the same payload as PayloadBuilderCareTeam in the PHP SDK for every
// setter PHP can run.
//
// addIdentifier is included: the PHP method used to call $this->add(), which Builder does not define;
// that defect was fixed in the canonical SDK, so identifier parity is now real.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
)

func TestCareTeamMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "care_team_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	b := builder.NewCareTeamBuilder()
	b.SetId("ct-1")
	b.AddIdentifier("urn:oid:1.2.3", "ct-1", "official", "PRN", "")
	b.SetStatus("active")
	b.SetCode("http://snomed.info/sct", "768242006", "Care team")
	b.SetSubject("Patient/1000000001", str("Budi Santoso"))
	b.SetEncounter("Encounter/2000000001")

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
