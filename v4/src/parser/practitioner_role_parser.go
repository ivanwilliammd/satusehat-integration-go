package parser

// PractitionerRoleParser extracts fields from a FHIR PractitionerRole resource map.
//
// The element list is derived from PayloadBuilderPractitionerRole in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type PractitionerRoleParser struct {
	data map[string]any
}

// NewPractitionerRoleParser creates a parser for the given PractitionerRole payload.
func NewPractitionerRoleParser(data map[string]any) *PractitionerRoleParser {
	if data == nil {
		data = map[string]any{}
	}
	return &PractitionerRoleParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *PractitionerRoleParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *PractitionerRoleParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *PractitionerRoleParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *PractitionerRoleParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetActive returns the active element.
func (p *PractitionerRoleParser) GetActive() any {
	return p.data["active"]
}

// GetPractitioner returns the practitioner element.
func (p *PractitionerRoleParser) GetPractitioner() any {
	return p.data["practitioner"]
}

// GetOrganization returns the organization element.
func (p *PractitionerRoleParser) GetOrganization() any {
	return p.data["organization"]
}

// GetCode returns the code element.
func (p *PractitionerRoleParser) GetCode() any {
	return p.data["code"]
}

// GetSpecialty returns the specialty element.
func (p *PractitionerRoleParser) GetSpecialty() any {
	return p.data["specialty"]
}

// GetLocation returns the location element.
func (p *PractitionerRoleParser) GetLocation() any {
	return p.data["location"]
}

// GetHealthcareService returns the healthcareService element.
func (p *PractitionerRoleParser) GetHealthcareService() any {
	return p.data["healthcareService"]
}

// GetTelecom returns the telecom element.
func (p *PractitionerRoleParser) GetTelecom() any {
	return p.data["telecom"]
}

// GetAvailableTime returns the availableTime element.
func (p *PractitionerRoleParser) GetAvailableTime() any {
	return p.data["availableTime"]
}

// GetNotAvailable returns the notAvailable element.
func (p *PractitionerRoleParser) GetNotAvailable() any {
	return p.data["notAvailable"]
}

// GetEndpoint returns the endpoint element.
func (p *PractitionerRoleParser) GetEndpoint() any {
	return p.data["endpoint"]
}

// GetAvailabilityExceptions returns the availabilityExceptions element.
func (p *PractitionerRoleParser) GetAvailabilityExceptions() any {
	return p.data["availabilityExceptions"]
}

// GetExtension returns the extension element.
func (p *PractitionerRoleParser) GetExtension() any {
	return p.data["extension"]
}
