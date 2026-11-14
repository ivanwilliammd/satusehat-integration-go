package parser

// PurificationDecisionParser extracts fields from a FHIR PurificationDecision resource map.
//
// The element list is derived from PayloadBuilderPurificationDecision in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type PurificationDecisionParser struct {
	data map[string]any
}

// NewPurificationDecisionParser creates a parser for the given PurificationDecision payload.
func NewPurificationDecisionParser(data map[string]any) *PurificationDecisionParser {
	if data == nil {
		data = map[string]any{}
	}
	return &PurificationDecisionParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *PurificationDecisionParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *PurificationDecisionParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *PurificationDecisionParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *PurificationDecisionParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *PurificationDecisionParser) GetStatus() any {
	return p.data["status"]
}

// GetInsurer returns the insurer element.
func (p *PurificationDecisionParser) GetInsurer() any {
	return p.data["insurer"]
}

// GetProvider returns the provider element.
func (p *PurificationDecisionParser) GetProvider() any {
	return p.data["provider"]
}

// GetClaimResponse returns the claimResponse element.
func (p *PurificationDecisionParser) GetClaimResponse() any {
	return p.data["claimResponse"]
}

// GetCreated returns the created element.
func (p *PurificationDecisionParser) GetCreated() any {
	return p.data["created"]
}
