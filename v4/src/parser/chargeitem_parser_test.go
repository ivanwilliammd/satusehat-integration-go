package parser_test

// ChargeItemParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestChargeItemParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "ChargeItem.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "ChargeItem" {
		t.Fatalf("resourceType = %v, want ChargeItem", data["resourceType"])
	}

	pr := parser.NewChargeItemParser(data)
	if pr.GetRaw() == nil {
		t.Fatal("GetRaw() nil")
	}
	if !reflect.DeepEqual(pr.GetResourceType(), data["resourceType"]) {
		t.Fatalf("GetResourceType() = %v, want %v", pr.GetResourceType(), data["resourceType"])
	}
	if !reflect.DeepEqual(pr.GetStatus(), data["status"]) {
		t.Fatalf("GetStatus() = %v, want %v", pr.GetStatus(), data["status"])
	}
	if !reflect.DeepEqual(pr.GetCode(), data["code"]) {
		t.Fatalf("GetCode() = %v, want %v", pr.GetCode(), data["code"])
	}
	if !reflect.DeepEqual(pr.GetSubject(), data["subject"]) {
		t.Fatalf("GetSubject() = %v, want %v", pr.GetSubject(), data["subject"])
	}
	if !reflect.DeepEqual(pr.GetContext(), data["context"]) {
		t.Fatalf("GetContext() = %v, want %v", pr.GetContext(), data["context"])
	}
	if !reflect.DeepEqual(pr.GetOccurrencePeriod(), data["occurrencePeriod"]) {
		t.Fatalf("GetOccurrencePeriod() = %v, want %v", pr.GetOccurrencePeriod(), data["occurrencePeriod"])
	}
	if !reflect.DeepEqual(pr.GetPerformer(), data["performer"]) {
		t.Fatalf("GetPerformer() = %v, want %v", pr.GetPerformer(), data["performer"])
	}
	if !reflect.DeepEqual(pr.GetQuantity(), data["quantity"]) {
		t.Fatalf("GetQuantity() = %v, want %v", pr.GetQuantity(), data["quantity"])
	}
	if !reflect.DeepEqual(pr.GetAccount(), data["account"]) {
		t.Fatalf("GetAccount() = %v, want %v", pr.GetAccount(), data["account"])
	}
	if !reflect.DeepEqual(pr.GetExtension(), data["extension"]) {
		t.Fatalf("GetExtension() = %v, want %v", pr.GetExtension(), data["extension"])
	}
}
