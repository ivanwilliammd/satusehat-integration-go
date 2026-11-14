package parser_test

// ClaimParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestClaimParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "Claim.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "Claim" {
		t.Fatalf("resourceType = %v, want Claim", data["resourceType"])
	}

	pr := parser.NewClaimParser(data)
	if pr.GetRaw() == nil {
		t.Fatal("GetRaw() nil")
	}
	if !reflect.DeepEqual(pr.GetResourceType(), data["resourceType"]) {
		t.Fatalf("GetResourceType() = %v, want %v", pr.GetResourceType(), data["resourceType"])
	}
	if !reflect.DeepEqual(pr.GetIdentifier(), data["identifier"]) {
		t.Fatalf("GetIdentifier() = %v, want %v", pr.GetIdentifier(), data["identifier"])
	}
	if !reflect.DeepEqual(pr.GetStatus(), data["status"]) {
		t.Fatalf("GetStatus() = %v, want %v", pr.GetStatus(), data["status"])
	}
	if !reflect.DeepEqual(pr.GetType(), data["type"]) {
		t.Fatalf("GetType() = %v, want %v", pr.GetType(), data["type"])
	}
	if !reflect.DeepEqual(pr.GetUse(), data["use"]) {
		t.Fatalf("GetUse() = %v, want %v", pr.GetUse(), data["use"])
	}
	if !reflect.DeepEqual(pr.GetPatient(), data["patient"]) {
		t.Fatalf("GetPatient() = %v, want %v", pr.GetPatient(), data["patient"])
	}
	if !reflect.DeepEqual(pr.GetBillablePeriod(), data["billablePeriod"]) {
		t.Fatalf("GetBillablePeriod() = %v, want %v", pr.GetBillablePeriod(), data["billablePeriod"])
	}
	if !reflect.DeepEqual(pr.GetCreated(), data["created"]) {
		t.Fatalf("GetCreated() = %v, want %v", pr.GetCreated(), data["created"])
	}
	if !reflect.DeepEqual(pr.GetEnterer(), data["enterer"]) {
		t.Fatalf("GetEnterer() = %v, want %v", pr.GetEnterer(), data["enterer"])
	}
	if !reflect.DeepEqual(pr.GetInsurer(), data["insurer"]) {
		t.Fatalf("GetInsurer() = %v, want %v", pr.GetInsurer(), data["insurer"])
	}
	if !reflect.DeepEqual(pr.GetProvider(), data["provider"]) {
		t.Fatalf("GetProvider() = %v, want %v", pr.GetProvider(), data["provider"])
	}
	if !reflect.DeepEqual(pr.GetPriority(), data["priority"]) {
		t.Fatalf("GetPriority() = %v, want %v", pr.GetPriority(), data["priority"])
	}
	if !reflect.DeepEqual(pr.GetPayee(), data["payee"]) {
		t.Fatalf("GetPayee() = %v, want %v", pr.GetPayee(), data["payee"])
	}
	if !reflect.DeepEqual(pr.GetFacility(), data["facility"]) {
		t.Fatalf("GetFacility() = %v, want %v", pr.GetFacility(), data["facility"])
	}
	if !reflect.DeepEqual(pr.GetSupportingInfo(), data["supportingInfo"]) {
		t.Fatalf("GetSupportingInfo() = %v, want %v", pr.GetSupportingInfo(), data["supportingInfo"])
	}
	if !reflect.DeepEqual(pr.GetDiagnosis(), data["diagnosis"]) {
		t.Fatalf("GetDiagnosis() = %v, want %v", pr.GetDiagnosis(), data["diagnosis"])
	}
	if !reflect.DeepEqual(pr.GetProcedure(), data["procedure"]) {
		t.Fatalf("GetProcedure() = %v, want %v", pr.GetProcedure(), data["procedure"])
	}
	if !reflect.DeepEqual(pr.GetInsurance(), data["insurance"]) {
		t.Fatalf("GetInsurance() = %v, want %v", pr.GetInsurance(), data["insurance"])
	}
	if !reflect.DeepEqual(pr.GetTotal(), data["total"]) {
		t.Fatalf("GetTotal() = %v, want %v", pr.GetTotal(), data["total"])
	}
	if !reflect.DeepEqual(pr.GetExtension(), data["extension"]) {
		t.Fatalf("GetExtension() = %v, want %v", pr.GetExtension(), data["extension"])
	}
}
