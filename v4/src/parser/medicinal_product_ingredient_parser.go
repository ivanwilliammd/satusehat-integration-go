package parser

// MedicinalProductIngredientParser extracts fields from a FHIR MedicinalProductIngredient resource map.
//
// The element list is derived from PayloadBuilderMedicinalProductIngredient in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type MedicinalProductIngredientParser struct {
	data map[string]any
}

// NewMedicinalProductIngredientParser creates a parser for the given MedicinalProductIngredient payload.
func NewMedicinalProductIngredientParser(data map[string]any) *MedicinalProductIngredientParser {
	if data == nil {
		data = map[string]any{}
	}
	return &MedicinalProductIngredientParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *MedicinalProductIngredientParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *MedicinalProductIngredientParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *MedicinalProductIngredientParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *MedicinalProductIngredientParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *MedicinalProductIngredientParser) GetStatus() any {
	return p.data["status"]
}

// GetRole returns the role element.
func (p *MedicinalProductIngredientParser) GetRole() any {
	return p.data["role"]
}

// GetSubstance returns the substance element.
func (p *MedicinalProductIngredientParser) GetSubstance() any {
	return p.data["substance"]
}

// GetQuantity returns the quantity element.
func (p *MedicinalProductIngredientParser) GetQuantity() any {
	return p.data["quantity"]
}
