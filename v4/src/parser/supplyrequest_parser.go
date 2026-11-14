package parser

// SupplyRequestParser extracts fields from a FHIR SupplyRequest resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/SupplyRequest.json, copied unchanged to
// src/parser/testdata/SupplyRequest.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type SupplyRequestParser struct {
	payload map[string]any
}

// NewSupplyRequestParser creates a new SupplyRequestParser instance.
func NewSupplyRequestParser(payload map[string]any) *SupplyRequestParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &SupplyRequestParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *SupplyRequestParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *SupplyRequestParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetIdentifier returns the identifier element unchanged.
func (p *SupplyRequestParser) GetIdentifier() any {
	return p.payload["identifier"]
}

// GetStatus returns the status element unchanged.
func (p *SupplyRequestParser) GetStatus() any {
	return p.payload["status"]
}

// GetItemReference returns the itemReference element unchanged.
func (p *SupplyRequestParser) GetItemReference() any {
	return p.payload["itemReference"]
}

// GetQuantity returns the quantity element unchanged.
func (p *SupplyRequestParser) GetQuantity() any {
	return p.payload["quantity"]
}

// GetAuthoredOn returns the authoredOn element unchanged.
func (p *SupplyRequestParser) GetAuthoredOn() any {
	return p.payload["authoredOn"]
}

// GetRequester returns the requester element unchanged.
func (p *SupplyRequestParser) GetRequester() any {
	return p.payload["requester"]
}

// GetSupplier returns the supplier element unchanged.
func (p *SupplyRequestParser) GetSupplier() any {
	return p.payload["supplier"]
}

// GetReasonCode returns the reasonCode element unchanged.
func (p *SupplyRequestParser) GetReasonCode() any {
	return p.payload["reasonCode"]
}
