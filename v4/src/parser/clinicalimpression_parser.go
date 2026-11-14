package parser

// ClinicalImpressionParser extracts fields from a FHIR ClinicalImpression resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/ClinicalImpression.json, copied unchanged to
// src/parser/testdata/ClinicalImpression.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type ClinicalImpressionParser struct {
	payload map[string]any
}

// NewClinicalImpressionParser creates a new ClinicalImpressionParser instance.
func NewClinicalImpressionParser(payload map[string]any) *ClinicalImpressionParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &ClinicalImpressionParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *ClinicalImpressionParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *ClinicalImpressionParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetIdentifier returns the identifier element unchanged.
func (p *ClinicalImpressionParser) GetIdentifier() any {
	return p.payload["identifier"]
}

// GetStatus returns the status element unchanged.
func (p *ClinicalImpressionParser) GetStatus() any {
	return p.payload["status"]
}

// GetCode returns the code element unchanged.
func (p *ClinicalImpressionParser) GetCode() any {
	return p.payload["code"]
}

// GetDescription returns the description element unchanged.
func (p *ClinicalImpressionParser) GetDescription() any {
	return p.payload["description"]
}

// GetSubject returns the subject element unchanged.
func (p *ClinicalImpressionParser) GetSubject() any {
	return p.payload["subject"]
}

// GetEncounter returns the encounter element unchanged.
func (p *ClinicalImpressionParser) GetEncounter() any {
	return p.payload["encounter"]
}

// GetEffectiveDateTime returns the effectiveDateTime element unchanged.
func (p *ClinicalImpressionParser) GetEffectiveDateTime() any {
	return p.payload["effectiveDateTime"]
}

// GetDate returns the date element unchanged.
func (p *ClinicalImpressionParser) GetDate() any {
	return p.payload["date"]
}

// GetAssessor returns the assessor element unchanged.
func (p *ClinicalImpressionParser) GetAssessor() any {
	return p.payload["assessor"]
}

// GetProblem returns the problem element unchanged.
func (p *ClinicalImpressionParser) GetProblem() any {
	return p.payload["problem"]
}

// GetInvestigation returns the investigation element unchanged.
func (p *ClinicalImpressionParser) GetInvestigation() any {
	return p.payload["investigation"]
}

// GetSummary returns the summary element unchanged.
func (p *ClinicalImpressionParser) GetSummary() any {
	return p.payload["summary"]
}

// GetFinding returns the finding element unchanged.
func (p *ClinicalImpressionParser) GetFinding() any {
	return p.payload["finding"]
}

// GetPrognosisCodeableConcept returns the prognosisCodeableConcept element unchanged.
func (p *ClinicalImpressionParser) GetPrognosisCodeableConcept() any {
	return p.payload["prognosisCodeableConcept"]
}
