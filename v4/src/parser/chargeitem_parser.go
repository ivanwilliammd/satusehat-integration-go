package parser

// ChargeItemParser extracts fields from a FHIR ChargeItem resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/ChargeItem.json, copied unchanged to
// src/parser/testdata/ChargeItem.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type ChargeItemParser struct {
	payload map[string]any
}

// NewChargeItemParser creates a new ChargeItemParser instance.
func NewChargeItemParser(payload map[string]any) *ChargeItemParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &ChargeItemParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *ChargeItemParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *ChargeItemParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetStatus returns the status element unchanged.
func (p *ChargeItemParser) GetStatus() any {
	return p.payload["status"]
}

// GetCode returns the code element unchanged.
func (p *ChargeItemParser) GetCode() any {
	return p.payload["code"]
}

// GetSubject returns the subject element unchanged.
func (p *ChargeItemParser) GetSubject() any {
	return p.payload["subject"]
}

// GetContext returns the context element unchanged.
func (p *ChargeItemParser) GetContext() any {
	return p.payload["context"]
}

// GetOccurrencePeriod returns the occurrencePeriod element unchanged.
func (p *ChargeItemParser) GetOccurrencePeriod() any {
	return p.payload["occurrencePeriod"]
}

// GetPerformer returns the performer element unchanged.
func (p *ChargeItemParser) GetPerformer() any {
	return p.payload["performer"]
}

// GetQuantity returns the quantity element unchanged.
func (p *ChargeItemParser) GetQuantity() any {
	return p.payload["quantity"]
}

// GetAccount returns the account element unchanged.
func (p *ChargeItemParser) GetAccount() any {
	return p.payload["account"]
}

// GetExtension returns the extension element unchanged.
func (p *ChargeItemParser) GetExtension() any {
	return p.payload["extension"]
}
