package builder_test

// EndpointBuilder must produce the same payload as PayloadBuilderEndpoint in the PHP SDK, and must
// reject the same statuses PHP rejects.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestEndpointMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "endpoint_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	b := builder.NewEndpointBuilder()
	b.SetId("end-1")
	b.AddIdentifier("http://sys-ids.kemkes.go.id/endpoint/1", str("end-1"))
	b.AddIdentifier(&datatype.Identifier{System: "urn:oid:1.2.3", Value: "alt-1"}, nil)
	b.SetStatus("active")
	b.SetConnectionType("hl7-fhir-rest", "HL7 FHIR", "")
	b.SetName("Endpoint SATUSEHAT")
	b.SetManagingOrganization("1000000001", str("RS Contoh"))
	b.AddContact("phone", "0254000002", "work")
	b.SetPeriod("2026-01-01", "")
	b.AddPayloadType("any", "Any", "")
	b.AddPayloadMimeType("application/fhir+json")
	b.SetAddress("https://api.satusehat.kemkes.go.id/fhir-r4/v1")
	b.AddHeader("Authorization: Bearer x")

	if err := b.Err(); err != nil {
		t.Fatalf("tidak diharapkan error: %v", err)
	}

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

// PHP throws InvalidArgumentException for an unknown status; Go records it and writes nothing.
func TestEndpointRejectsInvalidStatus(t *testing.T) {
	b := builder.NewEndpointBuilder()
	b.SetStatus("bogus")
	if b.Err() == nil {
		t.Fatal("status tidak sah seharusnya menghasilkan error")
	}
	if _, ada := b.Build()["status"]; ada {
		t.Fatal("status tidak sah tidak boleh ditulis ke payload")
	}
}
