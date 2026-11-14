package parser_test

// ChargeItemResponseParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestChargeItemResponseParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "ChargeItemResponse.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "ChargeItemResponse" {
		t.Fatalf("resourceType = %v, want ChargeItemResponse", data["resourceType"])
	}

	pr := parser.NewChargeItemResponseParser(data)
	if pr.GetRaw() == nil {
		t.Fatal("GetRaw() nil")
	}
	if !reflect.DeepEqual(pr.GetResourceType(), data["resourceType"]) {
		t.Fatalf("GetResourceType() = %v, want %v", pr.GetResourceType(), data["resourceType"])
	}
	if !reflect.DeepEqual(pr.GetId(), data["id"]) {
		t.Fatalf("GetId() = %v, want %v", pr.GetId(), data["id"])
	}
	if !reflect.DeepEqual(pr.GetIdentifier(), data["identifier"]) {
		t.Fatalf("GetIdentifier() = %v, want %v", pr.GetIdentifier(), data["identifier"])
	}
	if !reflect.DeepEqual(pr.GetStatus(), data["status"]) {
		t.Fatalf("GetStatus() = %v, want %v", pr.GetStatus(), data["status"])
	}
	if !reflect.DeepEqual(pr.GetChargeItem(), data["chargeItem"]) {
		t.Fatalf("GetChargeItem() = %v, want %v", pr.GetChargeItem(), data["chargeItem"])
	}
	if !reflect.DeepEqual(pr.GetCoverage(), data["coverage"]) {
		t.Fatalf("GetCoverage() = %v, want %v", pr.GetCoverage(), data["coverage"])
	}
	if !reflect.DeepEqual(pr.GetSubject(), data["subject"]) {
		t.Fatalf("GetSubject() = %v, want %v", pr.GetSubject(), data["subject"])
	}
	if !reflect.DeepEqual(pr.GetEncounter(), data["encounter"]) {
		t.Fatalf("GetEncounter() = %v, want %v", pr.GetEncounter(), data["encounter"])
	}
	if !reflect.DeepEqual(pr.GetPayor(), data["payor"]) {
		t.Fatalf("GetPayor() = %v, want %v", pr.GetPayor(), data["payor"])
	}
	if !reflect.DeepEqual(pr.GetQuantity(), data["quantity"]) {
		t.Fatalf("GetQuantity() = %v, want %v", pr.GetQuantity(), data["quantity"])
	}
	if !reflect.DeepEqual(pr.GetUnitPriceOverride(), data["unitPriceOverride"]) {
		t.Fatalf("GetUnitPriceOverride() = %v, want %v", pr.GetUnitPriceOverride(), data["unitPriceOverride"])
	}
	if !reflect.DeepEqual(pr.GetFactorOverride(), data["factorOverride"]) {
		t.Fatalf("GetFactorOverride() = %v, want %v", pr.GetFactorOverride(), data["factorOverride"])
	}
	if !reflect.DeepEqual(pr.GetOverrideReason(), data["overrideReason"]) {
		t.Fatalf("GetOverrideReason() = %v, want %v", pr.GetOverrideReason(), data["overrideReason"])
	}
	if !reflect.DeepEqual(pr.GetTotal(), data["total"]) {
		t.Fatalf("GetTotal() = %v, want %v", pr.GetTotal(), data["total"])
	}
	if !reflect.DeepEqual(pr.GetExcess(), data["excess"]) {
		t.Fatalf("GetExcess() = %v, want %v", pr.GetExcess(), data["excess"])
	}
}
