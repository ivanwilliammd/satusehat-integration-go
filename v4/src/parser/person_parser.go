package parser

// PersonParser extracts fields from a FHIR Person resource map.
//
// The element list is derived from PayloadBuilderPerson in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type PersonParser struct {
	data map[string]any
}

// NewPersonParser creates a parser for the given Person payload.
func NewPersonParser(data map[string]any) *PersonParser {
	if data == nil {
		data = map[string]any{}
	}
	return &PersonParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *PersonParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *PersonParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *PersonParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *PersonParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *PersonParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *PersonParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *PersonParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *PersonParser) GetEncounter() any {
	return p.data["encounter"]
}
