package parser

// MedicinalProductManufacturedParser extracts fields from a FHIR MedicinalProductManufactured resource map.
//
// The element list is derived from PayloadBuilderMedicinalProductManufactured in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type MedicinalProductManufacturedParser struct {
	data map[string]any
}

// NewMedicinalProductManufacturedParser creates a parser for the given MedicinalProductManufactured payload.
func NewMedicinalProductManufacturedParser(data map[string]any) *MedicinalProductManufacturedParser {
	if data == nil {
		data = map[string]any{}
	}
	return &MedicinalProductManufacturedParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *MedicinalProductManufacturedParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *MedicinalProductManufacturedParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *MedicinalProductManufacturedParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *MedicinalProductManufacturedParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *MedicinalProductManufacturedParser) GetStatus() any {
	return p.data["status"]
}

// GetQuantity returns the quantity element.
func (p *MedicinalProductManufacturedParser) GetQuantity() any {
	return p.data["quantity"]
}

// GetManufacturer returns the manufacturer element.
func (p *MedicinalProductManufacturedParser) GetManufacturer() any {
	return p.data["manufacturer"]
}

// GetIngredient returns the ingredient element.
func (p *MedicinalProductManufacturedParser) GetIngredient() any {
	return p.data["ingredient"]
}
