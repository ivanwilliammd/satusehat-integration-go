package parser

// MessageHeaderParser extracts fields from a FHIR MessageHeader resource map.
//
// The element list is derived from PayloadBuilderMessageHeader in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type MessageHeaderParser struct {
	data map[string]any
}

// NewMessageHeaderParser creates a parser for the given MessageHeader payload.
func NewMessageHeaderParser(data map[string]any) *MessageHeaderParser {
	if data == nil {
		data = map[string]any{}
	}
	return &MessageHeaderParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *MessageHeaderParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *MessageHeaderParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *MessageHeaderParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *MessageHeaderParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *MessageHeaderParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *MessageHeaderParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *MessageHeaderParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *MessageHeaderParser) GetEncounter() any {
	return p.data["encounter"]
}
