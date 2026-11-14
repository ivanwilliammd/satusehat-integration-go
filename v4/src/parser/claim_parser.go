package parser

// ClaimParser extracts fields from a FHIR Claim resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/Claim.json, copied unchanged to
// src/parser/testdata/Claim.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type ClaimParser struct {
	payload map[string]any
}

// NewClaimParser creates a new ClaimParser instance.
func NewClaimParser(payload map[string]any) *ClaimParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &ClaimParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *ClaimParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *ClaimParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetIdentifier returns the identifier element unchanged.
func (p *ClaimParser) GetIdentifier() any {
	return p.payload["identifier"]
}

// GetStatus returns the status element unchanged.
func (p *ClaimParser) GetStatus() any {
	return p.payload["status"]
}

// GetType returns the type element unchanged.
func (p *ClaimParser) GetType() any {
	return p.payload["type"]
}

// GetUse returns the use element unchanged.
func (p *ClaimParser) GetUse() any {
	return p.payload["use"]
}

// GetPatient returns the patient element unchanged.
func (p *ClaimParser) GetPatient() any {
	return p.payload["patient"]
}

// GetBillablePeriod returns the billablePeriod element unchanged.
func (p *ClaimParser) GetBillablePeriod() any {
	return p.payload["billablePeriod"]
}

// GetCreated returns the created element unchanged.
func (p *ClaimParser) GetCreated() any {
	return p.payload["created"]
}

// GetEnterer returns the enterer element unchanged.
func (p *ClaimParser) GetEnterer() any {
	return p.payload["enterer"]
}

// GetInsurer returns the insurer element unchanged.
func (p *ClaimParser) GetInsurer() any {
	return p.payload["insurer"]
}

// GetProvider returns the provider element unchanged.
func (p *ClaimParser) GetProvider() any {
	return p.payload["provider"]
}

// GetPriority returns the priority element unchanged.
func (p *ClaimParser) GetPriority() any {
	return p.payload["priority"]
}

// GetPayee returns the payee element unchanged.
func (p *ClaimParser) GetPayee() any {
	return p.payload["payee"]
}

// GetFacility returns the facility element unchanged.
func (p *ClaimParser) GetFacility() any {
	return p.payload["facility"]
}

// GetSupportingInfo returns the supportingInfo element unchanged.
func (p *ClaimParser) GetSupportingInfo() any {
	return p.payload["supportingInfo"]
}

// GetDiagnosis returns the diagnosis element unchanged.
func (p *ClaimParser) GetDiagnosis() any {
	return p.payload["diagnosis"]
}

// GetProcedure returns the procedure element unchanged.
func (p *ClaimParser) GetProcedure() any {
	return p.payload["procedure"]
}

// GetInsurance returns the insurance element unchanged.
func (p *ClaimParser) GetInsurance() any {
	return p.payload["insurance"]
}

// GetTotal returns the total element unchanged.
func (p *ClaimParser) GetTotal() any {
	return p.payload["total"]
}

// GetExtension returns the extension element unchanged.
func (p *ClaimParser) GetExtension() any {
	return p.payload["extension"]
}
