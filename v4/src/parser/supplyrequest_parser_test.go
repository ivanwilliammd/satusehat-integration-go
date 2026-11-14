package parser_test

// SupplyRequestParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestSupplyRequestParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "SupplyRequest.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "SupplyRequest" {
		t.Fatalf("resourceType = %v, want SupplyRequest", data["resourceType"])
	}

	pr := parser.NewSupplyRequestParser(data)
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
	if !reflect.DeepEqual(pr.GetItemReference(), data["itemReference"]) {
		t.Fatalf("GetItemReference() = %v, want %v", pr.GetItemReference(), data["itemReference"])
	}
	if !reflect.DeepEqual(pr.GetQuantity(), data["quantity"]) {
		t.Fatalf("GetQuantity() = %v, want %v", pr.GetQuantity(), data["quantity"])
	}
	if !reflect.DeepEqual(pr.GetAuthoredOn(), data["authoredOn"]) {
		t.Fatalf("GetAuthoredOn() = %v, want %v", pr.GetAuthoredOn(), data["authoredOn"])
	}
	if !reflect.DeepEqual(pr.GetRequester(), data["requester"]) {
		t.Fatalf("GetRequester() = %v, want %v", pr.GetRequester(), data["requester"])
	}
	if !reflect.DeepEqual(pr.GetSupplier(), data["supplier"]) {
		t.Fatalf("GetSupplier() = %v, want %v", pr.GetSupplier(), data["supplier"])
	}
	if !reflect.DeepEqual(pr.GetReasonCode(), data["reasonCode"]) {
		t.Fatalf("GetReasonCode() = %v, want %v", pr.GetReasonCode(), data["reasonCode"])
	}
}
