package parser

// CapabilityStatementParser extracts fields from a FHIR CapabilityStatement resource map.
//
// The element list is derived from PayloadBuilderCapabilityStatement in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type CapabilityStatementParser struct {
	data map[string]any
}

// NewCapabilityStatementParser creates a parser for the given CapabilityStatement payload.
func NewCapabilityStatementParser(data map[string]any) *CapabilityStatementParser {
	if data == nil {
		data = map[string]any{}
	}
	return &CapabilityStatementParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *CapabilityStatementParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *CapabilityStatementParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *CapabilityStatementParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *CapabilityStatementParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *CapabilityStatementParser) GetStatus() any {
	return p.data["status"]
}

// GetDate returns the date element.
func (p *CapabilityStatementParser) GetDate() any {
	return p.data["date"]
}

// GetKind returns the kind element.
func (p *CapabilityStatementParser) GetKind() any {
	return p.data["kind"]
}

// GetFhirVersion returns the fhirVersion element.
func (p *CapabilityStatementParser) GetFhirVersion() any {
	return p.data["fhirVersion"]
}

// GetFormat returns the format element.
func (p *CapabilityStatementParser) GetFormat() any {
	return p.data["format"]
}

// GetRest returns the rest element.
func (p *CapabilityStatementParser) GetRest() any {
	return p.data["rest"]
}
