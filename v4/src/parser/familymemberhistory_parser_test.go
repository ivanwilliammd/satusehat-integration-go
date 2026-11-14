package parser_test

// FamilyMemberHistoryParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestFamilyMemberHistoryParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "FamilyMemberHistory.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "FamilyMemberHistory" {
		t.Fatalf("resourceType = %v, want FamilyMemberHistory", data["resourceType"])
	}

	pr := parser.NewFamilyMemberHistoryParser(data)
	if pr.GetRaw() == nil {
		t.Fatal("GetRaw() nil")
	}
	if !reflect.DeepEqual(pr.GetResourceType(), data["resourceType"]) {
		t.Fatalf("GetResourceType() = %v, want %v", pr.GetResourceType(), data["resourceType"])
	}
	if !reflect.DeepEqual(pr.GetStatus(), data["status"]) {
		t.Fatalf("GetStatus() = %v, want %v", pr.GetStatus(), data["status"])
	}
	if !reflect.DeepEqual(pr.GetRelationship(), data["relationship"]) {
		t.Fatalf("GetRelationship() = %v, want %v", pr.GetRelationship(), data["relationship"])
	}
	if !reflect.DeepEqual(pr.GetDeceasedBoolean(), data["deceasedBoolean"]) {
		t.Fatalf("GetDeceasedBoolean() = %v, want %v", pr.GetDeceasedBoolean(), data["deceasedBoolean"])
	}
	if !reflect.DeepEqual(pr.GetPatient(), data["patient"]) {
		t.Fatalf("GetPatient() = %v, want %v", pr.GetPatient(), data["patient"])
	}
	if !reflect.DeepEqual(pr.GetDate(), data["date"]) {
		t.Fatalf("GetDate() = %v, want %v", pr.GetDate(), data["date"])
	}
	if !reflect.DeepEqual(pr.GetCondition(), data["condition"]) {
		t.Fatalf("GetCondition() = %v, want %v", pr.GetCondition(), data["condition"])
	}
}
