package parser

// AppointmentResponseParser extracts fields from a FHIR AppointmentResponse resource map.
//
// The element list is derived from PayloadBuilderAppointmentResponse in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type AppointmentResponseParser struct {
	data map[string]any
}

// NewAppointmentResponseParser creates a parser for the given AppointmentResponse payload.
func NewAppointmentResponseParser(data map[string]any) *AppointmentResponseParser {
	if data == nil {
		data = map[string]any{}
	}
	return &AppointmentResponseParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *AppointmentResponseParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *AppointmentResponseParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *AppointmentResponseParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *AppointmentResponseParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *AppointmentResponseParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *AppointmentResponseParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *AppointmentResponseParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *AppointmentResponseParser) GetEncounter() any {
	return p.data["encounter"]
}
