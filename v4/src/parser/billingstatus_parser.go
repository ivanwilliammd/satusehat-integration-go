package parser

// BillingStatusParser extracts fields from a FHIR BillingStatus resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/BillingStatus.json, copied unchanged to
// src/parser/testdata/BillingStatus.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type BillingStatusParser struct {
	payload map[string]any
}

// NewBillingStatusParser creates a new BillingStatusParser instance.
func NewBillingStatusParser(payload map[string]any) *BillingStatusParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &BillingStatusParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *BillingStatusParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *BillingStatusParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetId returns the id element unchanged.
func (p *BillingStatusParser) GetId() any {
	return p.payload["id"]
}

// GetIdentifier returns the identifier element unchanged.
func (p *BillingStatusParser) GetIdentifier() any {
	return p.payload["identifier"]
}

// GetStatus returns the status element unchanged.
func (p *BillingStatusParser) GetStatus() any {
	return p.payload["status"]
}

// GetInsurer returns the insurer element unchanged.
func (p *BillingStatusParser) GetInsurer() any {
	return p.payload["insurer"]
}

// GetRecipient returns the recipient element unchanged.
func (p *BillingStatusParser) GetRecipient() any {
	return p.payload["recipient"]
}

// GetSubject returns the subject element unchanged.
func (p *BillingStatusParser) GetSubject() any {
	return p.payload["subject"]
}

// GetRequest returns the request element unchanged.
func (p *BillingStatusParser) GetRequest() any {
	return p.payload["request"]
}
