package parser

// RiskEvidenceSynthesisParser extracts fields from a FHIR RiskEvidenceSynthesis resource map.
//
// The element list is derived from PayloadBuilderRiskEvidenceSynthesis in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type RiskEvidenceSynthesisParser struct {
	data map[string]any
}

// NewRiskEvidenceSynthesisParser creates a parser for the given RiskEvidenceSynthesis payload.
func NewRiskEvidenceSynthesisParser(data map[string]any) *RiskEvidenceSynthesisParser {
	if data == nil {
		data = map[string]any{}
	}
	return &RiskEvidenceSynthesisParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *RiskEvidenceSynthesisParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *RiskEvidenceSynthesisParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *RiskEvidenceSynthesisParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *RiskEvidenceSynthesisParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *RiskEvidenceSynthesisParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *RiskEvidenceSynthesisParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *RiskEvidenceSynthesisParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *RiskEvidenceSynthesisParser) GetEncounter() any {
	return p.data["encounter"]
}
