package parser

// CoverageEligibilityRequestParser extracts fields from a FHIR CoverageEligibilityRequest resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/CoverageEligibilityRequest.json, copied unchanged to
// src/parser/testdata/CoverageEligibilityRequest.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type CoverageEligibilityRequestParser struct {
	payload map[string]any
}

// NewCoverageEligibilityRequestParser creates a new CoverageEligibilityRequestParser instance.
func NewCoverageEligibilityRequestParser(payload map[string]any) *CoverageEligibilityRequestParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &CoverageEligibilityRequestParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *CoverageEligibilityRequestParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *CoverageEligibilityRequestParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetIdentifier returns the identifier element unchanged.
func (p *CoverageEligibilityRequestParser) GetIdentifier() any {
	return p.payload["identifier"]
}

// GetStatus returns the status element unchanged.
func (p *CoverageEligibilityRequestParser) GetStatus() any {
	return p.payload["status"]
}

// GetPriority returns the priority element unchanged.
func (p *CoverageEligibilityRequestParser) GetPriority() any {
	return p.payload["priority"]
}

// GetPurpose returns the purpose element unchanged.
func (p *CoverageEligibilityRequestParser) GetPurpose() any {
	return p.payload["purpose"]
}

// GetPatient returns the patient element unchanged.
func (p *CoverageEligibilityRequestParser) GetPatient() any {
	return p.payload["patient"]
}

// GetServicedDate returns the servicedDate element unchanged.
func (p *CoverageEligibilityRequestParser) GetServicedDate() any {
	return p.payload["servicedDate"]
}

// GetCreated returns the created element unchanged.
func (p *CoverageEligibilityRequestParser) GetCreated() any {
	return p.payload["created"]
}

// GetEnterer returns the enterer element unchanged.
func (p *CoverageEligibilityRequestParser) GetEnterer() any {
	return p.payload["enterer"]
}

// GetProvider returns the provider element unchanged.
func (p *CoverageEligibilityRequestParser) GetProvider() any {
	return p.payload["provider"]
}

// GetInsurer returns the insurer element unchanged.
func (p *CoverageEligibilityRequestParser) GetInsurer() any {
	return p.payload["insurer"]
}

// GetFacility returns the facility element unchanged.
func (p *CoverageEligibilityRequestParser) GetFacility() any {
	return p.payload["facility"]
}

// GetInsurance returns the insurance element unchanged.
func (p *CoverageEligibilityRequestParser) GetInsurance() any {
	return p.payload["insurance"]
}

// GetSupportingInfo returns the supportingInfo element unchanged.
func (p *CoverageEligibilityRequestParser) GetSupportingInfo() any {
	return p.payload["supportingInfo"]
}

// GetContained returns the contained element unchanged.
func (p *CoverageEligibilityRequestParser) GetContained() any {
	return p.payload["contained"]
}
