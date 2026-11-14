package parser

// MedicationAdministrationParser extracts fields from a FHIR MedicationAdministration resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/MedicationAdministration.json, copied unchanged to
// src/parser/testdata/MedicationAdministration.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type MedicationAdministrationParser struct {
	payload map[string]any
}

// NewMedicationAdministrationParser creates a new MedicationAdministrationParser instance.
func NewMedicationAdministrationParser(payload map[string]any) *MedicationAdministrationParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &MedicationAdministrationParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *MedicationAdministrationParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *MedicationAdministrationParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetContained returns the contained element unchanged.
func (p *MedicationAdministrationParser) GetContained() any {
	return p.payload["contained"]
}

// GetStatus returns the status element unchanged.
func (p *MedicationAdministrationParser) GetStatus() any {
	return p.payload["status"]
}

// GetCategory returns the category element unchanged.
func (p *MedicationAdministrationParser) GetCategory() any {
	return p.payload["category"]
}

// GetMedicationReference returns the medicationReference element unchanged.
func (p *MedicationAdministrationParser) GetMedicationReference() any {
	return p.payload["medicationReference"]
}

// GetSubject returns the subject element unchanged.
func (p *MedicationAdministrationParser) GetSubject() any {
	return p.payload["subject"]
}

// GetContext returns the context element unchanged.
func (p *MedicationAdministrationParser) GetContext() any {
	return p.payload["context"]
}

// GetEffectivePeriod returns the effectivePeriod element unchanged.
func (p *MedicationAdministrationParser) GetEffectivePeriod() any {
	return p.payload["effectivePeriod"]
}

// GetPerformer returns the performer element unchanged.
func (p *MedicationAdministrationParser) GetPerformer() any {
	return p.payload["performer"]
}

// GetReasonCode returns the reasonCode element unchanged.
func (p *MedicationAdministrationParser) GetReasonCode() any {
	return p.payload["reasonCode"]
}

// GetRequest returns the request element unchanged.
func (p *MedicationAdministrationParser) GetRequest() any {
	return p.payload["request"]
}

// GetDosage returns the dosage element unchanged.
func (p *MedicationAdministrationParser) GetDosage() any {
	return p.payload["dosage"]
}
