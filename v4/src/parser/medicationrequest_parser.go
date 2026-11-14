package parser

// MedicationRequestParser extracts fields from a FHIR MedicationRequest resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/MedicationRequest.json, copied unchanged to
// src/parser/testdata/MedicationRequest.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type MedicationRequestParser struct {
	payload map[string]any
}

// NewMedicationRequestParser creates a new MedicationRequestParser instance.
func NewMedicationRequestParser(payload map[string]any) *MedicationRequestParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &MedicationRequestParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *MedicationRequestParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *MedicationRequestParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetContained returns the contained element unchanged.
func (p *MedicationRequestParser) GetContained() any {
	return p.payload["contained"]
}

// GetIdentifier returns the identifier element unchanged.
func (p *MedicationRequestParser) GetIdentifier() any {
	return p.payload["identifier"]
}

// GetStatus returns the status element unchanged.
func (p *MedicationRequestParser) GetStatus() any {
	return p.payload["status"]
}

// GetIntent returns the intent element unchanged.
func (p *MedicationRequestParser) GetIntent() any {
	return p.payload["intent"]
}

// GetCategory returns the category element unchanged.
func (p *MedicationRequestParser) GetCategory() any {
	return p.payload["category"]
}

// GetPriority returns the priority element unchanged.
func (p *MedicationRequestParser) GetPriority() any {
	return p.payload["priority"]
}

// GetMedicationReference returns the medicationReference element unchanged.
func (p *MedicationRequestParser) GetMedicationReference() any {
	return p.payload["medicationReference"]
}

// GetSubject returns the subject element unchanged.
func (p *MedicationRequestParser) GetSubject() any {
	return p.payload["subject"]
}

// GetEncounter returns the encounter element unchanged.
func (p *MedicationRequestParser) GetEncounter() any {
	return p.payload["encounter"]
}

// GetAuthoredOn returns the authoredOn element unchanged.
func (p *MedicationRequestParser) GetAuthoredOn() any {
	return p.payload["authoredOn"]
}

// GetRequester returns the requester element unchanged.
func (p *MedicationRequestParser) GetRequester() any {
	return p.payload["requester"]
}

// GetCourseOfTherapyType returns the courseOfTherapyType element unchanged.
func (p *MedicationRequestParser) GetCourseOfTherapyType() any {
	return p.payload["courseOfTherapyType"]
}

// GetDosageInstruction returns the dosageInstruction element unchanged.
func (p *MedicationRequestParser) GetDosageInstruction() any {
	return p.payload["dosageInstruction"]
}

// GetDispenseRequest returns the dispenseRequest element unchanged.
func (p *MedicationRequestParser) GetDispenseRequest() any {
	return p.payload["dispenseRequest"]
}
