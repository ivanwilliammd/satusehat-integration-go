package builder_test

// MedicationRequestBuilder must produce the same payload as PayloadBuilderMedicationRequest in the
// PHP SDK, including the nested dispenseRequest element built by three separate setters.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestMedicationRequestMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "medication_request_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	allowed := true
	tablet := func() *datatype.Quantity { return &datatype.Quantity{Value: ptr(10.0), Unit: str("tablet")} }
	b := builder.NewMedicationRequestBuilder()
	b.SetId("7000000001")
	b.AddIdentifier(&datatype.Identifier{System: "http://sys-ids.kemkes.go.id/medicationrequest/1000000001", Value: "7000000001"})
	b.SetStatus("active")
	b.SetIntent("order")
	b.AddCategory(concept("http://terminology.hl7.org/CodeSystem/medicationrequest-category", "outpatient", "Outpatient"))
	b.SetMedicationCodeableConcept(concept("http://terminology.kemkes.go.id/CodeSystem/medication-code", "910002021", "Paracetamol 500 mg"))
	b.SetSubject(&datatype.Reference{Reference: "Patient/1000000001", Display: "Budi Santoso"})
	b.SetEncounter(&datatype.Reference{Reference: "Encounter/2000000001"})
	b.SetAuthoredOn("2026-01-05T09:10:00+07:00")
	b.SetRequester(&datatype.Reference{Reference: "Practitioner/1000000002", Display: "dr. Ani"})
	b.AddReasonCode(concept("http://hl7.org/fhir/sid/icd-10", "A09", "Diarrhoea"))
	b.AddDosageInstruction(&datatype.Dosage{Text: str("3 x 1 tablet sesudah makan")})
	dur := &datatype.Duration{Quantity: *qty(3, "day", "http://unitsofmeasure.org", "d")}
	b.SetDispenseRequest(nil, 1, tablet(), dur, nil)
	b.SetInitialFill(tablet())
	b.SetSubstitution(concept("http://terminology.hl7.org/CodeSystem/v3-substanceAdminSubstitution", "G", "generic"), &allowed)

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
