package parser_test

// AccountParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestAccountParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "Account.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "Account" {
		t.Fatalf("resourceType = %v, want Account", data["resourceType"])
	}

	pr := parser.NewAccountParser(data)
	if pr.GetRaw() == nil {
		t.Fatal("GetRaw() nil")
	}
	if !reflect.DeepEqual(pr.GetResourceType(), data["resourceType"]) {
		t.Fatalf("GetResourceType() = %v, want %v", pr.GetResourceType(), data["resourceType"])
	}
	if !reflect.DeepEqual(pr.GetStatus(), data["status"]) {
		t.Fatalf("GetStatus() = %v, want %v", pr.GetStatus(), data["status"])
	}
	if !reflect.DeepEqual(pr.GetType(), data["type"]) {
		t.Fatalf("GetType() = %v, want %v", pr.GetType(), data["type"])
	}
	if !reflect.DeepEqual(pr.GetName(), data["name"]) {
		t.Fatalf("GetName() = %v, want %v", pr.GetName(), data["name"])
	}
	if !reflect.DeepEqual(pr.GetSubject(), data["subject"]) {
		t.Fatalf("GetSubject() = %v, want %v", pr.GetSubject(), data["subject"])
	}
	if !reflect.DeepEqual(pr.GetServicePeriod(), data["servicePeriod"]) {
		t.Fatalf("GetServicePeriod() = %v, want %v", pr.GetServicePeriod(), data["servicePeriod"])
	}
	if !reflect.DeepEqual(pr.GetOwner(), data["owner"]) {
		t.Fatalf("GetOwner() = %v, want %v", pr.GetOwner(), data["owner"])
	}
	if !reflect.DeepEqual(pr.GetGuarantor(), data["guarantor"]) {
		t.Fatalf("GetGuarantor() = %v, want %v", pr.GetGuarantor(), data["guarantor"])
	}
	if !reflect.DeepEqual(pr.GetDescription(), data["description"]) {
		t.Fatalf("GetDescription() = %v, want %v", pr.GetDescription(), data["description"])
	}
}
