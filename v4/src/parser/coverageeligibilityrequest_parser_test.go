package parser_test

// CoverageEligibilityRequestParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestCoverageEligibilityRequestParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "CoverageEligibilityRequest.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "CoverageEligibilityRequest" {
		t.Fatalf("resourceType = %v, want CoverageEligibilityRequest", data["resourceType"])
	}

	pr := parser.NewCoverageEligibilityRequestParser(data)
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
	if !reflect.DeepEqual(pr.GetPriority(), data["priority"]) {
		t.Fatalf("GetPriority() = %v, want %v", pr.GetPriority(), data["priority"])
	}
	if !reflect.DeepEqual(pr.GetPurpose(), data["purpose"]) {
		t.Fatalf("GetPurpose() = %v, want %v", pr.GetPurpose(), data["purpose"])
	}
	if !reflect.DeepEqual(pr.GetPatient(), data["patient"]) {
		t.Fatalf("GetPatient() = %v, want %v", pr.GetPatient(), data["patient"])
	}
	if !reflect.DeepEqual(pr.GetServicedDate(), data["servicedDate"]) {
		t.Fatalf("GetServicedDate() = %v, want %v", pr.GetServicedDate(), data["servicedDate"])
	}
	if !reflect.DeepEqual(pr.GetCreated(), data["created"]) {
		t.Fatalf("GetCreated() = %v, want %v", pr.GetCreated(), data["created"])
	}
	if !reflect.DeepEqual(pr.GetEnterer(), data["enterer"]) {
		t.Fatalf("GetEnterer() = %v, want %v", pr.GetEnterer(), data["enterer"])
	}
	if !reflect.DeepEqual(pr.GetProvider(), data["provider"]) {
		t.Fatalf("GetProvider() = %v, want %v", pr.GetProvider(), data["provider"])
	}
	if !reflect.DeepEqual(pr.GetInsurer(), data["insurer"]) {
		t.Fatalf("GetInsurer() = %v, want %v", pr.GetInsurer(), data["insurer"])
	}
	if !reflect.DeepEqual(pr.GetFacility(), data["facility"]) {
		t.Fatalf("GetFacility() = %v, want %v", pr.GetFacility(), data["facility"])
	}
	if !reflect.DeepEqual(pr.GetInsurance(), data["insurance"]) {
		t.Fatalf("GetInsurance() = %v, want %v", pr.GetInsurance(), data["insurance"])
	}
	if !reflect.DeepEqual(pr.GetSupportingInfo(), data["supportingInfo"]) {
		t.Fatalf("GetSupportingInfo() = %v, want %v", pr.GetSupportingInfo(), data["supportingInfo"])
	}
	if !reflect.DeepEqual(pr.GetContained(), data["contained"]) {
		t.Fatalf("GetContained() = %v, want %v", pr.GetContained(), data["contained"])
	}
}
