package parser_test

// CoverageParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestCoverageParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "Coverage.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "Coverage" {
		t.Fatalf("resourceType = %v, want Coverage", data["resourceType"])
	}

	pr := parser.NewCoverageParser(data)
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
	if !reflect.DeepEqual(pr.GetSubscriberId(), data["subscriberId"]) {
		t.Fatalf("GetSubscriberId() = %v, want %v", pr.GetSubscriberId(), data["subscriberId"])
	}
	if !reflect.DeepEqual(pr.GetBeneficiary(), data["beneficiary"]) {
		t.Fatalf("GetBeneficiary() = %v, want %v", pr.GetBeneficiary(), data["beneficiary"])
	}
	if !reflect.DeepEqual(pr.GetPayor(), data["payor"]) {
		t.Fatalf("GetPayor() = %v, want %v", pr.GetPayor(), data["payor"])
	}
	if !reflect.DeepEqual(pr.GetClass(), data["class"]) {
		t.Fatalf("GetClass() = %v, want %v", pr.GetClass(), data["class"])
	}
	if !reflect.DeepEqual(pr.GetExtension(), data["extension"]) {
		t.Fatalf("GetExtension() = %v, want %v", pr.GetExtension(), data["extension"])
	}
}
