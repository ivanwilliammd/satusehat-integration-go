package parser_test

// PaymentNoticeParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestPaymentNoticeParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "PaymentNotice.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "PaymentNotice" {
		t.Fatalf("resourceType = %v, want PaymentNotice", data["resourceType"])
	}

	pr := parser.NewPaymentNoticeParser(data)
	if pr.GetRaw() == nil {
		t.Fatal("GetRaw() nil")
	}
	if !reflect.DeepEqual(pr.GetResourceType(), data["resourceType"]) {
		t.Fatalf("GetResourceType() = %v, want %v", pr.GetResourceType(), data["resourceType"])
	}
	if !reflect.DeepEqual(pr.GetStatus(), data["status"]) {
		t.Fatalf("GetStatus() = %v, want %v", pr.GetStatus(), data["status"])
	}
	if !reflect.DeepEqual(pr.GetRequest(), data["request"]) {
		t.Fatalf("GetRequest() = %v, want %v", pr.GetRequest(), data["request"])
	}
	if !reflect.DeepEqual(pr.GetResponse(), data["response"]) {
		t.Fatalf("GetResponse() = %v, want %v", pr.GetResponse(), data["response"])
	}
	if !reflect.DeepEqual(pr.GetCreated(), data["created"]) {
		t.Fatalf("GetCreated() = %v, want %v", pr.GetCreated(), data["created"])
	}
	if !reflect.DeepEqual(pr.GetPaymentDate(), data["paymentDate"]) {
		t.Fatalf("GetPaymentDate() = %v, want %v", pr.GetPaymentDate(), data["paymentDate"])
	}
	if !reflect.DeepEqual(pr.GetProvider(), data["provider"]) {
		t.Fatalf("GetProvider() = %v, want %v", pr.GetProvider(), data["provider"])
	}
	if !reflect.DeepEqual(pr.GetPayee(), data["payee"]) {
		t.Fatalf("GetPayee() = %v, want %v", pr.GetPayee(), data["payee"])
	}
	if !reflect.DeepEqual(pr.GetRecipient(), data["recipient"]) {
		t.Fatalf("GetRecipient() = %v, want %v", pr.GetRecipient(), data["recipient"])
	}
	if !reflect.DeepEqual(pr.GetAmount(), data["amount"]) {
		t.Fatalf("GetAmount() = %v, want %v", pr.GetAmount(), data["amount"])
	}
	if !reflect.DeepEqual(pr.GetPayment(), data["payment"]) {
		t.Fatalf("GetPayment() = %v, want %v", pr.GetPayment(), data["payment"])
	}
	if !reflect.DeepEqual(pr.GetPaymentStatus(), data["paymentStatus"]) {
		t.Fatalf("GetPaymentStatus() = %v, want %v", pr.GetPaymentStatus(), data["paymentStatus"])
	}
}
