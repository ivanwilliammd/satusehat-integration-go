package parser

// FamilyMemberHistoryParser extracts fields from a FHIR FamilyMemberHistory resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/FamilyMemberHistory.json, copied unchanged to
// src/parser/testdata/FamilyMemberHistory.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type FamilyMemberHistoryParser struct {
	payload map[string]any
}

// NewFamilyMemberHistoryParser creates a new FamilyMemberHistoryParser instance.
func NewFamilyMemberHistoryParser(payload map[string]any) *FamilyMemberHistoryParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &FamilyMemberHistoryParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *FamilyMemberHistoryParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *FamilyMemberHistoryParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetStatus returns the status element unchanged.
func (p *FamilyMemberHistoryParser) GetStatus() any {
	return p.payload["status"]
}

// GetRelationship returns the relationship element unchanged.
func (p *FamilyMemberHistoryParser) GetRelationship() any {
	return p.payload["relationship"]
}

// GetDeceasedBoolean returns the deceasedBoolean element unchanged.
func (p *FamilyMemberHistoryParser) GetDeceasedBoolean() any {
	return p.payload["deceasedBoolean"]
}

// GetPatient returns the patient element unchanged.
func (p *FamilyMemberHistoryParser) GetPatient() any {
	return p.payload["patient"]
}

// GetDate returns the date element unchanged.
func (p *FamilyMemberHistoryParser) GetDate() any {
	return p.payload["date"]
}

// GetCondition returns the condition element unchanged.
func (p *FamilyMemberHistoryParser) GetCondition() any {
	return p.payload["condition"]
}
