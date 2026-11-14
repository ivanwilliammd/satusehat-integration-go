package parser_test

// DeviceDispenseParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestDeviceDispenseParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "DeviceDispense.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "DeviceDispense" {
		t.Fatalf("resourceType = %v, want DeviceDispense", data["resourceType"])
	}

	pr := parser.NewDeviceDispenseParser(data)
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
	if !reflect.DeepEqual(pr.GetBasedOn(), data["basedOn"]) {
		t.Fatalf("GetBasedOn() = %v, want %v", pr.GetBasedOn(), data["basedOn"])
	}
	if !reflect.DeepEqual(pr.GetDevice(), data["device"]) {
		t.Fatalf("GetDevice() = %v, want %v", pr.GetDevice(), data["device"])
	}
	if !reflect.DeepEqual(pr.GetSubject(), data["subject"]) {
		t.Fatalf("GetSubject() = %v, want %v", pr.GetSubject(), data["subject"])
	}
	if !reflect.DeepEqual(pr.GetReceiver(), data["receiver"]) {
		t.Fatalf("GetReceiver() = %v, want %v", pr.GetReceiver(), data["receiver"])
	}
	if !reflect.DeepEqual(pr.GetPerformer(), data["performer"]) {
		t.Fatalf("GetPerformer() = %v, want %v", pr.GetPerformer(), data["performer"])
	}
	if !reflect.DeepEqual(pr.GetWhenPrepared(), data["whenPrepared"]) {
		t.Fatalf("GetWhenPrepared() = %v, want %v", pr.GetWhenPrepared(), data["whenPrepared"])
	}
	if !reflect.DeepEqual(pr.GetWhenHandedOver(), data["whenHandedOver"]) {
		t.Fatalf("GetWhenHandedOver() = %v, want %v", pr.GetWhenHandedOver(), data["whenHandedOver"])
	}
	if !reflect.DeepEqual(pr.GetDestination(), data["destination"]) {
		t.Fatalf("GetDestination() = %v, want %v", pr.GetDestination(), data["destination"])
	}
	if !reflect.DeepEqual(pr.GetNote(), data["note"]) {
		t.Fatalf("GetNote() = %v, want %v", pr.GetNote(), data["note"])
	}
}
