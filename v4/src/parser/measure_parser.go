package parser

// MeasureParser extracts fields from a FHIR Measure resource map.
//
// The element list is derived from PayloadBuilderMeasure in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type MeasureParser struct {
	data map[string]any
}

// NewMeasureParser creates a parser for the given Measure payload.
func NewMeasureParser(data map[string]any) *MeasureParser {
	if data == nil {
		data = map[string]any{}
	}
	return &MeasureParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *MeasureParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *MeasureParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *MeasureParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *MeasureParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *MeasureParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *MeasureParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *MeasureParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *MeasureParser) GetEncounter() any {
	return p.data["encounter"]
}
