package builder_test

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
)

// A dropped argument cannot be seen in a method count or a compile. These two checks read the payload
// and assert the reference survived.
func TestSetSubjectKeepsTheReference(t *testing.T) {
	kasus := map[string]func() map[string]interface{}{
		"account": func() map[string]interface{} {
			b := builder.NewAccountBuilder()
			b.SetSubject("Patient/1000000001")
			return b.Build()
		},
		"claim": func() map[string]interface{} {
			b := builder.NewClaimBuilder()
			b.SetSubject("Patient/1000000001")
			return b.Build()
		},
		"audit_event": func() map[string]interface{} {
			b := builder.NewAuditEventBuilder()
			b.SetSubject("Patient/1000000001")
			return b.Build()
		},
	}

	for nama, buat := range kasus {
		subject, ada := buat()["subject"].(map[string]interface{})
		if !ada {
			t.Fatalf("%s: subject tidak ditulis", nama)
		}
		if subject["reference"] != "Patient/1000000001" {
			t.Fatalf("%s: reference = %v, seharusnya Patient/1000000001", nama, subject["reference"])
		}
	}
}
