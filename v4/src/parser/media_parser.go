package parser

// MediaParser extracts fields from a FHIR Media resource map.
//
// The element list is derived from PayloadBuilderMedia in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type MediaParser struct {
	data map[string]any
}

// NewMediaParser creates a parser for the given Media payload.
func NewMediaParser(data map[string]any) *MediaParser {
	if data == nil {
		data = map[string]any{}
	}
	return &MediaParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *MediaParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *MediaParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *MediaParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *MediaParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *MediaParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *MediaParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *MediaParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *MediaParser) GetEncounter() any {
	return p.data["encounter"]
}
