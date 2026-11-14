package parser

// MessageDefinitionParser extracts fields from a FHIR MessageDefinition resource map.
//
// The element list is derived from PayloadBuilderMessageDefinition in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type MessageDefinitionParser struct {
	data map[string]any
}

// NewMessageDefinitionParser creates a parser for the given MessageDefinition payload.
func NewMessageDefinitionParser(data map[string]any) *MessageDefinitionParser {
	if data == nil {
		data = map[string]any{}
	}
	return &MessageDefinitionParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *MessageDefinitionParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *MessageDefinitionParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *MessageDefinitionParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *MessageDefinitionParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *MessageDefinitionParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *MessageDefinitionParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *MessageDefinitionParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *MessageDefinitionParser) GetEncounter() any {
	return p.data["encounter"]
}
