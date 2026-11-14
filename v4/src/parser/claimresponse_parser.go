package parser

// ClaimResponseParser extracts fields from a FHIR ClaimResponse resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/ClaimResponse.json, copied unchanged to
// src/parser/testdata/ClaimResponse.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type ClaimResponseParser struct {
	payload map[string]any
}

// NewClaimResponseParser creates a new ClaimResponseParser instance.
func NewClaimResponseParser(payload map[string]any) *ClaimResponseParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &ClaimResponseParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *ClaimResponseParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *ClaimResponseParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetStatus returns the status element unchanged.
func (p *ClaimResponseParser) GetStatus() any {
	return p.payload["status"]
}

// GetIdentifier returns the identifier element unchanged.
func (p *ClaimResponseParser) GetIdentifier() any {
	return p.payload["identifier"]
}

// GetType returns the type element unchanged.
func (p *ClaimResponseParser) GetType() any {
	return p.payload["type"]
}

// GetUse returns the use element unchanged.
func (p *ClaimResponseParser) GetUse() any {
	return p.payload["use"]
}

// GetPatient returns the patient element unchanged.
func (p *ClaimResponseParser) GetPatient() any {
	return p.payload["patient"]
}

// GetCreated returns the created element unchanged.
func (p *ClaimResponseParser) GetCreated() any {
	return p.payload["created"]
}

// GetInsurer returns the insurer element unchanged.
func (p *ClaimResponseParser) GetInsurer() any {
	return p.payload["insurer"]
}

// GetRequestor returns the requestor element unchanged.
func (p *ClaimResponseParser) GetRequestor() any {
	return p.payload["requestor"]
}

// GetRequest returns the request element unchanged.
func (p *ClaimResponseParser) GetRequest() any {
	return p.payload["request"]
}

// GetOutcome returns the outcome element unchanged.
func (p *ClaimResponseParser) GetOutcome() any {
	return p.payload["outcome"]
}

// GetPayeeType returns the payeeType element unchanged.
func (p *ClaimResponseParser) GetPayeeType() any {
	return p.payload["payeeType"]
}

// GetAdjudication returns the adjudication element unchanged.
func (p *ClaimResponseParser) GetAdjudication() any {
	return p.payload["adjudication"]
}

// GetItem returns the item element unchanged.
func (p *ClaimResponseParser) GetItem() any {
	return p.payload["item"]
}

// GetTotal returns the total element unchanged.
func (p *ClaimResponseParser) GetTotal() any {
	return p.payload["total"]
}

// GetPayment returns the payment element unchanged.
func (p *ClaimResponseParser) GetPayment() any {
	return p.payload["payment"]
}
