package builder_test

// AllergyIntoleranceBuilder must produce the same payload as PayloadBuilderAllergyIntolerance in the
// PHP SDK, including the baked-in clinical/verification systems and the nested reaction.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestAllergyIntoleranceMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "allergy_intolerance_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	b := builder.NewAllergyIntoleranceBuilder()
	b.AddIdentifier(&datatype.Identifier{System: "http://sys-ids.kemkes.go.id/allergy/1000000001", Value: "A1"})
	b.SetClinicalStatus("active")
	b.SetVerificationStatus("confirmed")
	b.SetType("allergy")
	b.AddCategory("medication")
	b.SetCriticality("high")
	b.SetCode(concept("http://snomed.info/sct", "387517004", "Paracetamol"))
	b.SetPatient(&datatype.Reference{Reference: "Patient/1000000001", Display: "Budi Santoso"})
	b.SetEncounter(&datatype.Reference{Reference: "Encounter/2000000001"})
	b.SetOnsetDateTime("2026-01-05T08:00:00+07:00")
	b.SetRecordedDate("2026-01-05T08:20:00+07:00")
	b.SetRecorder(&datatype.Reference{Reference: "Practitioner/1000000002", Display: "dr. Ani"})
	b.SetAsserter(&datatype.Reference{Reference: "Patient/1000000001"})
	b.SetLastOccurrence("2026-01-05T08:05:00+07:00")
	b.AddNote("Riwayat gatal setelah minum paracetamol.")
	b.AddReaction(
		concept("http://snomed.info/sct", "387517004", "Paracetamol"),
		concept("http://snomed.info/sct", "126485001", "Urticaria"),
		"Ruam merah pada lengan",
		"2026-01-05T08:05:00+07:00",
		"moderate",
		concept("http://terminology.hl7.org/CodeSystem/v3-RouteOfAdministration", "PO", "Oral"),
		"Muncul 15 menit setelah minum",
	)

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
