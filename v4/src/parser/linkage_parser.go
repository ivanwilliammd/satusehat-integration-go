package parser

// LinkageParser extracts fields from a FHIR Linkage resource map.
//
// The element list is derived from PayloadBuilderLinkage in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type LinkageParser struct {
	data map[string]any
}

// NewLinkageParser creates a parser for the given Linkage payload.
func NewLinkageParser(data map[string]any) *LinkageParser {
	if data == nil {
		data = map[string]any{}
	}
	return &LinkageParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *LinkageParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *LinkageParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *LinkageParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *LinkageParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *LinkageParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *LinkageParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *LinkageParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *LinkageParser) GetEncounter() any {
	return p.data["encounter"]
}
