package parser_test

// QuestionnaireResponseParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestQuestionnaireResponseParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "QuestionnaireResponse.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "QuestionnaireResponse" {
		t.Fatalf("resourceType = %v, want QuestionnaireResponse", data["resourceType"])
	}

	pr := parser.NewQuestionnaireResponseParser(data)
	if pr.GetRaw() == nil {
		t.Fatal("GetRaw() nil")
	}
	if !reflect.DeepEqual(pr.GetResourceType(), data["resourceType"]) {
		t.Fatalf("GetResourceType() = %v, want %v", pr.GetResourceType(), data["resourceType"])
	}
	if !reflect.DeepEqual(pr.GetQuestionnaire(), data["questionnaire"]) {
		t.Fatalf("GetQuestionnaire() = %v, want %v", pr.GetQuestionnaire(), data["questionnaire"])
	}
	if !reflect.DeepEqual(pr.GetStatus(), data["status"]) {
		t.Fatalf("GetStatus() = %v, want %v", pr.GetStatus(), data["status"])
	}
	if !reflect.DeepEqual(pr.GetSubject(), data["subject"]) {
		t.Fatalf("GetSubject() = %v, want %v", pr.GetSubject(), data["subject"])
	}
	if !reflect.DeepEqual(pr.GetEncounter(), data["encounter"]) {
		t.Fatalf("GetEncounter() = %v, want %v", pr.GetEncounter(), data["encounter"])
	}
	if !reflect.DeepEqual(pr.GetAuthored(), data["authored"]) {
		t.Fatalf("GetAuthored() = %v, want %v", pr.GetAuthored(), data["authored"])
	}
	if !reflect.DeepEqual(pr.GetAuthor(), data["author"]) {
		t.Fatalf("GetAuthor() = %v, want %v", pr.GetAuthor(), data["author"])
	}
	if !reflect.DeepEqual(pr.GetSource(), data["source"]) {
		t.Fatalf("GetSource() = %v, want %v", pr.GetSource(), data["source"])
	}
	if !reflect.DeepEqual(pr.GetItem(), data["item"]) {
		t.Fatalf("GetItem() = %v, want %v", pr.GetItem(), data["item"])
	}
}
