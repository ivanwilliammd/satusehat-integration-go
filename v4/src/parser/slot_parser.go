package parser

// SlotParser extracts fields from a FHIR Slot resource map.
//
// The element list is derived from PayloadBuilderSlot in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type SlotParser struct {
	data map[string]any
}

// NewSlotParser creates a parser for the given Slot payload.
func NewSlotParser(data map[string]any) *SlotParser {
	if data == nil {
		data = map[string]any{}
	}
	return &SlotParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *SlotParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *SlotParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *SlotParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *SlotParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *SlotParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *SlotParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *SlotParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *SlotParser) GetEncounter() any {
	return p.data["encounter"]
}
