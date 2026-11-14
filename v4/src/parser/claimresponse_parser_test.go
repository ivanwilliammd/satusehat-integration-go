package parser_test

// ClaimResponseParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestClaimResponseParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "ClaimResponse.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "ClaimResponse" {
		t.Fatalf("resourceType = %v, want ClaimResponse", data["resourceType"])
	}

	pr := parser.NewClaimResponseParser(data)
	if pr.GetRaw() == nil {
		t.Fatal("GetRaw() nil")
	}
	if !reflect.DeepEqual(pr.GetResourceType(), data["resourceType"]) {
		t.Fatalf("GetResourceType() = %v, want %v", pr.GetResourceType(), data["resourceType"])
	}
	if !reflect.DeepEqual(pr.GetStatus(), data["status"]) {
		t.Fatalf("GetStatus() = %v, want %v", pr.GetStatus(), data["status"])
	}
	if !reflect.DeepEqual(pr.GetIdentifier(), data["identifier"]) {
		t.Fatalf("GetIdentifier() = %v, want %v", pr.GetIdentifier(), data["identifier"])
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
	if !reflect.DeepEqual(pr.GetCreated(), data["created"]) {
		t.Fatalf("GetCreated() = %v, want %v", pr.GetCreated(), data["created"])
	}
	if !reflect.DeepEqual(pr.GetInsurer(), data["insurer"]) {
		t.Fatalf("GetInsurer() = %v, want %v", pr.GetInsurer(), data["insurer"])
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
	if !reflect.DeepEqual(pr.GetPayeeType(), data["payeeType"]) {
		t.Fatalf("GetPayeeType() = %v, want %v", pr.GetPayeeType(), data["payeeType"])
	}
	if !reflect.DeepEqual(pr.GetAdjudication(), data["adjudication"]) {
		t.Fatalf("GetAdjudication() = %v, want %v", pr.GetAdjudication(), data["adjudication"])
	}
	if !reflect.DeepEqual(pr.GetItem(), data["item"]) {
		t.Fatalf("GetItem() = %v, want %v", pr.GetItem(), data["item"])
	}
	if !reflect.DeepEqual(pr.GetTotal(), data["total"]) {
		t.Fatalf("GetTotal() = %v, want %v", pr.GetTotal(), data["total"])
	}
	if !reflect.DeepEqual(pr.GetPayment(), data["payment"]) {
		t.Fatalf("GetPayment() = %v, want %v", pr.GetPayment(), data["payment"])
	}
}
