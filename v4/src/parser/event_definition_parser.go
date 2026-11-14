package parser

// EventDefinitionParser extracts fields from a FHIR EventDefinition resource map.
//
// The element list is derived from PayloadBuilderEventDefinition in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type EventDefinitionParser struct {
	data map[string]any
}

// NewEventDefinitionParser creates a parser for the given EventDefinition payload.
func NewEventDefinitionParser(data map[string]any) *EventDefinitionParser {
	if data == nil {
		data = map[string]any{}
	}
	return &EventDefinitionParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *EventDefinitionParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *EventDefinitionParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *EventDefinitionParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *EventDefinitionParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *EventDefinitionParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *EventDefinitionParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *EventDefinitionParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *EventDefinitionParser) GetEncounter() any {
	return p.data["encounter"]
}
