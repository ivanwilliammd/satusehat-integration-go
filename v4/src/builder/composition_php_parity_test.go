package builder_test

// CompositionBuilder must produce the same payload as PayloadBuilderComposition in the PHP SDK.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestCompositionMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "composition_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	b := builder.NewCompositionBuilder()
	b.SetId("comp-1")
	b.AddIdentifier(&datatype.Identifier{System: "http://sys-ids.kemkes.go.id/composition/1000000001", Value: "comp-1"})
	b.SetStatus("final")
	b.SetType(concept("http://loinc.org", "11488-4", "Consult note"))
	b.AddCategory(concept("http://terminology.hl7.org/CodeSystem/v3-ActCode", "LP173421-1", "Report"))
	b.SetTitle("Resume Medis Rawat Jalan")
	b.SetDate("2026-01-05T09:30:00+07:00")
	b.AddAuthor(&datatype.Reference{Reference: "Practitioner/1000000002", Display: "dr. Ani"})
	b.SetSubject(&datatype.Reference{Reference: "Patient/1000000001", Display: "Budi Santoso"})
	b.SetEncounter(&datatype.Reference{Reference: "Encounter/2000000001"})
	b.SetCustodian(&datatype.Reference{Reference: "Organization/1000000001", Display: "RS Contoh"})
	b.AddAttester("legal", &datatype.Reference{Reference: "Practitioner/1000000002"})
	b.AddSection(map[string]interface{}{
		"title": "Keluhan Utama",
		"code":  concept("http://loinc.org", "10154-3", "Chief complaint"),
		"text":  map[string]interface{}{"status": "generated", "div": "<div>Demam tiga hari</div>"},
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
