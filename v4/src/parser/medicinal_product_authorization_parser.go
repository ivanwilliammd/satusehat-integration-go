package parser

// MedicinalProductAuthorizationParser extracts fields from a FHIR MedicinalProductAuthorization resource map.
//
// The element list is derived from PayloadBuilderMedicinalProductAuthorization in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type MedicinalProductAuthorizationParser struct {
	data map[string]any
}

// NewMedicinalProductAuthorizationParser creates a parser for the given MedicinalProductAuthorization payload.
func NewMedicinalProductAuthorizationParser(data map[string]any) *MedicinalProductAuthorizationParser {
	if data == nil {
		data = map[string]any{}
	}
	return &MedicinalProductAuthorizationParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *MedicinalProductAuthorizationParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *MedicinalProductAuthorizationParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *MedicinalProductAuthorizationParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *MedicinalProductAuthorizationParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *MedicinalProductAuthorizationParser) GetStatus() any {
	return p.data["status"]
}

// GetCountry returns the country element.
func (p *MedicinalProductAuthorizationParser) GetCountry() any {
	return p.data["country"]
}

// GetRegulator returns the regulator element.
func (p *MedicinalProductAuthorizationParser) GetRegulator() any {
	return p.data["regulator"]
}

// GetValidityPeriod returns the validityPeriod element.
func (p *MedicinalProductAuthorizationParser) GetValidityPeriod() any {
	return p.data["validityPeriod"]
}
