package builder_test

// ConditionBuilder must produce the same payload as PayloadBuilderCondition in the PHP SDK.
// testdata/condition_php_golden.json is the PHP builder's output for a fully populated Condition.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestConditionMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "condition_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	b := builder.NewConditionBuilder()
	b.SetId("5000000001")
	b.AddIdentifier(&datatype.Identifier{System: "http://sys-ids.kemkes.go.id/condition/1000000001", Value: "5000000001"})
	b.SetClinicalStatus(concept("http://terminology.hl7.org/CodeSystem/condition-clinical", "active", "Active"))
	b.SetVerificationStatus(concept("http://terminology.hl7.org/CodeSystem/condition-ver-status", "confirmed", "Confirmed"))
	b.AddCategory(concept("http://terminology.hl7.org/CodeSystem/condition-category", "encounter-diagnosis", "Encounter Diagnosis"))
	b.SetCode(concept("http://hl7.org/fhir/sid/icd-10", "A09", "Diarrhoea and gastroenteritis of presumed infectious origin"))
	b.SetSubject(&datatype.Reference{Reference: "Patient/1000000001", Display: "Budi Santoso"})
	b.SetEncounter(&datatype.Reference{Reference: "Encounter/2000000001"})
	b.SetOnsetDateTime("2026-01-04T10:00:00+07:00")
	b.SetRecordedDate("2026-01-05T08:45:00+07:00")
	b.SetRecorder(&datatype.Reference{Reference: "Practitioner/1000000002", Display: "dr. Ani"})
	b.SetAsserter(&datatype.Reference{Reference: "Practitioner/1000000002", Display: "dr. Ani"})
	b.AddStage(concept("http://snomed.info/sct", "255407002", "Stage 1"), nil)
	b.AddEvidence(concept("http://loinc.org", "2339-0", "Glucose"), nil)
	b.AddNote("Pasien membaik dengan rehidrasi.")

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
