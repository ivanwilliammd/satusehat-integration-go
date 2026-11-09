package parser

// LocationParser extracts fields from a FHIR Location resource map.
type LocationParser struct {
	location map[string]any
}

// NewLocationParser creates a new LocationParser instance.
func NewLocationParser(location map[string]any) *LocationParser {
	if location == nil {
		location = map[string]any{}
	}
	return &LocationParser{location: location}
}

func (p *LocationParser) GetIdentifiers() any {
	return p.location["identifier"]
}

func (p *LocationParser) GetName() string {
	return strVal(p.location["name"])
}

func (p *LocationParser) GetStatus() string {
	return strVal(p.location["status"])
}

func (p *LocationParser) GetOperationalStatus() any {
	return p.location["operationalStatus"]
}

func (p *LocationParser) GetTelecom() any {
	return p.location["telecom"]
}

func (p *LocationParser) GetAddress() any {
	return p.location["address"]
}

func (p *LocationParser) GetPhysicalType() any {
	return p.location["physicalType"]
}

func (p *LocationParser) GetPosition() any {
	return p.location["position"]
}

func (p *LocationParser) GetManagingOrganization() any {
	return p.location["managingOrganization"]
}

func (p *LocationParser) GetManagingOrganizationReference() string {
	return removePrefix(nestedStr(p.location, "managingOrganization", "reference"), "Organization/")
}

func (p *LocationParser) GetPartOf() any {
	return p.location["partOf"]
}

func (p *LocationParser) GetPartOfReference() string {
	return removePrefix(nestedStr(p.location, "partOf", "reference"), "Location/")
}
