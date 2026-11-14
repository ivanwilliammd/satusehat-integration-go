package parser_test

// GoalParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestGoalParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "Goal.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "Goal" {
		t.Fatalf("resourceType = %v, want Goal", data["resourceType"])
	}

	pr := parser.NewGoalParser(data)
	if pr.GetRaw() == nil {
		t.Fatal("GetRaw() nil")
	}
	if !reflect.DeepEqual(pr.GetResourceType(), data["resourceType"]) {
		t.Fatalf("GetResourceType() = %v, want %v", pr.GetResourceType(), data["resourceType"])
	}
	if !reflect.DeepEqual(pr.GetId(), data["id"]) {
		t.Fatalf("GetId() = %v, want %v", pr.GetId(), data["id"])
	}
	if !reflect.DeepEqual(pr.GetLifecycleStatus(), data["lifecycleStatus"]) {
		t.Fatalf("GetLifecycleStatus() = %v, want %v", pr.GetLifecycleStatus(), data["lifecycleStatus"])
	}
	if !reflect.DeepEqual(pr.GetAchievementStatus(), data["achievementStatus"]) {
		t.Fatalf("GetAchievementStatus() = %v, want %v", pr.GetAchievementStatus(), data["achievementStatus"])
	}
	if !reflect.DeepEqual(pr.GetCategory(), data["category"]) {
		t.Fatalf("GetCategory() = %v, want %v", pr.GetCategory(), data["category"])
	}
	if !reflect.DeepEqual(pr.GetDescription(), data["description"]) {
		t.Fatalf("GetDescription() = %v, want %v", pr.GetDescription(), data["description"])
	}
	if !reflect.DeepEqual(pr.GetSubject(), data["subject"]) {
		t.Fatalf("GetSubject() = %v, want %v", pr.GetSubject(), data["subject"])
	}
	if !reflect.DeepEqual(pr.GetTarget(), data["target"]) {
		t.Fatalf("GetTarget() = %v, want %v", pr.GetTarget(), data["target"])
	}
	if !reflect.DeepEqual(pr.GetStatusDate(), data["statusDate"]) {
		t.Fatalf("GetStatusDate() = %v, want %v", pr.GetStatusDate(), data["statusDate"])
	}
	if !reflect.DeepEqual(pr.GetExpressedBy(), data["expressedBy"]) {
		t.Fatalf("GetExpressedBy() = %v, want %v", pr.GetExpressedBy(), data["expressedBy"])
	}
	if !reflect.DeepEqual(pr.GetAddresses(), data["addresses"]) {
		t.Fatalf("GetAddresses() = %v, want %v", pr.GetAddresses(), data["addresses"])
	}
	if !reflect.DeepEqual(pr.GetOutcomeCode(), data["outcomeCode"]) {
		t.Fatalf("GetOutcomeCode() = %v, want %v", pr.GetOutcomeCode(), data["outcomeCode"])
	}
}
