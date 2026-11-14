package builder_test

// ServiceRequestBuilder must produce the same payload as PayloadBuilderServiceRequest in the PHP SDK.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestServiceRequestMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "service_request_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	b := builder.NewServiceRequestBuilder()
	b.SetId("9400000001")
	b.AddIdentifier(&datatype.Identifier{System: "http://sys-ids.kemkes.go.id/servicerequest/1000000001", Value: "9400000001"})
	b.SetRequisition(&datatype.Identifier{System: "http://sys-ids.kemkes.go.id/requisition/1000000001", Value: "REQ-1"})
	b.SetStatus("active")
	b.SetIntent("order")
	b.AddCategory(concept("http://snomed.info/sct", "108252007", "Laboratory procedure"))
	b.SetPriority("urgent")
	b.SetDoNotPerform(false)
	b.SetCode(concept("http://loinc.org", "57021-8", "CBC panel"))
	q := &datatype.Quantity{Value: ptr(1.0), Unit: str("panel")}
	b.SetQuantityQuantity(q)
	b.SetSubject(&datatype.Reference{Reference: "Patient/1000000001", Display: "Budi Santoso"})
	b.SetEncounter(&datatype.Reference{Reference: "Encounter/2000000001"})
	b.SetOccurrenceDateTime("2026-01-05T08:00:00+07:00")
	b.SetAuthoredOn("2026-01-05T07:40:00+07:00")
	b.SetRequester(&datatype.Reference{Reference: "Practitioner/1000000002", Display: "dr. Ani"})
	b.AddPerformer(&datatype.Reference{Reference: "Organization/1000000001", Display: "Lab RS Contoh"})
	b.AddReasonCode(concept("http://hl7.org/fhir/sid/icd-10", "A09", "Diarrhoea"))
	b.AddSpecimen(&datatype.Reference{Reference: "Specimen/4000000001"})
	b.AddNote("Mohon hasil dipercepat.")
	b.SetPatientInstruction("Puasa 8 jam sebelum pengambilan darah.")

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
