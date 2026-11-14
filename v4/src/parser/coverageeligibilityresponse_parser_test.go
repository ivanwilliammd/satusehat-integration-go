package parser_test

// CoverageEligibilityResponseParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestCoverageEligibilityResponseParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "CoverageEligibilityResponse.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "CoverageEligibilityResponse" {
		t.Fatalf("resourceType = %v, want CoverageEligibilityResponse", data["resourceType"])
	}

	pr := parser.NewCoverageEligibilityResponseParser(data)
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
	if !reflect.DeepEqual(pr.GetPurpose(), data["purpose"]) {
		t.Fatalf("GetPurpose() = %v, want %v", pr.GetPurpose(), data["purpose"])
	}
	if !reflect.DeepEqual(pr.GetPatient(), data["patient"]) {
		t.Fatalf("GetPatient() = %v, want %v", pr.GetPatient(), data["patient"])
	}
	if !reflect.DeepEqual(pr.GetCreated(), data["created"]) {
		t.Fatalf("GetCreated() = %v, want %v", pr.GetCreated(), data["created"])
	}
	if !reflect.DeepEqual(pr.GetRequestor(), data["requestor"]) {
		t.Fatalf("GetRequestor() = %v, want %v", pr.GetRequestor(), data["requestor"])
	}
	if !reflect.DeepEqual(pr.GetRequest(), data["request"]) {
		t.Fatalf("GetRequest() = %v, want %v", pr.GetRequest(), data["request"])
	}
	if !reflect.DeepEqual(pr.GetOutcome(), data["outcome"]) {
		t.Fatalf("GetOutcome() = %v, want %v", pr.GetOutcome(), data["outcome"])
	}
	if !reflect.DeepEqual(pr.GetDisposition(), data["disposition"]) {
		t.Fatalf("GetDisposition() = %v, want %v", pr.GetDisposition(), data["disposition"])
	}
	if !reflect.DeepEqual(pr.GetInsurer(), data["insurer"]) {
		t.Fatalf("GetInsurer() = %v, want %v", pr.GetInsurer(), data["insurer"])
	}
	if !reflect.DeepEqual(pr.GetInsurance(), data["insurance"]) {
		t.Fatalf("GetInsurance() = %v, want %v", pr.GetInsurance(), data["insurance"])
	}
	if !reflect.DeepEqual(pr.GetExtension(), data["extension"]) {
		t.Fatalf("GetExtension() = %v, want %v", pr.GetExtension(), data["extension"])
	}
}
