package parser_test

// PaymentReconciliationParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestPaymentReconciliationParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "PaymentReconciliation.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "PaymentReconciliation" {
		t.Fatalf("resourceType = %v, want PaymentReconciliation", data["resourceType"])
	}

	pr := parser.NewPaymentReconciliationParser(data)
	if pr.GetRaw() == nil {
		t.Fatal("GetRaw() nil")
	}
	if !reflect.DeepEqual(pr.GetResourceType(), data["resourceType"]) {
		t.Fatalf("GetResourceType() = %v, want %v", pr.GetResourceType(), data["resourceType"])
	}
	if !reflect.DeepEqual(pr.GetStatus(), data["status"]) {
		t.Fatalf("GetStatus() = %v, want %v", pr.GetStatus(), data["status"])
	}
	if !reflect.DeepEqual(pr.GetPeriod(), data["period"]) {
		t.Fatalf("GetPeriod() = %v, want %v", pr.GetPeriod(), data["period"])
	}
	if !reflect.DeepEqual(pr.GetCreated(), data["created"]) {
		t.Fatalf("GetCreated() = %v, want %v", pr.GetCreated(), data["created"])
	}
	if !reflect.DeepEqual(pr.GetPaymentIssuer(), data["paymentIssuer"]) {
		t.Fatalf("GetPaymentIssuer() = %v, want %v", pr.GetPaymentIssuer(), data["paymentIssuer"])
	}
	if !reflect.DeepEqual(pr.GetRequestor(), data["requestor"]) {
		t.Fatalf("GetRequestor() = %v, want %v", pr.GetRequestor(), data["requestor"])
	}
	if !reflect.DeepEqual(pr.GetOutcome(), data["outcome"]) {
		t.Fatalf("GetOutcome() = %v, want %v", pr.GetOutcome(), data["outcome"])
	}
	if !reflect.DeepEqual(pr.GetDisposition(), data["disposition"]) {
		t.Fatalf("GetDisposition() = %v, want %v", pr.GetDisposition(), data["disposition"])
	}
	if !reflect.DeepEqual(pr.GetPaymentDate(), data["paymentDate"]) {
		t.Fatalf("GetPaymentDate() = %v, want %v", pr.GetPaymentDate(), data["paymentDate"])
	}
	if !reflect.DeepEqual(pr.GetPaymentAmount(), data["paymentAmount"]) {
		t.Fatalf("GetPaymentAmount() = %v, want %v", pr.GetPaymentAmount(), data["paymentAmount"])
	}
	if !reflect.DeepEqual(pr.GetDetail(), data["detail"]) {
		t.Fatalf("GetDetail() = %v, want %v", pr.GetDetail(), data["detail"])
	}
	if !reflect.DeepEqual(pr.GetProcessNote(), data["processNote"]) {
		t.Fatalf("GetProcessNote() = %v, want %v", pr.GetProcessNote(), data["processNote"])
	}
}
