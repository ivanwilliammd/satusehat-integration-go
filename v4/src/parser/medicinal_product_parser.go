package parser

// MedicinalProductParser extracts fields from a FHIR MedicinalProduct resource map.
//
// The element list is derived from PayloadBuilderMedicinalProduct in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type MedicinalProductParser struct {
	data map[string]any
}

// NewMedicinalProductParser creates a parser for the given MedicinalProduct payload.
func NewMedicinalProductParser(data map[string]any) *MedicinalProductParser {
	if data == nil {
		data = map[string]any{}
	}
	return &MedicinalProductParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *MedicinalProductParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *MedicinalProductParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *MedicinalProductParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *MedicinalProductParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *MedicinalProductParser) GetStatus() any {
	return p.data["status"]
}

// GetName returns the name element.
func (p *MedicinalProductParser) GetName() any {
	return p.data["name"]
}

// GetManufacturer returns the manufacturer element.
func (p *MedicinalProductParser) GetManufacturer() any {
	return p.data["manufacturer"]
}

// GetProductClass returns the productClass element.
func (p *MedicinalProductParser) GetProductClass() any {
	return p.data["productClass"]
}
