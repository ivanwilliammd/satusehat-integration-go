package parser_test

// BillingStatusParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestBillingStatusParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "BillingStatus.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "BillingStatus" {
		t.Fatalf("resourceType = %v, want BillingStatus", data["resourceType"])
	}

	pr := parser.NewBillingStatusParser(data)
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
	if !reflect.DeepEqual(pr.GetInsurer(), data["insurer"]) {
		t.Fatalf("GetInsurer() = %v, want %v", pr.GetInsurer(), data["insurer"])
	}
	if !reflect.DeepEqual(pr.GetRecipient(), data["recipient"]) {
		t.Fatalf("GetRecipient() = %v, want %v", pr.GetRecipient(), data["recipient"])
	}
	if !reflect.DeepEqual(pr.GetSubject(), data["subject"]) {
		t.Fatalf("GetSubject() = %v, want %v", pr.GetSubject(), data["subject"])
	}
	if !reflect.DeepEqual(pr.GetRequest(), data["request"]) {
		t.Fatalf("GetRequest() = %v, want %v", pr.GetRequest(), data["request"])
	}
}
