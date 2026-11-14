package parser_test

// InvoiceParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestInvoiceParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "Invoice.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "Invoice" {
		t.Fatalf("resourceType = %v, want Invoice", data["resourceType"])
	}

	pr := parser.NewInvoiceParser(data)
	if pr.GetRaw() == nil {
		t.Fatal("GetRaw() nil")
	}
	if !reflect.DeepEqual(pr.GetResourceType(), data["resourceType"]) {
		t.Fatalf("GetResourceType() = %v, want %v", pr.GetResourceType(), data["resourceType"])
	}
	if !reflect.DeepEqual(pr.GetStatus(), data["status"]) {
		t.Fatalf("GetStatus() = %v, want %v", pr.GetStatus(), data["status"])
	}
	if !reflect.DeepEqual(pr.GetSubject(), data["subject"]) {
		t.Fatalf("GetSubject() = %v, want %v", pr.GetSubject(), data["subject"])
	}
	if !reflect.DeepEqual(pr.GetDate(), data["date"]) {
		t.Fatalf("GetDate() = %v, want %v", pr.GetDate(), data["date"])
	}
	if !reflect.DeepEqual(pr.GetIssuer(), data["issuer"]) {
		t.Fatalf("GetIssuer() = %v, want %v", pr.GetIssuer(), data["issuer"])
	}
	if !reflect.DeepEqual(pr.GetAccount(), data["account"]) {
		t.Fatalf("GetAccount() = %v, want %v", pr.GetAccount(), data["account"])
	}
	if !reflect.DeepEqual(pr.GetLineItem(), data["lineItem"]) {
		t.Fatalf("GetLineItem() = %v, want %v", pr.GetLineItem(), data["lineItem"])
	}
	if !reflect.DeepEqual(pr.GetTotalNet(), data["totalNet"]) {
		t.Fatalf("GetTotalNet() = %v, want %v", pr.GetTotalNet(), data["totalNet"])
	}
}
