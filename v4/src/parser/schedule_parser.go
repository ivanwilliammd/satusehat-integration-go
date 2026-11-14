package parser

// ScheduleParser extracts fields from a FHIR Schedule resource map.
//
// The element list is derived from PayloadBuilderSchedule in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type ScheduleParser struct {
	data map[string]any
}

// NewScheduleParser creates a parser for the given Schedule payload.
func NewScheduleParser(data map[string]any) *ScheduleParser {
	if data == nil {
		data = map[string]any{}
	}
	return &ScheduleParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *ScheduleParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *ScheduleParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *ScheduleParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *ScheduleParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *ScheduleParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *ScheduleParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *ScheduleParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *ScheduleParser) GetEncounter() any {
	return p.data["encounter"]
}
