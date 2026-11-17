package datatype

import "testing"

func TestPtrReturnsAddressOfCopy(t *testing.T) {
	p := Ptr("official")
	if p == nil || *p != "official" {
		t.Fatalf("Ptr(string) = %v", p)
	}

	n := Ptr(80.0)
	if n == nil || *n != 80.0 {
		t.Fatalf("Ptr(float64) = %v", n)
	}
}
