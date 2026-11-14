package parser

// CareTeamParser extracts fields from a FHIR CareTeam resource map.
//
// The element list is derived from PayloadBuilderCareTeam in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type CareTeamParser struct {
	data map[string]any
}

// NewCareTeamParser creates a parser for the given CareTeam payload.
func NewCareTeamParser(data map[string]any) *CareTeamParser {
	if data == nil {
		data = map[string]any{}
	}
	return &CareTeamParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *CareTeamParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *CareTeamParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *CareTeamParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *CareTeamParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *CareTeamParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *CareTeamParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *CareTeamParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *CareTeamParser) GetEncounter() any {
	return p.data["encounter"]
}
