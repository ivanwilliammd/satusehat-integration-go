package parser

// FlagParser extracts fields from a FHIR Flag resource map.
//
// The element list is derived from PayloadBuilderFlag in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type FlagParser struct {
	data map[string]any
}

// NewFlagParser creates a parser for the given Flag payload.
func NewFlagParser(data map[string]any) *FlagParser {
	if data == nil {
		data = map[string]any{}
	}
	return &FlagParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *FlagParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *FlagParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *FlagParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *FlagParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *FlagParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *FlagParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *FlagParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *FlagParser) GetEncounter() any {
	return p.data["encounter"]
}
