package builder_test

// CarePlanBuilder must produce the same payload as PayloadBuilderCarePlan in the PHP SDK, including
// the activity wrapper ({detail: …}).

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestCarePlanMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "care_plan_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	b := builder.NewCarePlanBuilder()
	b.SetId("cp-1")
	b.AddIdentifier(&datatype.Identifier{System: "http://sys-ids.kemkes.go.id/careplan/1", Value: "cp-1"})
	b.SetStatus("active")
	b.SetIntent("plan")
	b.AddCategory(concept("http://snomed.info/sct", "38717000", "Ambulatory"))
	b.SetTitle("Rencana Perawatan Hipertensi")
	b.SetDescription("Pemantauan tekanan darah dan diet rendah garam")
	b.SetSubject(&datatype.Reference{Reference: "Patient/1000000001", Display: "Budi Santoso"})
	b.SetEncounter(&datatype.Reference{Reference: "Encounter/2000000001"})
	b.SetPeriod(&datatype.Period{Start: str("2026-01-05"), End: str("2026-04-05")})
	b.SetAuthor(&datatype.Reference{Reference: "Practitioner/1000000002", Display: "dr. Ani"})
	b.AddContributor(&datatype.Reference{Reference: "Practitioner/1000000003"})
	b.AddAddressee(&datatype.Reference{Reference: "Patient/1000000001"})
	b.AddGoal(&datatype.Reference{Reference: "Goal/g-1"})
	b.AddActivity(map[string]interface{}{
		"status":      "scheduled",
		"description": "Kontrol tekanan darah tiap 2 minggu",
		"code":        concept("http://snomed.info/sct", "386053000", "Evaluation procedure"),
	})

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
