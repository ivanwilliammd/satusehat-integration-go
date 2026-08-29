package datatype

import "testing"

func TestIdentifier(t *testing.T) {
	id := Identifier{System: "https://fhir.kemkes.go.id/id/nik", Value: "1234567890"}
	if id.Value != "1234567890" {
		t.Errorf("Expected NIK to be 1234567890, got %s", id.Value)
	}
}
