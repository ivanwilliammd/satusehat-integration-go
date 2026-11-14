package parser_test

// CarePlanParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestCarePlanParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "CarePlan.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "CarePlan" {
		t.Fatalf("resourceType = %v, want CarePlan", data["resourceType"])
	}

	pr := parser.NewCarePlanParser(data)
	if pr.GetRaw() == nil {
		t.Fatal("GetRaw() nil")
	}
	if !reflect.DeepEqual(pr.GetResourceType(), data["resourceType"]) {
		t.Fatalf("GetResourceType() = %v, want %v", pr.GetResourceType(), data["resourceType"])
	}
	if !reflect.DeepEqual(pr.GetTitle(), data["title"]) {
		t.Fatalf("GetTitle() = %v, want %v", pr.GetTitle(), data["title"])
	}
	if !reflect.DeepEqual(pr.GetDescription(), data["description"]) {
		t.Fatalf("GetDescription() = %v, want %v", pr.GetDescription(), data["description"])
	}
	if !reflect.DeepEqual(pr.GetCategory(), data["category"]) {
		t.Fatalf("GetCategory() = %v, want %v", pr.GetCategory(), data["category"])
	}
	if !reflect.DeepEqual(pr.GetStatus(), data["status"]) {
		t.Fatalf("GetStatus() = %v, want %v", pr.GetStatus(), data["status"])
	}
	if !reflect.DeepEqual(pr.GetIntent(), data["intent"]) {
		t.Fatalf("GetIntent() = %v, want %v", pr.GetIntent(), data["intent"])
	}
	if !reflect.DeepEqual(pr.GetSubject(), data["subject"]) {
		t.Fatalf("GetSubject() = %v, want %v", pr.GetSubject(), data["subject"])
	}
	if !reflect.DeepEqual(pr.GetEncounter(), data["encounter"]) {
		t.Fatalf("GetEncounter() = %v, want %v", pr.GetEncounter(), data["encounter"])
	}
	if !reflect.DeepEqual(pr.GetPeriod(), data["period"]) {
		t.Fatalf("GetPeriod() = %v, want %v", pr.GetPeriod(), data["period"])
	}
	if !reflect.DeepEqual(pr.GetCreated(), data["created"]) {
		t.Fatalf("GetCreated() = %v, want %v", pr.GetCreated(), data["created"])
	}
	if !reflect.DeepEqual(pr.GetAuthor(), data["author"]) {
		t.Fatalf("GetAuthor() = %v, want %v", pr.GetAuthor(), data["author"])
	}
	if !reflect.DeepEqual(pr.GetActivity(), data["activity"]) {
		t.Fatalf("GetActivity() = %v, want %v", pr.GetActivity(), data["activity"])
	}
}
