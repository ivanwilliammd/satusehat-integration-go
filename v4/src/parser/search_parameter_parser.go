package parser

// SearchParameterParser extracts fields from a FHIR SearchParameter resource map.
//
// The element list is derived from PayloadBuilderSearchParameter in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type SearchParameterParser struct {
	data map[string]any
}

// NewSearchParameterParser creates a parser for the given SearchParameter payload.
func NewSearchParameterParser(data map[string]any) *SearchParameterParser {
	if data == nil {
		data = map[string]any{}
	}
	return &SearchParameterParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *SearchParameterParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *SearchParameterParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *SearchParameterParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *SearchParameterParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *SearchParameterParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *SearchParameterParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *SearchParameterParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *SearchParameterParser) GetEncounter() any {
	return p.data["encounter"]
}
