package parser

// EnrollmentRequestParser extracts fields from a FHIR EnrollmentRequest resource map.
//
// The element list is derived from PayloadBuilderEnrollmentRequest in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type EnrollmentRequestParser struct {
	data map[string]any
}

// NewEnrollmentRequestParser creates a parser for the given EnrollmentRequest payload.
func NewEnrollmentRequestParser(data map[string]any) *EnrollmentRequestParser {
	if data == nil {
		data = map[string]any{}
	}
	return &EnrollmentRequestParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *EnrollmentRequestParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *EnrollmentRequestParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *EnrollmentRequestParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *EnrollmentRequestParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *EnrollmentRequestParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *EnrollmentRequestParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *EnrollmentRequestParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *EnrollmentRequestParser) GetEncounter() any {
	return p.data["encounter"]
}
