package parser

// ConditionParser extracts fields from a FHIR Condition resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/Condition.json, copied unchanged to
// src/parser/testdata/Condition.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type ConditionParser struct {
	payload map[string]any
}

// NewConditionParser creates a new ConditionParser instance.
func NewConditionParser(payload map[string]any) *ConditionParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &ConditionParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *ConditionParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *ConditionParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetClinicalStatus returns the clinicalStatus element unchanged.
func (p *ConditionParser) GetClinicalStatus() any {
	return p.payload["clinicalStatus"]
}

// GetVerificationStatus returns the verificationStatus element unchanged.
func (p *ConditionParser) GetVerificationStatus() any {
	return p.payload["verificationStatus"]
}

// GetCategory returns the category element unchanged.
func (p *ConditionParser) GetCategory() any {
	return p.payload["category"]
}

// GetCode returns the code element unchanged.
func (p *ConditionParser) GetCode() any {
	return p.payload["code"]
}

// GetStage returns the stage element unchanged.
func (p *ConditionParser) GetStage() any {
	return p.payload["stage"]
}

// GetSubject returns the subject element unchanged.
func (p *ConditionParser) GetSubject() any {
	return p.payload["subject"]
}

// GetEncounter returns the encounter element unchanged.
func (p *ConditionParser) GetEncounter() any {
	return p.payload["encounter"]
}

// GetRecordedDate returns the recordedDate element unchanged.
func (p *ConditionParser) GetRecordedDate() any {
	return p.payload["recordedDate"]
}
