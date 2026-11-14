package parser

// ImagingSelectionParser extracts fields from a FHIR ImagingSelection resource map.
//
// The element list is derived from PayloadBuilderImagingSelection in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type ImagingSelectionParser struct {
	data map[string]any
}

// NewImagingSelectionParser creates a parser for the given ImagingSelection payload.
func NewImagingSelectionParser(data map[string]any) *ImagingSelectionParser {
	if data == nil {
		data = map[string]any{}
	}
	return &ImagingSelectionParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *ImagingSelectionParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *ImagingSelectionParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *ImagingSelectionParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *ImagingSelectionParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *ImagingSelectionParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *ImagingSelectionParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *ImagingSelectionParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *ImagingSelectionParser) GetEncounter() any {
	return p.data["encounter"]
}
