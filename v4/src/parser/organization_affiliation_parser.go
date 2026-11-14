package parser

// OrganizationAffiliationParser extracts fields from a FHIR OrganizationAffiliation resource map.
//
// The element list is derived from PayloadBuilderOrganizationAffiliation in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type OrganizationAffiliationParser struct {
	data map[string]any
}

// NewOrganizationAffiliationParser creates a parser for the given OrganizationAffiliation payload.
func NewOrganizationAffiliationParser(data map[string]any) *OrganizationAffiliationParser {
	if data == nil {
		data = map[string]any{}
	}
	return &OrganizationAffiliationParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *OrganizationAffiliationParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *OrganizationAffiliationParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *OrganizationAffiliationParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *OrganizationAffiliationParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *OrganizationAffiliationParser) GetStatus() any {
	return p.data["status"]
}

// GetOrganization returns the organization element.
func (p *OrganizationAffiliationParser) GetOrganization() any {
	return p.data["organization"]
}

// GetParticipatingOrganization returns the participatingOrganization element.
func (p *OrganizationAffiliationParser) GetParticipatingOrganization() any {
	return p.data["participatingOrganization"]
}

// GetNetwork returns the network element.
func (p *OrganizationAffiliationParser) GetNetwork() any {
	return p.data["network"]
}

// GetCode returns the code element.
func (p *OrganizationAffiliationParser) GetCode() any {
	return p.data["code"]
}
