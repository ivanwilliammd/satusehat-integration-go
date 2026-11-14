package parser

// BasicParser extracts fields from a FHIR Basic resource map.
//
// The element list is derived from PayloadBuilderBasic in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type BasicParser struct {
	data map[string]any
}

// NewBasicParser creates a parser for the given Basic payload.
func NewBasicParser(data map[string]any) *BasicParser {
	if data == nil {
		data = map[string]any{}
	}
	return &BasicParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *BasicParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *BasicParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *BasicParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *BasicParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *BasicParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *BasicParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *BasicParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *BasicParser) GetEncounter() any {
	return p.data["encounter"]
}
