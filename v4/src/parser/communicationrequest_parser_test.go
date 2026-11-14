package parser_test

// CommunicationRequestParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestCommunicationRequestParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "CommunicationRequest.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "CommunicationRequest" {
		t.Fatalf("resourceType = %v, want CommunicationRequest", data["resourceType"])
	}

	pr := parser.NewCommunicationRequestParser(data)
	if pr.GetRaw() == nil {
		t.Fatal("GetRaw() nil")
	}
	if !reflect.DeepEqual(pr.GetResourceType(), data["resourceType"]) {
		t.Fatalf("GetResourceType() = %v, want %v", pr.GetResourceType(), data["resourceType"])
	}
	if !reflect.DeepEqual(pr.GetStatus(), data["status"]) {
		t.Fatalf("GetStatus() = %v, want %v", pr.GetStatus(), data["status"])
	}
	if !reflect.DeepEqual(pr.GetCategory(), data["category"]) {
		t.Fatalf("GetCategory() = %v, want %v", pr.GetCategory(), data["category"])
	}
	if !reflect.DeepEqual(pr.GetPayload(), data["payload"]) {
		t.Fatalf("GetPayload() = %v, want %v", pr.GetPayload(), data["payload"])
	}
	if !reflect.DeepEqual(pr.GetAbout(), data["about"]) {
		t.Fatalf("GetAbout() = %v, want %v", pr.GetAbout(), data["about"])
	}
	if !reflect.DeepEqual(pr.GetSender(), data["sender"]) {
		t.Fatalf("GetSender() = %v, want %v", pr.GetSender(), data["sender"])
	}
	if !reflect.DeepEqual(pr.GetRecipient(), data["recipient"]) {
		t.Fatalf("GetRecipient() = %v, want %v", pr.GetRecipient(), data["recipient"])
	}
}
