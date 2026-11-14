package parser

// ResearchStudyParser extracts fields from a FHIR ResearchStudy resource map.
//
// The element list is derived from PayloadBuilderResearchStudy in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type ResearchStudyParser struct {
	data map[string]any
}

// NewResearchStudyParser creates a parser for the given ResearchStudy payload.
func NewResearchStudyParser(data map[string]any) *ResearchStudyParser {
	if data == nil {
		data = map[string]any{}
	}
	return &ResearchStudyParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *ResearchStudyParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *ResearchStudyParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *ResearchStudyParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *ResearchStudyParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *ResearchStudyParser) GetStatus() any {
	return p.data["status"]
}

// GetTitle returns the title element.
func (p *ResearchStudyParser) GetTitle() any {
	return p.data["title"]
}

// GetProtocol returns the protocol element.
func (p *ResearchStudyParser) GetProtocol() any {
	return p.data["protocol"]
}

// GetSponsor returns the sponsor element.
func (p *ResearchStudyParser) GetSponsor() any {
	return p.data["sponsor"]
}

// GetPhase returns the phase element.
func (p *ResearchStudyParser) GetPhase() any {
	return p.data["phase"]
}
