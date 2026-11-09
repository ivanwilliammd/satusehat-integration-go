package parser

// OrganizationParser extracts fields from a FHIR Organization resource map.
type OrganizationParser struct {
	organization map[string]any
}

// NewOrganizationParser creates a new OrganizationParser instance.
func NewOrganizationParser(organization map[string]any) *OrganizationParser {
	if organization == nil {
		organization = map[string]any{}
	}
	return &OrganizationParser{organization: organization}
}

func (p *OrganizationParser) GetIdentifiers() any {
	return p.organization["identifier"]
}

func (p *OrganizationParser) GetName() string {
	return strVal(p.organization["name"])
}

func (p *OrganizationParser) GetOperationalStatus() any {
	return p.organization["extension"]
}

func (p *OrganizationParser) GetPartOf() any {
	return p.organization["partOf"]
}

func (p *OrganizationParser) GetPartOfReference() string {
	return removePrefix(nestedStr(p.organization, "partOf", "reference"), "Organization/")
}

func (p *OrganizationParser) GetType() any {
	return p.organization["type"]
}

func (p *OrganizationParser) GetTelecom() any {
	return p.organization["telecom"]
}

func (p *OrganizationParser) GetAddress() any {
	return p.organization["address"]
}
