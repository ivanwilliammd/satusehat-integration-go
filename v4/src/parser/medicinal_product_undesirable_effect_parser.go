package parser

// MedicinalProductUndesirableEffectParser extracts fields from a FHIR MedicinalProductUndesirableEffect resource map.
//
// The element list is derived from PayloadBuilderMedicinalProductUndesirableEffect in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type MedicinalProductUndesirableEffectParser struct {
	data map[string]any
}

// NewMedicinalProductUndesirableEffectParser creates a parser for the given MedicinalProductUndesirableEffect payload.
func NewMedicinalProductUndesirableEffectParser(data map[string]any) *MedicinalProductUndesirableEffectParser {
	if data == nil {
		data = map[string]any{}
	}
	return &MedicinalProductUndesirableEffectParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *MedicinalProductUndesirableEffectParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *MedicinalProductUndesirableEffectParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *MedicinalProductUndesirableEffectParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *MedicinalProductUndesirableEffectParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *MedicinalProductUndesirableEffectParser) GetStatus() any {
	return p.data["status"]
}

// GetSubject returns the subject element.
func (p *MedicinalProductUndesirableEffectParser) GetSubject() any {
	return p.data["subject"]
}

// GetSymptomConditionEffect returns the symptomConditionEffect element.
func (p *MedicinalProductUndesirableEffectParser) GetSymptomConditionEffect() any {
	return p.data["symptomConditionEffect"]
}

// GetClassification returns the classification element.
func (p *MedicinalProductUndesirableEffectParser) GetClassification() any {
	return p.data["classification"]
}
