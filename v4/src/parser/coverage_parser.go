package parser

// CoverageParser extracts fields from a FHIR Coverage resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/Coverage.json, copied unchanged to
// src/parser/testdata/Coverage.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type CoverageParser struct {
	payload map[string]any
}

// NewCoverageParser creates a new CoverageParser instance.
func NewCoverageParser(payload map[string]any) *CoverageParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &CoverageParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *CoverageParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *CoverageParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetIdentifier returns the identifier element unchanged.
func (p *CoverageParser) GetIdentifier() any {
	return p.payload["identifier"]
}

// GetStatus returns the status element unchanged.
func (p *CoverageParser) GetStatus() any {
	return p.payload["status"]
}

// GetType returns the type element unchanged.
func (p *CoverageParser) GetType() any {
	return p.payload["type"]
}

// GetSubscriberId returns the subscriberId element unchanged.
func (p *CoverageParser) GetSubscriberId() any {
	return p.payload["subscriberId"]
}

// GetBeneficiary returns the beneficiary element unchanged.
func (p *CoverageParser) GetBeneficiary() any {
	return p.payload["beneficiary"]
}

// GetPayor returns the payor element unchanged.
func (p *CoverageParser) GetPayor() any {
	return p.payload["payor"]
}

// GetClass returns the class element unchanged.
func (p *CoverageParser) GetClass() any {
	return p.payload["class"]
}

// GetExtension returns the extension element unchanged.
func (p *CoverageParser) GetExtension() any {
	return p.payload["extension"]
}
