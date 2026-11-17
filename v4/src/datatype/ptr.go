package datatype

// Ptr returns a pointer to v.
//
// Optional FHIR fields in this package are pointers so that "not set" stays distinguishable from the
// zero value; Ptr keeps that requirement from turning every payload into a block of local variables:
//
//	Use:    datatype.Ptr("official"),
//	Value:  datatype.Ptr(80.0),
//	Start:  datatype.Ptr("2026-11-17T02:00:00+00:00"),
func Ptr[T any](v T) *T {
	return &v
}
