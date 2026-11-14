package parser_test

// CompositionParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestCompositionParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "Composition.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "Composition" {
		t.Fatalf("resourceType = %v, want Composition", data["resourceType"])
	}

	pr := parser.NewCompositionParser(data)
	if pr.GetRaw() == nil {
		t.Fatal("GetRaw() nil")
	}
	if !reflect.DeepEqual(pr.GetResourceType(), data["resourceType"]) {
		t.Fatalf("GetResourceType() = %v, want %v", pr.GetResourceType(), data["resourceType"])
	}
	if !reflect.DeepEqual(pr.GetStatus(), data["status"]) {
		t.Fatalf("GetStatus() = %v, want %v", pr.GetStatus(), data["status"])
	}
	if !reflect.DeepEqual(pr.GetAttester(), data["attester"]) {
		t.Fatalf("GetAttester() = %v, want %v", pr.GetAttester(), data["attester"])
	}
	if !reflect.DeepEqual(pr.GetType(), data["type"]) {
		t.Fatalf("GetType() = %v, want %v", pr.GetType(), data["type"])
	}
	if !reflect.DeepEqual(pr.GetSubject(), data["subject"]) {
		t.Fatalf("GetSubject() = %v, want %v", pr.GetSubject(), data["subject"])
	}
	if !reflect.DeepEqual(pr.GetEncounter(), data["encounter"]) {
		t.Fatalf("GetEncounter() = %v, want %v", pr.GetEncounter(), data["encounter"])
	}
	if !reflect.DeepEqual(pr.GetDate(), data["date"]) {
		t.Fatalf("GetDate() = %v, want %v", pr.GetDate(), data["date"])
	}
	if !reflect.DeepEqual(pr.GetAuthor(), data["author"]) {
		t.Fatalf("GetAuthor() = %v, want %v", pr.GetAuthor(), data["author"])
	}
	if !reflect.DeepEqual(pr.GetTitle(), data["title"]) {
		t.Fatalf("GetTitle() = %v, want %v", pr.GetTitle(), data["title"])
	}
	if !reflect.DeepEqual(pr.GetSection(), data["section"]) {
		t.Fatalf("GetSection() = %v, want %v", pr.GetSection(), data["section"])
	}
}
