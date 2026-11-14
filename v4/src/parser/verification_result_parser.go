package parser

// VerificationResultParser extracts fields from a FHIR VerificationResult resource map.
//
// The element list is derived from PayloadBuilderVerificationResult in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type VerificationResultParser struct {
	data map[string]any
}

// NewVerificationResultParser creates a parser for the given VerificationResult payload.
func NewVerificationResultParser(data map[string]any) *VerificationResultParser {
	if data == nil {
		data = map[string]any{}
	}
	return &VerificationResultParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *VerificationResultParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *VerificationResultParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *VerificationResultParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *VerificationResultParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *VerificationResultParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *VerificationResultParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *VerificationResultParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *VerificationResultParser) GetEncounter() any {
	return p.data["encounter"]
}
