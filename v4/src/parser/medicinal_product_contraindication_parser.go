package parser

// MedicinalProductContraindicationParser extracts fields from a FHIR MedicinalProductContraindication resource map.
//
// The element list is derived from PayloadBuilderMedicinalProductContraindication in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type MedicinalProductContraindicationParser struct {
	data map[string]any
}

// NewMedicinalProductContraindicationParser creates a parser for the given MedicinalProductContraindication payload.
func NewMedicinalProductContraindicationParser(data map[string]any) *MedicinalProductContraindicationParser {
	if data == nil {
		data = map[string]any{}
	}
	return &MedicinalProductContraindicationParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *MedicinalProductContraindicationParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *MedicinalProductContraindicationParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *MedicinalProductContraindicationParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *MedicinalProductContraindicationParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *MedicinalProductContraindicationParser) GetStatus() any {
	return p.data["status"]
}

// GetSubject returns the subject element.
func (p *MedicinalProductContraindicationParser) GetSubject() any {
	return p.data["subject"]
}

// GetDisease returns the disease element.
func (p *MedicinalProductContraindicationParser) GetDisease() any {
	return p.data["disease"]
}

// GetComorbidity returns the comorbidity element.
func (p *MedicinalProductContraindicationParser) GetComorbidity() any {
	return p.data["comorbidity"]
}
