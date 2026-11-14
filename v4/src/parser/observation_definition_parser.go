package parser

// ObservationDefinitionParser extracts fields from a FHIR ObservationDefinition resource map.
//
// The element list is derived from PayloadBuilderObservationDefinition in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type ObservationDefinitionParser struct {
	data map[string]any
}

// NewObservationDefinitionParser creates a parser for the given ObservationDefinition payload.
func NewObservationDefinitionParser(data map[string]any) *ObservationDefinitionParser {
	if data == nil {
		data = map[string]any{}
	}
	return &ObservationDefinitionParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *ObservationDefinitionParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *ObservationDefinitionParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *ObservationDefinitionParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *ObservationDefinitionParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *ObservationDefinitionParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *ObservationDefinitionParser) GetCode() any {
	return p.data["code"]
}

// GetCategory returns the category element.
func (p *ObservationDefinitionParser) GetCategory() any {
	return p.data["category"]
}

// GetMethod returns the method element.
func (p *ObservationDefinitionParser) GetMethod() any {
	return p.data["method"]
}
