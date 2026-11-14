package parser

// DeviceDefinitionParser extracts fields from a FHIR DeviceDefinition resource map.
//
// The element list is derived from PayloadBuilderDeviceDefinition in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type DeviceDefinitionParser struct {
	data map[string]any
}

// NewDeviceDefinitionParser creates a parser for the given DeviceDefinition payload.
func NewDeviceDefinitionParser(data map[string]any) *DeviceDefinitionParser {
	if data == nil {
		data = map[string]any{}
	}
	return &DeviceDefinitionParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *DeviceDefinitionParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *DeviceDefinitionParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *DeviceDefinitionParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *DeviceDefinitionParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *DeviceDefinitionParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *DeviceDefinitionParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *DeviceDefinitionParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *DeviceDefinitionParser) GetEncounter() any {
	return p.data["encounter"]
}
