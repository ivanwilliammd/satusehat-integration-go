package builder_test

// DiagnosticReportBuilder must produce the same payload as PayloadBuilderDiagnosticReport in the
// PHP SDK, including the two-element result list.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestDiagnosticReportMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "diagnostic_report_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	b := builder.NewDiagnosticReportBuilder()
	b.SetId("9500000001")
	b.AddIdentifier(&datatype.Identifier{System: "http://sys-ids.kemkes.go.id/diagnosticreport/1000000001", Value: "9500000001"})
	b.SetStatus("final")
	b.AddCategory(concept("http://terminology.hl7.org/CodeSystem/v2-0074", "LAB", "Laboratory"))
	b.SetCode(concept("http://loinc.org", "57021-8", "CBC panel"))
	b.SetSubject(&datatype.Reference{Reference: "Patient/1000000001", Display: "Budi Santoso"})
	b.SetEncounter(&datatype.Reference{Reference: "Encounter/2000000001"})
	b.SetEffectiveDateTime("2026-01-05T08:30:00+07:00")
	b.SetIssued("2026-01-05T09:00:00+07:00")
	b.AddPerformer(&datatype.Reference{Reference: "Organization/1000000001", Display: "Lab RS Contoh"})
	b.AddResult(&datatype.Reference{Reference: "Observation/3000000001"})
	b.AddResult(&datatype.Reference{Reference: "Observation/3000000002"})
	b.AddSpecimen(&datatype.Reference{Reference: "Specimen/4000000001"})
	b.AddBasedOn(&datatype.Reference{Reference: "ServiceRequest/9400000001"})
	b.SetConclusion("Dalam batas normal")
	b.AddConclusionCode(concept("http://snomed.info/sct", "281900007", "No abnormality detected"))

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
