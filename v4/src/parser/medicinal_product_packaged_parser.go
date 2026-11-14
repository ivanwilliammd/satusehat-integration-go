package parser

// MedicinalProductPackagedParser extracts fields from a FHIR MedicinalProductPackaged resource map.
//
// The element list is derived from PayloadBuilderMedicinalProductPackaged in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type MedicinalProductPackagedParser struct {
	data map[string]any
}

// NewMedicinalProductPackagedParser creates a parser for the given MedicinalProductPackaged payload.
func NewMedicinalProductPackagedParser(data map[string]any) *MedicinalProductPackagedParser {
	if data == nil {
		data = map[string]any{}
	}
	return &MedicinalProductPackagedParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *MedicinalProductPackagedParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *MedicinalProductPackagedParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *MedicinalProductPackagedParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *MedicinalProductPackagedParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *MedicinalProductPackagedParser) GetStatus() any {
	return p.data["status"]
}

// GetPackageItem returns the packageItem element.
func (p *MedicinalProductPackagedParser) GetPackageItem() any {
	return p.data["packageItem"]
}

// GetDescription returns the description element.
func (p *MedicinalProductPackagedParser) GetDescription() any {
	return p.data["description"]
}

// GetManufacturer returns the manufacturer element.
func (p *MedicinalProductPackagedParser) GetManufacturer() any {
	return p.data["manufacturer"]
}
