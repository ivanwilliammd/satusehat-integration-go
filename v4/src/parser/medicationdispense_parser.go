package parser

// MedicationDispenseParser extracts fields from a FHIR MedicationDispense resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/MedicationDispense.json, copied unchanged to
// src/parser/testdata/MedicationDispense.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type MedicationDispenseParser struct {
	payload map[string]any
}

// NewMedicationDispenseParser creates a new MedicationDispenseParser instance.
func NewMedicationDispenseParser(payload map[string]any) *MedicationDispenseParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &MedicationDispenseParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *MedicationDispenseParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *MedicationDispenseParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetContained returns the contained element unchanged.
func (p *MedicationDispenseParser) GetContained() any {
	return p.payload["contained"]
}

// GetIdentifier returns the identifier element unchanged.
func (p *MedicationDispenseParser) GetIdentifier() any {
	return p.payload["identifier"]
}

// GetStatus returns the status element unchanged.
func (p *MedicationDispenseParser) GetStatus() any {
	return p.payload["status"]
}

// GetCategory returns the category element unchanged.
func (p *MedicationDispenseParser) GetCategory() any {
	return p.payload["category"]
}

// GetMedicationReference returns the medicationReference element unchanged.
func (p *MedicationDispenseParser) GetMedicationReference() any {
	return p.payload["medicationReference"]
}

// GetSubject returns the subject element unchanged.
func (p *MedicationDispenseParser) GetSubject() any {
	return p.payload["subject"]
}

// GetContext returns the context element unchanged.
func (p *MedicationDispenseParser) GetContext() any {
	return p.payload["context"]
}

// GetPerformer returns the performer element unchanged.
func (p *MedicationDispenseParser) GetPerformer() any {
	return p.payload["performer"]
}

// GetLocation returns the location element unchanged.
func (p *MedicationDispenseParser) GetLocation() any {
	return p.payload["location"]
}

// GetAuthorizingPrescription returns the authorizingPrescription element unchanged.
func (p *MedicationDispenseParser) GetAuthorizingPrescription() any {
	return p.payload["authorizingPrescription"]
}

// GetQuantity returns the quantity element unchanged.
func (p *MedicationDispenseParser) GetQuantity() any {
	return p.payload["quantity"]
}

// GetDaysSupply returns the daysSupply element unchanged.
func (p *MedicationDispenseParser) GetDaysSupply() any {
	return p.payload["daysSupply"]
}

// GetWhenPrepared returns the whenPrepared element unchanged.
func (p *MedicationDispenseParser) GetWhenPrepared() any {
	return p.payload["whenPrepared"]
}

// GetWhenHandedOver returns the whenHandedOver element unchanged.
func (p *MedicationDispenseParser) GetWhenHandedOver() any {
	return p.payload["whenHandedOver"]
}

// GetDosageInstruction returns the dosageInstruction element unchanged.
func (p *MedicationDispenseParser) GetDosageInstruction() any {
	return p.payload["dosageInstruction"]
}
