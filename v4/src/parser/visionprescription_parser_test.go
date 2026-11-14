package parser_test

// VisionPrescriptionParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestVisionPrescriptionParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "VisionPrescription.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "VisionPrescription" {
		t.Fatalf("resourceType = %v, want VisionPrescription", data["resourceType"])
	}

	pr := parser.NewVisionPrescriptionParser(data)
	if pr.GetRaw() == nil {
		t.Fatal("GetRaw() nil")
	}
	if !reflect.DeepEqual(pr.GetResourceType(), data["resourceType"]) {
		t.Fatalf("GetResourceType() = %v, want %v", pr.GetResourceType(), data["resourceType"])
	}
	if !reflect.DeepEqual(pr.GetStatus(), data["status"]) {
		t.Fatalf("GetStatus() = %v, want %v", pr.GetStatus(), data["status"])
	}
	if !reflect.DeepEqual(pr.GetCreated(), data["created"]) {
		t.Fatalf("GetCreated() = %v, want %v", pr.GetCreated(), data["created"])
	}
	if !reflect.DeepEqual(pr.GetPatient(), data["patient"]) {
		t.Fatalf("GetPatient() = %v, want %v", pr.GetPatient(), data["patient"])
	}
	if !reflect.DeepEqual(pr.GetEncounter(), data["encounter"]) {
		t.Fatalf("GetEncounter() = %v, want %v", pr.GetEncounter(), data["encounter"])
	}
	if !reflect.DeepEqual(pr.GetDateWritten(), data["dateWritten"]) {
		t.Fatalf("GetDateWritten() = %v, want %v", pr.GetDateWritten(), data["dateWritten"])
	}
	if !reflect.DeepEqual(pr.GetPrescriber(), data["prescriber"]) {
		t.Fatalf("GetPrescriber() = %v, want %v", pr.GetPrescriber(), data["prescriber"])
	}
	if !reflect.DeepEqual(pr.GetLensSpecification(), data["lensSpecification"]) {
		t.Fatalf("GetLensSpecification() = %v, want %v", pr.GetLensSpecification(), data["lensSpecification"])
	}
}
