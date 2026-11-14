package builder_test

// MedicationAdministrationBuilder must produce the same payload as
// PayloadBuilderMedicationAdministration in the PHP SDK. The reference setters accept a string or a
// Reference there, so both forms are exercised here.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestMedicationAdministrationMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "medication_administration_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	b := builder.NewMedicationAdministrationBuilder()
	b.SetId("9100000001")
	b.AddIdentifier(&datatype.Identifier{System: "http://sys-ids.kemkes.go.id/medicationadministration/1000000001", Value: "9100000001"})
	b.SetStatus("completed")
	b.SetMedicationReference("Medication/6000000001")                    // string form
	b.SetSubject("Patient/1000000001")                                   // string form
	b.SetContext(&datatype.Reference{Reference: "Encounter/2000000001"}) // Reference form
	b.SetEffectivePeriod(&datatype.Period{
		Start: ptr("2026-01-05T08:00:00+07:00"),
		End:   ptr("2026-01-05T08:10:00+07:00"),
	})
	b.AddPerformer(&datatype.Reference{Reference: "Practitioner/1000000002", Display: "Ns. Dwi"})
	b.AddReasonCode(concept("http://hl7.org/fhir/sid/icd-10", "A09", "Diarrhoea"))
	b.SetRequest(&datatype.Reference{Reference: "MedicationRequest/7000000001"})
	b.SetDosage(
		qty(500, "mg", "http://unitsofmeasure.org", "mg"),
		concept("http://terminology.hl7.org/CodeSystem/v3-RouteOfAdministration", "PO", "Oral"),
		concept("http://snomed.info/sct", "181251004", "Structure of mouth"),
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
