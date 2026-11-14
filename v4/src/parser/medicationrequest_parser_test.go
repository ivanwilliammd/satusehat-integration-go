package parser_test

// MedicationRequestParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestMedicationRequestParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "MedicationRequest.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "MedicationRequest" {
		t.Fatalf("resourceType = %v, want MedicationRequest", data["resourceType"])
	}

	pr := parser.NewMedicationRequestParser(data)
	if pr.GetRaw() == nil {
		t.Fatal("GetRaw() nil")
	}
	if !reflect.DeepEqual(pr.GetResourceType(), data["resourceType"]) {
		t.Fatalf("GetResourceType() = %v, want %v", pr.GetResourceType(), data["resourceType"])
	}
	if !reflect.DeepEqual(pr.GetContained(), data["contained"]) {
		t.Fatalf("GetContained() = %v, want %v", pr.GetContained(), data["contained"])
	}
	if !reflect.DeepEqual(pr.GetIdentifier(), data["identifier"]) {
		t.Fatalf("GetIdentifier() = %v, want %v", pr.GetIdentifier(), data["identifier"])
	}
	if !reflect.DeepEqual(pr.GetStatus(), data["status"]) {
		t.Fatalf("GetStatus() = %v, want %v", pr.GetStatus(), data["status"])
	}
	if !reflect.DeepEqual(pr.GetIntent(), data["intent"]) {
		t.Fatalf("GetIntent() = %v, want %v", pr.GetIntent(), data["intent"])
	}
	if !reflect.DeepEqual(pr.GetCategory(), data["category"]) {
		t.Fatalf("GetCategory() = %v, want %v", pr.GetCategory(), data["category"])
	}
	if !reflect.DeepEqual(pr.GetPriority(), data["priority"]) {
		t.Fatalf("GetPriority() = %v, want %v", pr.GetPriority(), data["priority"])
	}
	if !reflect.DeepEqual(pr.GetMedicationReference(), data["medicationReference"]) {
		t.Fatalf("GetMedicationReference() = %v, want %v", pr.GetMedicationReference(), data["medicationReference"])
	}
	if !reflect.DeepEqual(pr.GetSubject(), data["subject"]) {
		t.Fatalf("GetSubject() = %v, want %v", pr.GetSubject(), data["subject"])
	}
	if !reflect.DeepEqual(pr.GetEncounter(), data["encounter"]) {
		t.Fatalf("GetEncounter() = %v, want %v", pr.GetEncounter(), data["encounter"])
	}
	if !reflect.DeepEqual(pr.GetAuthoredOn(), data["authoredOn"]) {
		t.Fatalf("GetAuthoredOn() = %v, want %v", pr.GetAuthoredOn(), data["authoredOn"])
	}
	if !reflect.DeepEqual(pr.GetRequester(), data["requester"]) {
		t.Fatalf("GetRequester() = %v, want %v", pr.GetRequester(), data["requester"])
	}
	if !reflect.DeepEqual(pr.GetCourseOfTherapyType(), data["courseOfTherapyType"]) {
		t.Fatalf("GetCourseOfTherapyType() = %v, want %v", pr.GetCourseOfTherapyType(), data["courseOfTherapyType"])
	}
	if !reflect.DeepEqual(pr.GetDosageInstruction(), data["dosageInstruction"]) {
		t.Fatalf("GetDosageInstruction() = %v, want %v", pr.GetDosageInstruction(), data["dosageInstruction"])
	}
	if !reflect.DeepEqual(pr.GetDispenseRequest(), data["dispenseRequest"]) {
		t.Fatalf("GetDispenseRequest() = %v, want %v", pr.GetDispenseRequest(), data["dispenseRequest"])
	}
}
