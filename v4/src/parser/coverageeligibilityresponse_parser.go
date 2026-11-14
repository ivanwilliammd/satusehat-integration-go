package parser

// CoverageEligibilityResponseParser extracts fields from a FHIR CoverageEligibilityResponse resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/CoverageEligibilityResponse.json, copied unchanged to
// src/parser/testdata/CoverageEligibilityResponse.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type CoverageEligibilityResponseParser struct {
	payload map[string]any
}

// NewCoverageEligibilityResponseParser creates a new CoverageEligibilityResponseParser instance.
func NewCoverageEligibilityResponseParser(payload map[string]any) *CoverageEligibilityResponseParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &CoverageEligibilityResponseParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *CoverageEligibilityResponseParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *CoverageEligibilityResponseParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetIdentifier returns the identifier element unchanged.
func (p *CoverageEligibilityResponseParser) GetIdentifier() any {
	return p.payload["identifier"]
}

// GetStatus returns the status element unchanged.
func (p *CoverageEligibilityResponseParser) GetStatus() any {
	return p.payload["status"]
}

// GetPurpose returns the purpose element unchanged.
func (p *CoverageEligibilityResponseParser) GetPurpose() any {
	return p.payload["purpose"]
}

// GetPatient returns the patient element unchanged.
func (p *CoverageEligibilityResponseParser) GetPatient() any {
	return p.payload["patient"]
}

// GetCreated returns the created element unchanged.
func (p *CoverageEligibilityResponseParser) GetCreated() any {
	return p.payload["created"]
}

// GetRequestor returns the requestor element unchanged.
func (p *CoverageEligibilityResponseParser) GetRequestor() any {
	return p.payload["requestor"]
}

// GetRequest returns the request element unchanged.
func (p *CoverageEligibilityResponseParser) GetRequest() any {
	return p.payload["request"]
}

// GetOutcome returns the outcome element unchanged.
func (p *CoverageEligibilityResponseParser) GetOutcome() any {
	return p.payload["outcome"]
}

// GetDisposition returns the disposition element unchanged.
func (p *CoverageEligibilityResponseParser) GetDisposition() any {
	return p.payload["disposition"]
}

// GetInsurer returns the insurer element unchanged.
func (p *CoverageEligibilityResponseParser) GetInsurer() any {
	return p.payload["insurer"]
}

// GetInsurance returns the insurance element unchanged.
func (p *CoverageEligibilityResponseParser) GetInsurance() any {
	return p.payload["insurance"]
}

// GetExtension returns the extension element unchanged.
func (p *CoverageEligibilityResponseParser) GetExtension() any {
	return p.payload["extension"]
}
