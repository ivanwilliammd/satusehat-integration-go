package builder_test

// CoverageBuilder must produce the same payload as PayloadBuilderCoverage in the PHP SDK. PHP takes
// primitives there, so the Go calls look different but the JSON must be identical.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestCoverageMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "coverage_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	b := builder.NewCoverageBuilder()
	b.SetId("cov-1")
	b.AddIdentifier(&datatype.Identifier{System: "http://sys-ids.kemkes.go.id/coverage/1000000001", Value: "cov-1"})
	b.SetStatus("active")
	b.SetType(concept("http://terminology.hl7.org/CodeSystem/v3-ActCode", "EHCPOL", "extended healthcare"))
	b.SetPolicyHolder(&datatype.Reference{Reference: "Patient/1000000001"})
	b.SetSubscriber(&datatype.Reference{Reference: "Patient/1000000001"})
	b.SetSubscriberId("BPJS-0001")
	b.SetBeneficiary(&datatype.Reference{Reference: "Patient/1000000001"})
	b.SetDependent("1")
	b.SetRelationship(&datatype.CodeableConcept{Coding: []datatype.Coding{{
		System: "http://terminology.hl7.org/CodeSystem/subscriber-relationship", Code: "self",
	}}})
	b.AddPayor(&datatype.Reference{Reference: "Organization/1000000001", Display: "BPJS Kesehatan"})
	b.AddClass(&datatype.CodeableConcept{Coding: []datatype.Coding{{
		System: "http://terminology.hl7.org/CodeSystem/coverage-class", Code: "group",
	}}}, "GRP-1", "Grup A")
	b.SetOrder(1)
	b.SetNetwork("JKN")
	b.AddCostToBeneficiary(50000, "IDR")
	b.SetPeriod("2026-01-01", "2026-12-31")

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
