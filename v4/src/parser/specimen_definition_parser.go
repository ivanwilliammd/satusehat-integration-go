package parser

// SpecimenDefinitionParser extracts fields from a FHIR SpecimenDefinition resource map.
//
// The element list is derived from PayloadBuilderSpecimenDefinition in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type SpecimenDefinitionParser struct {
	data map[string]any
}

// NewSpecimenDefinitionParser creates a parser for the given SpecimenDefinition payload.
func NewSpecimenDefinitionParser(data map[string]any) *SpecimenDefinitionParser {
	if data == nil {
		data = map[string]any{}
	}
	return &SpecimenDefinitionParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *SpecimenDefinitionParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *SpecimenDefinitionParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *SpecimenDefinitionParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *SpecimenDefinitionParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *SpecimenDefinitionParser) GetStatus() any {
	return p.data["status"]
}

// GetType returns the type element.
func (p *SpecimenDefinitionParser) GetType() any {
	return p.data["type"]
}

// GetCollection returns the collection element.
func (p *SpecimenDefinitionParser) GetCollection() any {
	return p.data["collection"]
}

// GetPreparation returns the preparation element.
func (p *SpecimenDefinitionParser) GetPreparation() any {
	return p.data["preparation"]
}
