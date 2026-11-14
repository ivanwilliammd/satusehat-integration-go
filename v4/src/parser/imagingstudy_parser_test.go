package parser_test

// ImagingStudyParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestImagingStudyParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "ImagingStudy.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "ImagingStudy" {
		t.Fatalf("resourceType = %v, want ImagingStudy", data["resourceType"])
	}

	pr := parser.NewImagingStudyParser(data)
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
	if !reflect.DeepEqual(pr.GetModality(), data["modality"]) {
		t.Fatalf("GetModality() = %v, want %v", pr.GetModality(), data["modality"])
	}
	if !reflect.DeepEqual(pr.GetSubject(), data["subject"]) {
		t.Fatalf("GetSubject() = %v, want %v", pr.GetSubject(), data["subject"])
	}
	if !reflect.DeepEqual(pr.GetStarted(), data["started"]) {
		t.Fatalf("GetStarted() = %v, want %v", pr.GetStarted(), data["started"])
	}
	if !reflect.DeepEqual(pr.GetBasedOn(), data["basedOn"]) {
		t.Fatalf("GetBasedOn() = %v, want %v", pr.GetBasedOn(), data["basedOn"])
	}
	if !reflect.DeepEqual(pr.GetNumberOfSeries(), data["numberOfSeries"]) {
		t.Fatalf("GetNumberOfSeries() = %v, want %v", pr.GetNumberOfSeries(), data["numberOfSeries"])
	}
	if !reflect.DeepEqual(pr.GetNumberOfInstances(), data["numberOfInstances"]) {
		t.Fatalf("GetNumberOfInstances() = %v, want %v", pr.GetNumberOfInstances(), data["numberOfInstances"])
	}
	if !reflect.DeepEqual(pr.GetDescription(), data["description"]) {
		t.Fatalf("GetDescription() = %v, want %v", pr.GetDescription(), data["description"])
	}
	if !reflect.DeepEqual(pr.GetSeries(), data["series"]) {
		t.Fatalf("GetSeries() = %v, want %v", pr.GetSeries(), data["series"])
	}
}
