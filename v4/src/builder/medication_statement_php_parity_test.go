package builder_test

// MedicationStatementBuilder must produce the same payload as PayloadBuilderMedicationStatement in
// the PHP SDK, including the reference prefixing rules and the dosage element name.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestMedicationStatementMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "medication_statement_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	b := builder.NewMedicationStatementBuilder()
	b.SetId("9000000001")
	b.AddIdentifier(&datatype.Identifier{System: "http://sys-ids.kemkes.go.id/medicationstatement/1000000001", Value: "9000000001"})
	b.SetStatus("active")
	b.AddStatusReason("replaced", "Replaced", "http://terminology.hl7.org/CodeSystem/medication-statement-status")
	b.SetCategory(concept("http://terminology.hl7.org/CodeSystem/medication-statement-category", "inpatient", "Inpatient"))
	b.SetMedicationCodeableConcept(concept("http://terminology.kemkes.go.id/CodeSystem/medication-code", "910002021", "Paracetamol 500 mg"))
	b.SetSubject("1000000001", "Budi Santoso")
	b.SetContext("2000000001", "")
	b.SetDateAsserted("2026-01-05T09:00:00+07:00")
	b.SetEffectiveDateTime("2026-01-04T20:00:00+07:00")
	b.SetInformationSource("1000000002", "dr. Ani")
	b.SetDerivedFrom("MedicationRequest/7000000001")
	b.SetReasonCode(concept("http://hl7.org/fhir/sid/icd-10", "A09", "Diarrhoea"))
	b.SetReasonReference(&datatype.Reference{Reference: "Condition/5000000001"})
	b.AddNote("Dilaporkan sendiri oleh pasien.")
	b.AddDosageInstruction("3 x 1 tablet sesudah makan", 3, 1, "d")

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
