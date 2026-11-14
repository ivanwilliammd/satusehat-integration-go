package parser_test

// DeviceUseStatementParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestDeviceUseStatementParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "DeviceUseStatement.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "DeviceUseStatement" {
		t.Fatalf("resourceType = %v, want DeviceUseStatement", data["resourceType"])
	}

	pr := parser.NewDeviceUseStatementParser(data)
	if pr.GetRaw() == nil {
		t.Fatal("GetRaw() nil")
	}
	if !reflect.DeepEqual(pr.GetResourceType(), data["resourceType"]) {
		t.Fatalf("GetResourceType() = %v, want %v", pr.GetResourceType(), data["resourceType"])
	}
	if !reflect.DeepEqual(pr.GetIdentifier(), data["identifier"]) {
		t.Fatalf("GetIdentifier() = %v, want %v", pr.GetIdentifier(), data["identifier"])
	}
	if !reflect.DeepEqual(pr.GetStatus(), data["status"]) {
		t.Fatalf("GetStatus() = %v, want %v", pr.GetStatus(), data["status"])
	}
	if !reflect.DeepEqual(pr.GetDevice(), data["device"]) {
		t.Fatalf("GetDevice() = %v, want %v", pr.GetDevice(), data["device"])
	}
	if !reflect.DeepEqual(pr.GetSubject(), data["subject"]) {
		t.Fatalf("GetSubject() = %v, want %v", pr.GetSubject(), data["subject"])
	}
	if !reflect.DeepEqual(pr.GetTimingPeriod(), data["timingPeriod"]) {
		t.Fatalf("GetTimingPeriod() = %v, want %v", pr.GetTimingPeriod(), data["timingPeriod"])
	}
	if !reflect.DeepEqual(pr.GetRecordedOn(), data["recordedOn"]) {
		t.Fatalf("GetRecordedOn() = %v, want %v", pr.GetRecordedOn(), data["recordedOn"])
	}
	if !reflect.DeepEqual(pr.GetSource(), data["source"]) {
		t.Fatalf("GetSource() = %v, want %v", pr.GetSource(), data["source"])
	}
	if !reflect.DeepEqual(pr.GetReasonCode(), data["reasonCode"]) {
		t.Fatalf("GetReasonCode() = %v, want %v", pr.GetReasonCode(), data["reasonCode"])
	}
	if !reflect.DeepEqual(pr.GetNote(), data["note"]) {
		t.Fatalf("GetNote() = %v, want %v", pr.GetNote(), data["note"])
	}
}
