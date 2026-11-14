package parser

// ResourceGuideParser extracts fields from a FHIR ResourceGuide resource map.
//
// The element list is derived from PayloadBuilderResourceGuide in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type ResourceGuideParser struct {
	data map[string]any
}

// NewResourceGuideParser creates a parser for the given ResourceGuide payload.
func NewResourceGuideParser(data map[string]any) *ResourceGuideParser {
	if data == nil {
		data = map[string]any{}
	}
	return &ResourceGuideParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *ResourceGuideParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *ResourceGuideParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *ResourceGuideParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *ResourceGuideParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *ResourceGuideParser) GetStatus() any {
	return p.data["status"]
}

// GetName returns the name element.
func (p *ResourceGuideParser) GetName() any {
	return p.data["name"]
}

// GetDescription returns the description element.
func (p *ResourceGuideParser) GetDescription() any {
	return p.data["description"]
}

// GetVersion returns the version element.
func (p *ResourceGuideParser) GetVersion() any {
	return p.data["version"]
}

// GetPublisher returns the publisher element.
func (p *ResourceGuideParser) GetPublisher() any {
	return p.data["publisher"]
}
