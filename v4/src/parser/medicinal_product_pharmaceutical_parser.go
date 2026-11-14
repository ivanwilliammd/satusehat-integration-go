package parser

// MedicinalProductPharmaceuticalParser extracts fields from a FHIR MedicinalProductPharmaceutical resource map.
//
// The element list is derived from PayloadBuilderMedicinalProductPharmaceutical in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type MedicinalProductPharmaceuticalParser struct {
	data map[string]any
}

// NewMedicinalProductPharmaceuticalParser creates a parser for the given MedicinalProductPharmaceutical payload.
func NewMedicinalProductPharmaceuticalParser(data map[string]any) *MedicinalProductPharmaceuticalParser {
	if data == nil {
		data = map[string]any{}
	}
	return &MedicinalProductPharmaceuticalParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *MedicinalProductPharmaceuticalParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *MedicinalProductPharmaceuticalParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *MedicinalProductPharmaceuticalParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *MedicinalProductPharmaceuticalParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *MedicinalProductPharmaceuticalParser) GetStatus() any {
	return p.data["status"]
}

// GetRoute returns the route element.
func (p *MedicinalProductPharmaceuticalParser) GetRoute() any {
	return p.data["route"]
}

// GetDoseForm returns the doseForm element.
func (p *MedicinalProductPharmaceuticalParser) GetDoseForm() any {
	return p.data["doseForm"]
}

// GetIngredient returns the ingredient element.
func (p *MedicinalProductPharmaceuticalParser) GetIngredient() any {
	return p.data["ingredient"]
}
