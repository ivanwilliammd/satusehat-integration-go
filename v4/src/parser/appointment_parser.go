package parser

// AppointmentParser extracts fields from a FHIR Appointment resource map.
//
// The element list is derived from PayloadBuilderAppointment in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type AppointmentParser struct {
	data map[string]any
}

// NewAppointmentParser creates a parser for the given Appointment payload.
func NewAppointmentParser(data map[string]any) *AppointmentParser {
	if data == nil {
		data = map[string]any{}
	}
	return &AppointmentParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *AppointmentParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *AppointmentParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *AppointmentParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *AppointmentParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *AppointmentParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *AppointmentParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *AppointmentParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *AppointmentParser) GetEncounter() any {
	return p.data["encounter"]
}
