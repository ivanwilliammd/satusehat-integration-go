package parser

// ChargeItemResponseParser extracts fields from a FHIR ChargeItemResponse resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/ChargeItemResponse.json, copied unchanged to
// src/parser/testdata/ChargeItemResponse.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type ChargeItemResponseParser struct {
	payload map[string]any
}

// NewChargeItemResponseParser creates a new ChargeItemResponseParser instance.
func NewChargeItemResponseParser(payload map[string]any) *ChargeItemResponseParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &ChargeItemResponseParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *ChargeItemResponseParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *ChargeItemResponseParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetId returns the id element unchanged.
func (p *ChargeItemResponseParser) GetId() any {
	return p.payload["id"]
}

// GetIdentifier returns the identifier element unchanged.
func (p *ChargeItemResponseParser) GetIdentifier() any {
	return p.payload["identifier"]
}

// GetStatus returns the status element unchanged.
func (p *ChargeItemResponseParser) GetStatus() any {
	return p.payload["status"]
}

// GetChargeItem returns the chargeItem element unchanged.
func (p *ChargeItemResponseParser) GetChargeItem() any {
	return p.payload["chargeItem"]
}

// GetCoverage returns the coverage element unchanged.
func (p *ChargeItemResponseParser) GetCoverage() any {
	return p.payload["coverage"]
}

// GetSubject returns the subject element unchanged.
func (p *ChargeItemResponseParser) GetSubject() any {
	return p.payload["subject"]
}

// GetEncounter returns the encounter element unchanged.
func (p *ChargeItemResponseParser) GetEncounter() any {
	return p.payload["encounter"]
}

// GetPayor returns the payor element unchanged.
func (p *ChargeItemResponseParser) GetPayor() any {
	return p.payload["payor"]
}

// GetQuantity returns the quantity element unchanged.
func (p *ChargeItemResponseParser) GetQuantity() any {
	return p.payload["quantity"]
}

// GetUnitPriceOverride returns the unitPriceOverride element unchanged.
func (p *ChargeItemResponseParser) GetUnitPriceOverride() any {
	return p.payload["unitPriceOverride"]
}

// GetFactorOverride returns the factorOverride element unchanged.
func (p *ChargeItemResponseParser) GetFactorOverride() any {
	return p.payload["factorOverride"]
}

// GetOverrideReason returns the overrideReason element unchanged.
func (p *ChargeItemResponseParser) GetOverrideReason() any {
	return p.payload["overrideReason"]
}

// GetTotal returns the total element unchanged.
func (p *ChargeItemResponseParser) GetTotal() any {
	return p.payload["total"]
}

// GetExcess returns the excess element unchanged.
func (p *ChargeItemResponseParser) GetExcess() any {
	return p.payload["excess"]
}
