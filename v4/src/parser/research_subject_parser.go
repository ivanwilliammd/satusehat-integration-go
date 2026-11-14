package parser

// ResearchSubjectParser extracts fields from a FHIR ResearchSubject resource map.
//
// The element list is derived from PayloadBuilderResearchSubject in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type ResearchSubjectParser struct {
	data map[string]any
}

// NewResearchSubjectParser creates a parser for the given ResearchSubject payload.
func NewResearchSubjectParser(data map[string]any) *ResearchSubjectParser {
	if data == nil {
		data = map[string]any{}
	}
	return &ResearchSubjectParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *ResearchSubjectParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *ResearchSubjectParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *ResearchSubjectParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *ResearchSubjectParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *ResearchSubjectParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *ResearchSubjectParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *ResearchSubjectParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *ResearchSubjectParser) GetEncounter() any {
	return p.data["encounter"]
}
