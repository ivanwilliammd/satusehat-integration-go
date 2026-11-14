package parser

// AccountParser extracts fields from a FHIR Account resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/Account.json, copied unchanged to
// src/parser/testdata/Account.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type AccountParser struct {
	payload map[string]any
}

// NewAccountParser creates a new AccountParser instance.
func NewAccountParser(payload map[string]any) *AccountParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &AccountParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *AccountParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *AccountParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetStatus returns the status element unchanged.
func (p *AccountParser) GetStatus() any {
	return p.payload["status"]
}

// GetType returns the type element unchanged.
func (p *AccountParser) GetType() any {
	return p.payload["type"]
}

// GetName returns the name element unchanged.
func (p *AccountParser) GetName() any {
	return p.payload["name"]
}

// GetSubject returns the subject element unchanged.
func (p *AccountParser) GetSubject() any {
	return p.payload["subject"]
}

// GetServicePeriod returns the servicePeriod element unchanged.
func (p *AccountParser) GetServicePeriod() any {
	return p.payload["servicePeriod"]
}

// GetOwner returns the owner element unchanged.
func (p *AccountParser) GetOwner() any {
	return p.payload["owner"]
}

// GetGuarantor returns the guarantor element unchanged.
func (p *AccountParser) GetGuarantor() any {
	return p.payload["guarantor"]
}

// GetDescription returns the description element unchanged.
func (p *AccountParser) GetDescription() any {
	return p.payload["description"]
}
