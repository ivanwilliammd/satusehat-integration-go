package parser_test

// EpisodeOfCareParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestEpisodeOfCareParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "EpisodeOfCare.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "EpisodeOfCare" {
		t.Fatalf("resourceType = %v, want EpisodeOfCare", data["resourceType"])
	}

	pr := parser.NewEpisodeOfCareParser(data)
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
	if !reflect.DeepEqual(pr.GetStatusHistory(), data["statusHistory"]) {
		t.Fatalf("GetStatusHistory() = %v, want %v", pr.GetStatusHistory(), data["statusHistory"])
	}
	if !reflect.DeepEqual(pr.GetType(), data["type"]) {
		t.Fatalf("GetType() = %v, want %v", pr.GetType(), data["type"])
	}
	if !reflect.DeepEqual(pr.GetDiagnosis(), data["diagnosis"]) {
		t.Fatalf("GetDiagnosis() = %v, want %v", pr.GetDiagnosis(), data["diagnosis"])
	}
	if !reflect.DeepEqual(pr.GetPatient(), data["patient"]) {
		t.Fatalf("GetPatient() = %v, want %v", pr.GetPatient(), data["patient"])
	}
	if !reflect.DeepEqual(pr.GetManagingOrganization(), data["managingOrganization"]) {
		t.Fatalf("GetManagingOrganization() = %v, want %v", pr.GetManagingOrganization(), data["managingOrganization"])
	}
	if !reflect.DeepEqual(pr.GetPeriod(), data["period"]) {
		t.Fatalf("GetPeriod() = %v, want %v", pr.GetPeriod(), data["period"])
	}
	if !reflect.DeepEqual(pr.GetCareManager(), data["careManager"]) {
		t.Fatalf("GetCareManager() = %v, want %v", pr.GetCareManager(), data["careManager"])
	}
}
