package parser

// EnrollmentResponseParser extracts fields from a FHIR EnrollmentResponse resource map.
//
// The element list is derived from PayloadBuilderEnrollmentResponse in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type EnrollmentResponseParser struct {
	data map[string]any
}

// NewEnrollmentResponseParser creates a parser for the given EnrollmentResponse payload.
func NewEnrollmentResponseParser(data map[string]any) *EnrollmentResponseParser {
	if data == nil {
		data = map[string]any{}
	}
	return &EnrollmentResponseParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *EnrollmentResponseParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *EnrollmentResponseParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *EnrollmentResponseParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *EnrollmentResponseParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *EnrollmentResponseParser) GetStatus() any {
	return p.data["status"]
}

// GetRequest returns the request element.
func (p *EnrollmentResponseParser) GetRequest() any {
	return p.data["request"]
}

// GetOutcome returns the outcome element.
func (p *EnrollmentResponseParser) GetOutcome() any {
	return p.data["outcome"]
}

// GetDisposition returns the disposition element.
func (p *EnrollmentResponseParser) GetDisposition() any {
	return p.data["disposition"]
}

// GetCreated returns the created element.
func (p *EnrollmentResponseParser) GetCreated() any {
	return p.data["created"]
}
