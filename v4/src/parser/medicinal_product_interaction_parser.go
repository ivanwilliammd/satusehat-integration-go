package parser

// MedicinalProductInteractionParser extracts fields from a FHIR MedicinalProductInteraction resource map.
//
// The element list is derived from PayloadBuilderMedicinalProductInteraction in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type MedicinalProductInteractionParser struct {
	data map[string]any
}

// NewMedicinalProductInteractionParser creates a parser for the given MedicinalProductInteraction payload.
func NewMedicinalProductInteractionParser(data map[string]any) *MedicinalProductInteractionParser {
	if data == nil {
		data = map[string]any{}
	}
	return &MedicinalProductInteractionParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *MedicinalProductInteractionParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *MedicinalProductInteractionParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *MedicinalProductInteractionParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *MedicinalProductInteractionParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *MedicinalProductInteractionParser) GetStatus() any {
	return p.data["status"]
}

// GetSubject returns the subject element.
func (p *MedicinalProductInteractionParser) GetSubject() any {
	return p.data["subject"]
}

// GetDescription returns the description element.
func (p *MedicinalProductInteractionParser) GetDescription() any {
	return p.data["description"]
}

// GetInteractant returns the interactant element.
func (p *MedicinalProductInteractionParser) GetInteractant() any {
	return p.data["interactant"]
}
