package builder_test

// DocumentReferenceBuilder must produce the same payload as PayloadBuilderDocumentReference in the
// PHP SDK, including meta-ish context passed as a raw map and the list-shaped type element.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestDocumentReferenceMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "document_reference_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	b := builder.NewDocumentReferenceBuilder()
	b.SetId("doc-1")
	b.AddIdentifier(&datatype.Identifier{System: "http://sys-ids.kemkes.go.id/document/1000000001", Value: "doc-1"})
	b.SetMasterIdentifier(&datatype.Identifier{System: "urn:oid:1.2.3.4", Value: "master-1"})
	b.SetStatus("current")
	b.SetDocStatus("final")
	b.AddType(concept("http://loinc.org", "11488-4", "Consult note"))
	b.AddCategory(concept("http://terminology.hl7.org/CodeSystem/v3-ActCode", "LP173421-1", "Report"))
	b.SetSubject(&datatype.Reference{Reference: "Patient/1000000001", Display: "Budi Santoso"})
	b.SetDate("2026-01-05T10:00:00+07:00")
	b.AddAuthor(&datatype.Reference{Reference: "Practitioner/1000000002", Display: "dr. Ani"})
	b.SetCustodian(&datatype.Reference{Reference: "Organization/1000000001"})
	b.AddRelatesTo("replaces", &datatype.Reference{Reference: "DocumentReference/doc-0"})
	b.SetDescription("Resume medis rawat jalan")
	b.AddSecurityLabel(concept("http://terminology.hl7.org/CodeSystem/v3-Confidentiality", "N", "Normal"))
	b.AddContent(
		&datatype.Attachment{ContentType: str("application/pdf"), URL: str("https://example.org/doc-1.pdf")},
		concept("http://ihe.net/fhir/ValueSet/IHE.FormatCode.codesystem", "urn:ihe:iti:xds:2017:mimeTypeSufficient", "PDF"),
	)
	b.SetContext(map[string]interface{}{"encounter": []interface{}{map[string]interface{}{"reference": "Encounter/2000000001"}}})

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
