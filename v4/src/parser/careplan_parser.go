package parser

// CarePlanParser extracts fields from a FHIR CarePlan resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/CarePlan.json, copied unchanged to
// src/parser/testdata/CarePlan.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type CarePlanParser struct {
	payload map[string]any
}

// NewCarePlanParser creates a new CarePlanParser instance.
func NewCarePlanParser(payload map[string]any) *CarePlanParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &CarePlanParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *CarePlanParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *CarePlanParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetTitle returns the title element unchanged.
func (p *CarePlanParser) GetTitle() any {
	return p.payload["title"]
}

// GetDescription returns the description element unchanged.
func (p *CarePlanParser) GetDescription() any {
	return p.payload["description"]
}

// GetCategory returns the category element unchanged.
func (p *CarePlanParser) GetCategory() any {
	return p.payload["category"]
}

// GetStatus returns the status element unchanged.
func (p *CarePlanParser) GetStatus() any {
	return p.payload["status"]
}

// GetIntent returns the intent element unchanged.
func (p *CarePlanParser) GetIntent() any {
	return p.payload["intent"]
}

// GetSubject returns the subject element unchanged.
func (p *CarePlanParser) GetSubject() any {
	return p.payload["subject"]
}

// GetEncounter returns the encounter element unchanged.
func (p *CarePlanParser) GetEncounter() any {
	return p.payload["encounter"]
}

// GetPeriod returns the period element unchanged.
func (p *CarePlanParser) GetPeriod() any {
	return p.payload["period"]
}

// GetCreated returns the created element unchanged.
func (p *CarePlanParser) GetCreated() any {
	return p.payload["created"]
}

// GetAuthor returns the author element unchanged.
func (p *CarePlanParser) GetAuthor() any {
	return p.payload["author"]
}

// GetActivity returns the activity element unchanged.
func (p *CarePlanParser) GetActivity() any {
	return p.payload["activity"]
}
