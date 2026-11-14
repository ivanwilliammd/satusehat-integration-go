package parser

// MedicationStatementParser extracts fields from a FHIR MedicationStatement resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/MedicationStatement.json, copied unchanged to
// src/parser/testdata/MedicationStatement.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type MedicationStatementParser struct {
	payload map[string]any
}

// NewMedicationStatementParser creates a new MedicationStatementParser instance.
func NewMedicationStatementParser(payload map[string]any) *MedicationStatementParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &MedicationStatementParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *MedicationStatementParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *MedicationStatementParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetContained returns the contained element unchanged.
func (p *MedicationStatementParser) GetContained() any {
	return p.payload["contained"]
}

// GetStatus returns the status element unchanged.
func (p *MedicationStatementParser) GetStatus() any {
	return p.payload["status"]
}

// GetCategory returns the category element unchanged.
func (p *MedicationStatementParser) GetCategory() any {
	return p.payload["category"]
}

// GetMedicationReference returns the medicationReference element unchanged.
func (p *MedicationStatementParser) GetMedicationReference() any {
	return p.payload["medicationReference"]
}

// GetSubject returns the subject element unchanged.
func (p *MedicationStatementParser) GetSubject() any {
	return p.payload["subject"]
}

// GetDosage returns the dosage element unchanged.
func (p *MedicationStatementParser) GetDosage() any {
	return p.payload["dosage"]
}

// GetEffectiveDateTime returns the effectiveDateTime element unchanged.
func (p *MedicationStatementParser) GetEffectiveDateTime() any {
	return p.payload["effectiveDateTime"]
}

// GetDateAsserted returns the dateAsserted element unchanged.
func (p *MedicationStatementParser) GetDateAsserted() any {
	return p.payload["dateAsserted"]
}

// GetInformationSource returns the informationSource element unchanged.
func (p *MedicationStatementParser) GetInformationSource() any {
	return p.payload["informationSource"]
}

// GetContext returns the context element unchanged.
func (p *MedicationStatementParser) GetContext() any {
	return p.payload["context"]
}
