package parser_test

// ConditionParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestConditionParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "Condition.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "Condition" {
		t.Fatalf("resourceType = %v, want Condition", data["resourceType"])
	}

	pr := parser.NewConditionParser(data)
	if pr.GetRaw() == nil {
		t.Fatal("GetRaw() nil")
	}
	if !reflect.DeepEqual(pr.GetResourceType(), data["resourceType"]) {
		t.Fatalf("GetResourceType() = %v, want %v", pr.GetResourceType(), data["resourceType"])
	}
	if !reflect.DeepEqual(pr.GetClinicalStatus(), data["clinicalStatus"]) {
		t.Fatalf("GetClinicalStatus() = %v, want %v", pr.GetClinicalStatus(), data["clinicalStatus"])
	}
	if !reflect.DeepEqual(pr.GetVerificationStatus(), data["verificationStatus"]) {
		t.Fatalf("GetVerificationStatus() = %v, want %v", pr.GetVerificationStatus(), data["verificationStatus"])
	}
	if !reflect.DeepEqual(pr.GetCategory(), data["category"]) {
		t.Fatalf("GetCategory() = %v, want %v", pr.GetCategory(), data["category"])
	}
	if !reflect.DeepEqual(pr.GetCode(), data["code"]) {
		t.Fatalf("GetCode() = %v, want %v", pr.GetCode(), data["code"])
	}
	if !reflect.DeepEqual(pr.GetStage(), data["stage"]) {
		t.Fatalf("GetStage() = %v, want %v", pr.GetStage(), data["stage"])
	}
	if !reflect.DeepEqual(pr.GetSubject(), data["subject"]) {
		t.Fatalf("GetSubject() = %v, want %v", pr.GetSubject(), data["subject"])
	}
	if !reflect.DeepEqual(pr.GetEncounter(), data["encounter"]) {
		t.Fatalf("GetEncounter() = %v, want %v", pr.GetEncounter(), data["encounter"])
	}
	if !reflect.DeepEqual(pr.GetRecordedDate(), data["recordedDate"]) {
		t.Fatalf("GetRecordedDate() = %v, want %v", pr.GetRecordedDate(), data["recordedDate"])
	}
}
