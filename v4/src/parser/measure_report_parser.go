package parser

// MeasureReportParser extracts fields from a FHIR MeasureReport resource map.
//
// The element list is derived from PayloadBuilderMeasureReport in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type MeasureReportParser struct {
	data map[string]any
}

// NewMeasureReportParser creates a parser for the given MeasureReport payload.
func NewMeasureReportParser(data map[string]any) *MeasureReportParser {
	if data == nil {
		data = map[string]any{}
	}
	return &MeasureReportParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *MeasureReportParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *MeasureReportParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *MeasureReportParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *MeasureReportParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *MeasureReportParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *MeasureReportParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *MeasureReportParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *MeasureReportParser) GetEncounter() any {
	return p.data["encounter"]
}
