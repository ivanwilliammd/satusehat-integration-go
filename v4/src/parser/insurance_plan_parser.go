package parser

// InsurancePlanParser extracts fields from a FHIR InsurancePlan resource map.
//
// The element list is derived from PayloadBuilderInsurancePlan in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type InsurancePlanParser struct {
	data map[string]any
}

// NewInsurancePlanParser creates a parser for the given InsurancePlan payload.
func NewInsurancePlanParser(data map[string]any) *InsurancePlanParser {
	if data == nil {
		data = map[string]any{}
	}
	return &InsurancePlanParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *InsurancePlanParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *InsurancePlanParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *InsurancePlanParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *InsurancePlanParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *InsurancePlanParser) GetStatus() any {
	return p.data["status"]
}

// GetName returns the name element.
func (p *InsurancePlanParser) GetName() any {
	return p.data["name"]
}

// GetType returns the type element.
func (p *InsurancePlanParser) GetType() any {
	return p.data["type"]
}

// GetAdministeredBy returns the administeredBy element.
func (p *InsurancePlanParser) GetAdministeredBy() any {
	return p.data["administeredBy"]
}

// GetCoverage returns the coverage element.
func (p *InsurancePlanParser) GetCoverage() any {
	return p.data["coverage"]
}
