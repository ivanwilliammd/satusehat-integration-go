package parser_test

// ImmunizationParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestImmunizationParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "Immunization.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "Immunization" {
		t.Fatalf("resourceType = %v, want Immunization", data["resourceType"])
	}

	pr := parser.NewImmunizationParser(data)
	if pr.GetRaw() == nil {
		t.Fatal("GetRaw() nil")
	}
	if !reflect.DeepEqual(pr.GetResourceType(), data["resourceType"]) {
		t.Fatalf("GetResourceType() = %v, want %v", pr.GetResourceType(), data["resourceType"])
	}
	if !reflect.DeepEqual(pr.GetStatus(), data["status"]) {
		t.Fatalf("GetStatus() = %v, want %v", pr.GetStatus(), data["status"])
	}
	if !reflect.DeepEqual(pr.GetVaccineCode(), data["vaccineCode"]) {
		t.Fatalf("GetVaccineCode() = %v, want %v", pr.GetVaccineCode(), data["vaccineCode"])
	}
	if !reflect.DeepEqual(pr.GetPatient(), data["patient"]) {
		t.Fatalf("GetPatient() = %v, want %v", pr.GetPatient(), data["patient"])
	}
	if !reflect.DeepEqual(pr.GetEncounter(), data["encounter"]) {
		t.Fatalf("GetEncounter() = %v, want %v", pr.GetEncounter(), data["encounter"])
	}
	if !reflect.DeepEqual(pr.GetOccurrenceDateTime(), data["occurrenceDateTime"]) {
		t.Fatalf("GetOccurrenceDateTime() = %v, want %v", pr.GetOccurrenceDateTime(), data["occurrenceDateTime"])
	}
	if !reflect.DeepEqual(pr.GetRecorded(), data["recorded"]) {
		t.Fatalf("GetRecorded() = %v, want %v", pr.GetRecorded(), data["recorded"])
	}
	if !reflect.DeepEqual(pr.GetReaction(), data["reaction"]) {
		t.Fatalf("GetReaction() = %v, want %v", pr.GetReaction(), data["reaction"])
	}
	if !reflect.DeepEqual(pr.GetPrimarySource(), data["primarySource"]) {
		t.Fatalf("GetPrimarySource() = %v, want %v", pr.GetPrimarySource(), data["primarySource"])
	}
	if !reflect.DeepEqual(pr.GetLocation(), data["location"]) {
		t.Fatalf("GetLocation() = %v, want %v", pr.GetLocation(), data["location"])
	}
	if !reflect.DeepEqual(pr.GetLotNumber(), data["lotNumber"]) {
		t.Fatalf("GetLotNumber() = %v, want %v", pr.GetLotNumber(), data["lotNumber"])
	}
	if !reflect.DeepEqual(pr.GetExpirationDate(), data["expirationDate"]) {
		t.Fatalf("GetExpirationDate() = %v, want %v", pr.GetExpirationDate(), data["expirationDate"])
	}
	if !reflect.DeepEqual(pr.GetRoute(), data["route"]) {
		t.Fatalf("GetRoute() = %v, want %v", pr.GetRoute(), data["route"])
	}
	if !reflect.DeepEqual(pr.GetDoseQuantity(), data["doseQuantity"]) {
		t.Fatalf("GetDoseQuantity() = %v, want %v", pr.GetDoseQuantity(), data["doseQuantity"])
	}
	if !reflect.DeepEqual(pr.GetPerformer(), data["performer"]) {
		t.Fatalf("GetPerformer() = %v, want %v", pr.GetPerformer(), data["performer"])
	}
	if !reflect.DeepEqual(pr.GetReasonCode(), data["reasonCode"]) {
		t.Fatalf("GetReasonCode() = %v, want %v", pr.GetReasonCode(), data["reasonCode"])
	}
	if !reflect.DeepEqual(pr.GetProtocolApplied(), data["protocolApplied"]) {
		t.Fatalf("GetProtocolApplied() = %v, want %v", pr.GetProtocolApplied(), data["protocolApplied"])
	}
}
