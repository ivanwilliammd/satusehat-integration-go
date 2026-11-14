package parser

// AdverseEventParser extracts fields from a FHIR AdverseEvent resource map.
//
// The element list is derived from PayloadBuilderAdverseEvent in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type AdverseEventParser struct {
	data map[string]any
}

// NewAdverseEventParser creates a parser for the given AdverseEvent payload.
func NewAdverseEventParser(data map[string]any) *AdverseEventParser {
	if data == nil {
		data = map[string]any{}
	}
	return &AdverseEventParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *AdverseEventParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *AdverseEventParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *AdverseEventParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *AdverseEventParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *AdverseEventParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *AdverseEventParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *AdverseEventParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *AdverseEventParser) GetEncounter() any {
	return p.data["encounter"]
}
