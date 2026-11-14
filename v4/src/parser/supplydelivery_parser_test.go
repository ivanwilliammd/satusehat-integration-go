package parser_test

// SupplyDeliveryParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestSupplyDeliveryParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "SupplyDelivery.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "SupplyDelivery" {
		t.Fatalf("resourceType = %v, want SupplyDelivery", data["resourceType"])
	}

	pr := parser.NewSupplyDeliveryParser(data)
	if pr.GetRaw() == nil {
		t.Fatal("GetRaw() nil")
	}
	if !reflect.DeepEqual(pr.GetResourceType(), data["resourceType"]) {
		t.Fatalf("GetResourceType() = %v, want %v", pr.GetResourceType(), data["resourceType"])
	}
	if !reflect.DeepEqual(pr.GetIdentifier(), data["identifier"]) {
		t.Fatalf("GetIdentifier() = %v, want %v", pr.GetIdentifier(), data["identifier"])
	}
	if !reflect.DeepEqual(pr.GetBasedOn(), data["basedOn"]) {
		t.Fatalf("GetBasedOn() = %v, want %v", pr.GetBasedOn(), data["basedOn"])
	}
	if !reflect.DeepEqual(pr.GetStatus(), data["status"]) {
		t.Fatalf("GetStatus() = %v, want %v", pr.GetStatus(), data["status"])
	}
	if !reflect.DeepEqual(pr.GetPatient(), data["patient"]) {
		t.Fatalf("GetPatient() = %v, want %v", pr.GetPatient(), data["patient"])
	}
	if !reflect.DeepEqual(pr.GetSuppliedItem(), data["suppliedItem"]) {
		t.Fatalf("GetSuppliedItem() = %v, want %v", pr.GetSuppliedItem(), data["suppliedItem"])
	}
	if !reflect.DeepEqual(pr.GetOccurrenceDateTime(), data["occurrenceDateTime"]) {
		t.Fatalf("GetOccurrenceDateTime() = %v, want %v", pr.GetOccurrenceDateTime(), data["occurrenceDateTime"])
	}
	if !reflect.DeepEqual(pr.GetSupplier(), data["supplier"]) {
		t.Fatalf("GetSupplier() = %v, want %v", pr.GetSupplier(), data["supplier"])
	}
	if !reflect.DeepEqual(pr.GetDestination(), data["destination"]) {
		t.Fatalf("GetDestination() = %v, want %v", pr.GetDestination(), data["destination"])
	}
	if !reflect.DeepEqual(pr.GetReceiver(), data["receiver"]) {
		t.Fatalf("GetReceiver() = %v, want %v", pr.GetReceiver(), data["receiver"])
	}
}
