package parser

// MedicationKnowledgeParser extracts fields from a FHIR MedicationKnowledge resource map.
//
// The element list is derived from PayloadBuilderMedicationKnowledge in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type MedicationKnowledgeParser struct {
	data map[string]any
}

// NewMedicationKnowledgeParser creates a parser for the given MedicationKnowledge payload.
func NewMedicationKnowledgeParser(data map[string]any) *MedicationKnowledgeParser {
	if data == nil {
		data = map[string]any{}
	}
	return &MedicationKnowledgeParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *MedicationKnowledgeParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *MedicationKnowledgeParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *MedicationKnowledgeParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *MedicationKnowledgeParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *MedicationKnowledgeParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *MedicationKnowledgeParser) GetCode() any {
	return p.data["code"]
}

// GetManufacturer returns the manufacturer element.
func (p *MedicationKnowledgeParser) GetManufacturer() any {
	return p.data["manufacturer"]
}

// GetDoseForm returns the doseForm element.
func (p *MedicationKnowledgeParser) GetDoseForm() any {
	return p.data["doseForm"]
}
