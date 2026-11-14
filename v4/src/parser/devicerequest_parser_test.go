package parser_test

// DeviceRequestParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestDeviceRequestParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "DeviceRequest.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "DeviceRequest" {
		t.Fatalf("resourceType = %v, want DeviceRequest", data["resourceType"])
	}

	pr := parser.NewDeviceRequestParser(data)
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
	if !reflect.DeepEqual(pr.GetIntent(), data["intent"]) {
		t.Fatalf("GetIntent() = %v, want %v", pr.GetIntent(), data["intent"])
	}
	if !reflect.DeepEqual(pr.GetPriority(), data["priority"]) {
		t.Fatalf("GetPriority() = %v, want %v", pr.GetPriority(), data["priority"])
	}
	if !reflect.DeepEqual(pr.GetCodeReference(), data["codeReference"]) {
		t.Fatalf("GetCodeReference() = %v, want %v", pr.GetCodeReference(), data["codeReference"])
	}
	if !reflect.DeepEqual(pr.GetSubject(), data["subject"]) {
		t.Fatalf("GetSubject() = %v, want %v", pr.GetSubject(), data["subject"])
	}
	if !reflect.DeepEqual(pr.GetAuthoredOn(), data["authoredOn"]) {
		t.Fatalf("GetAuthoredOn() = %v, want %v", pr.GetAuthoredOn(), data["authoredOn"])
	}
	if !reflect.DeepEqual(pr.GetRequester(), data["requester"]) {
		t.Fatalf("GetRequester() = %v, want %v", pr.GetRequester(), data["requester"])
	}
	if !reflect.DeepEqual(pr.GetPerformer(), data["performer"]) {
		t.Fatalf("GetPerformer() = %v, want %v", pr.GetPerformer(), data["performer"])
	}
	if !reflect.DeepEqual(pr.GetReasonCode(), data["reasonCode"]) {
		t.Fatalf("GetReasonCode() = %v, want %v", pr.GetReasonCode(), data["reasonCode"])
	}
}
