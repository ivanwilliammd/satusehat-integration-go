package parser

// MedicationParser extracts fields from a FHIR Medication resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/Medication.json, copied unchanged to
// src/parser/testdata/Medication.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type MedicationParser struct {
	payload map[string]any
}

// NewMedicationParser creates a new MedicationParser instance.
func NewMedicationParser(payload map[string]any) *MedicationParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &MedicationParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *MedicationParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *MedicationParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetMeta returns the meta element unchanged.
func (p *MedicationParser) GetMeta() any {
	return p.payload["meta"]
}

// GetIdentifier returns the identifier element unchanged.
func (p *MedicationParser) GetIdentifier() any {
	return p.payload["identifier"]
}

// GetCode returns the code element unchanged.
func (p *MedicationParser) GetCode() any {
	return p.payload["code"]
}

// GetStatus returns the status element unchanged.
func (p *MedicationParser) GetStatus() any {
	return p.payload["status"]
}

// GetManufacturer returns the manufacturer element unchanged.
func (p *MedicationParser) GetManufacturer() any {
	return p.payload["manufacturer"]
}

// GetForm returns the form element unchanged.
func (p *MedicationParser) GetForm() any {
	return p.payload["form"]
}

// GetIngredient returns the ingredient element unchanged.
func (p *MedicationParser) GetIngredient() any {
	return p.payload["ingredient"]
}

// GetBatch returns the batch element unchanged.
func (p *MedicationParser) GetBatch() any {
	return p.payload["batch"]
}

// GetExtension returns the extension element unchanged.
func (p *MedicationParser) GetExtension() any {
	return p.payload["extension"]
}
