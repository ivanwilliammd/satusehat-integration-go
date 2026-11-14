package builder_test

// ProcedureBuilder must produce the same payload as PayloadBuilderProcedure in the PHP SDK.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestProcedureMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "procedure_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	b := builder.NewProcedureBuilder()
	b.SetId("9300000001")
	b.AddIdentifier(&datatype.Identifier{System: "http://sys-ids.kemkes.go.id/procedure/1000000001", Value: "9300000001"})
	b.SetStatus("completed")
	b.SetCategory(concept("http://snomed.info/sct", "103693007", "Diagnostic procedure"))
	b.SetCode(concept("http://hl7.org/fhir/sid/icd-9-cm", "88.72", "Diagnostic ultrasound of heart"))
	b.SetSubject(&datatype.Reference{Reference: "Patient/1000000001", Display: "Budi Santoso"})
	b.SetEncounter(&datatype.Reference{Reference: "Encounter/2000000001"})
	b.AddBasedOn(&datatype.Reference{Reference: "ServiceRequest/9400000001"})
	b.AddReasonCode(concept("http://hl7.org/fhir/sid/icd-10", "I50.9", "Heart failure"))
	b.SetPerformedPeriod(&datatype.Period{
		Start: ptr("2026-01-05T08:00:00+07:00"),
		End:   ptr("2026-01-05T08:30:00+07:00"),
	})
	b.AddPerformer(&datatype.Reference{Reference: "Practitioner/1000000002", Display: "dr. Ani"}, nil, nil)
	b.SetOutcome(concept("http://snomed.info/sct", "385669000", "Successful"))
	b.AddReport(&datatype.Reference{Reference: "DiagnosticReport/9500000001"})
	b.AddFollowUp(concept("http://snomed.info/sct", "182836005", "Review of medication"))
	b.AddNote("Pasien puasa 6 jam sebelum tindakan.")
	b.AddFocalDevice(concept("http://snomed.info/sct", "129442004", "Insertion"), nil, &datatype.Reference{Reference: "Device/9600000001"})
	b.AddUsedCode(concept("http://snomed.info/sct", "40701008", "Echocardiography"))
	b.AddBodySite(concept("http://snomed.info/sct", "302509004", "Heart structure"))
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
