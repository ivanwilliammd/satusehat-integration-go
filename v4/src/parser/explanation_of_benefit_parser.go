package parser

// ExplanationOfBenefitParser extracts fields from a FHIR ExplanationOfBenefit resource map.
//
// The element list is derived from PayloadBuilderExplanationOfBenefit in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type ExplanationOfBenefitParser struct {
	data map[string]any
}

// NewExplanationOfBenefitParser creates a parser for the given ExplanationOfBenefit payload.
func NewExplanationOfBenefitParser(data map[string]any) *ExplanationOfBenefitParser {
	if data == nil {
		data = map[string]any{}
	}
	return &ExplanationOfBenefitParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *ExplanationOfBenefitParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *ExplanationOfBenefitParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *ExplanationOfBenefitParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *ExplanationOfBenefitParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *ExplanationOfBenefitParser) GetStatus() any {
	return p.data["status"]
}

// GetType returns the type element.
func (p *ExplanationOfBenefitParser) GetType() any {
	return p.data["type"]
}

// GetUse returns the use element.
func (p *ExplanationOfBenefitParser) GetUse() any {
	return p.data["use"]
}

// GetPatient returns the patient element.
func (p *ExplanationOfBenefitParser) GetPatient() any {
	return p.data["patient"]
}

// GetCreated returns the created element.
func (p *ExplanationOfBenefitParser) GetCreated() any {
	return p.data["created"]
}

// GetProvider returns the provider element.
func (p *ExplanationOfBenefitParser) GetProvider() any {
	return p.data["provider"]
}
