package parser

// GuidanceResponseParser extracts fields from a FHIR GuidanceResponse resource map.
//
// The element list is derived from PayloadBuilderGuidanceResponse in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type GuidanceResponseParser struct {
	data map[string]any
}

// NewGuidanceResponseParser creates a parser for the given GuidanceResponse payload.
func NewGuidanceResponseParser(data map[string]any) *GuidanceResponseParser {
	if data == nil {
		data = map[string]any{}
	}
	return &GuidanceResponseParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *GuidanceResponseParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *GuidanceResponseParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *GuidanceResponseParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *GuidanceResponseParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *GuidanceResponseParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *GuidanceResponseParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *GuidanceResponseParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *GuidanceResponseParser) GetEncounter() any {
	return p.data["encounter"]
}
