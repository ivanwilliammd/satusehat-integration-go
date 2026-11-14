package parser

// MedicinalProductIndicationParser extracts fields from a FHIR MedicinalProductIndication resource map.
//
// The element list is derived from PayloadBuilderMedicinalProductIndication in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type MedicinalProductIndicationParser struct {
	data map[string]any
}

// NewMedicinalProductIndicationParser creates a parser for the given MedicinalProductIndication payload.
func NewMedicinalProductIndicationParser(data map[string]any) *MedicinalProductIndicationParser {
	if data == nil {
		data = map[string]any{}
	}
	return &MedicinalProductIndicationParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *MedicinalProductIndicationParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *MedicinalProductIndicationParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *MedicinalProductIndicationParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *MedicinalProductIndicationParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *MedicinalProductIndicationParser) GetStatus() any {
	return p.data["status"]
}

// GetSubject returns the subject element.
func (p *MedicinalProductIndicationParser) GetSubject() any {
	return p.data["subject"]
}

// GetDisease returns the disease element.
func (p *MedicinalProductIndicationParser) GetDisease() any {
	return p.data["disease"]
}

// GetComorbidity returns the comorbidity element.
func (p *MedicinalProductIndicationParser) GetComorbidity() any {
	return p.data["comorbidity"]
}
