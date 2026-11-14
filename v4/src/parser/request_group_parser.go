package parser

// RequestGroupParser extracts fields from a FHIR RequestGroup resource map.
//
// The element list is derived from PayloadBuilderRequestGroup in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type RequestGroupParser struct {
	data map[string]any
}

// NewRequestGroupParser creates a parser for the given RequestGroup payload.
func NewRequestGroupParser(data map[string]any) *RequestGroupParser {
	if data == nil {
		data = map[string]any{}
	}
	return &RequestGroupParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *RequestGroupParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *RequestGroupParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *RequestGroupParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *RequestGroupParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *RequestGroupParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *RequestGroupParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *RequestGroupParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *RequestGroupParser) GetEncounter() any {
	return p.data["encounter"]
}
