package parser

// ResearchDefinitionParser extracts fields from a FHIR ResearchDefinition resource map.
//
// The element list is derived from PayloadBuilderResearchDefinition in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type ResearchDefinitionParser struct {
	data map[string]any
}

// NewResearchDefinitionParser creates a parser for the given ResearchDefinition payload.
func NewResearchDefinitionParser(data map[string]any) *ResearchDefinitionParser {
	if data == nil {
		data = map[string]any{}
	}
	return &ResearchDefinitionParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *ResearchDefinitionParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *ResearchDefinitionParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *ResearchDefinitionParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *ResearchDefinitionParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *ResearchDefinitionParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *ResearchDefinitionParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *ResearchDefinitionParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *ResearchDefinitionParser) GetEncounter() any {
	return p.data["encounter"]
}
