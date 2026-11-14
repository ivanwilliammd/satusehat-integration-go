package parser

// ActivityDefinitionParser extracts fields from a FHIR ActivityDefinition resource map.
//
// The element list is derived from PayloadBuilderActivityDefinition in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type ActivityDefinitionParser struct {
	data map[string]any
}

// NewActivityDefinitionParser creates a parser for the given ActivityDefinition payload.
func NewActivityDefinitionParser(data map[string]any) *ActivityDefinitionParser {
	if data == nil {
		data = map[string]any{}
	}
	return &ActivityDefinitionParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *ActivityDefinitionParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *ActivityDefinitionParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *ActivityDefinitionParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *ActivityDefinitionParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *ActivityDefinitionParser) GetStatus() any {
	return p.data["status"]
}

// GetDescription returns the description element.
func (p *ActivityDefinitionParser) GetDescription() any {
	return p.data["description"]
}

// GetKind returns the kind element.
func (p *ActivityDefinitionParser) GetKind() any {
	return p.data["kind"]
}

// GetCode returns the code element.
func (p *ActivityDefinitionParser) GetCode() any {
	return p.data["code"]
}

// GetAuthor returns the author element.
func (p *ActivityDefinitionParser) GetAuthor() any {
	return p.data["author"]
}

// GetTiming returns the timing element.
func (p *ActivityDefinitionParser) GetTiming() any {
	return p.data["timing"]
}
