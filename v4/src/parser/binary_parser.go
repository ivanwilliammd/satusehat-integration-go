package parser

// BinaryParser extracts fields from a FHIR Binary resource map.
//
// The element list is derived from PayloadBuilderBinary in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type BinaryParser struct {
	data map[string]any
}

// NewBinaryParser creates a parser for the given Binary payload.
func NewBinaryParser(data map[string]any) *BinaryParser {
	if data == nil {
		data = map[string]any{}
	}
	return &BinaryParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *BinaryParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *BinaryParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *BinaryParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *BinaryParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *BinaryParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *BinaryParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *BinaryParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *BinaryParser) GetEncounter() any {
	return p.data["encounter"]
}
