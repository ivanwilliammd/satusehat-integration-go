package parser

// BodyStructureParser extracts fields from a FHIR BodyStructure resource map.
//
// The element list is derived from PayloadBuilderBodyStructure in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type BodyStructureParser struct {
	data map[string]any
}

// NewBodyStructureParser creates a parser for the given BodyStructure payload.
func NewBodyStructureParser(data map[string]any) *BodyStructureParser {
	if data == nil {
		data = map[string]any{}
	}
	return &BodyStructureParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *BodyStructureParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *BodyStructureParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *BodyStructureParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *BodyStructureParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *BodyStructureParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *BodyStructureParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *BodyStructureParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *BodyStructureParser) GetEncounter() any {
	return p.data["encounter"]
}
