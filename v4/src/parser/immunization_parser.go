package parser

// ImmunizationParser extracts fields from a FHIR Immunization resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/Immunization.json, copied unchanged to
// src/parser/testdata/Immunization.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type ImmunizationParser struct {
	payload map[string]any
}

// NewImmunizationParser creates a new ImmunizationParser instance.
func NewImmunizationParser(payload map[string]any) *ImmunizationParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &ImmunizationParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *ImmunizationParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *ImmunizationParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetStatus returns the status element unchanged.
func (p *ImmunizationParser) GetStatus() any {
	return p.payload["status"]
}

// GetVaccineCode returns the vaccineCode element unchanged.
func (p *ImmunizationParser) GetVaccineCode() any {
	return p.payload["vaccineCode"]
}

// GetPatient returns the patient element unchanged.
func (p *ImmunizationParser) GetPatient() any {
	return p.payload["patient"]
}

// GetEncounter returns the encounter element unchanged.
func (p *ImmunizationParser) GetEncounter() any {
	return p.payload["encounter"]
}

// GetOccurrenceDateTime returns the occurrenceDateTime element unchanged.
func (p *ImmunizationParser) GetOccurrenceDateTime() any {
	return p.payload["occurrenceDateTime"]
}

// GetRecorded returns the recorded element unchanged.
func (p *ImmunizationParser) GetRecorded() any {
	return p.payload["recorded"]
}

// GetReaction returns the reaction element unchanged.
func (p *ImmunizationParser) GetReaction() any {
	return p.payload["reaction"]
}

// GetPrimarySource returns the primarySource element unchanged.
func (p *ImmunizationParser) GetPrimarySource() any {
	return p.payload["primarySource"]
}

// GetLocation returns the location element unchanged.
func (p *ImmunizationParser) GetLocation() any {
	return p.payload["location"]
}

// GetLotNumber returns the lotNumber element unchanged.
func (p *ImmunizationParser) GetLotNumber() any {
	return p.payload["lotNumber"]
}

// GetExpirationDate returns the expirationDate element unchanged.
func (p *ImmunizationParser) GetExpirationDate() any {
	return p.payload["expirationDate"]
}

// GetRoute returns the route element unchanged.
func (p *ImmunizationParser) GetRoute() any {
	return p.payload["route"]
}

// GetDoseQuantity returns the doseQuantity element unchanged.
func (p *ImmunizationParser) GetDoseQuantity() any {
	return p.payload["doseQuantity"]
}

// GetPerformer returns the performer element unchanged.
func (p *ImmunizationParser) GetPerformer() any {
	return p.payload["performer"]
}

// GetReasonCode returns the reasonCode element unchanged.
func (p *ImmunizationParser) GetReasonCode() any {
	return p.payload["reasonCode"]
}

// GetProtocolApplied returns the protocolApplied element unchanged.
func (p *ImmunizationParser) GetProtocolApplied() any {
	return p.payload["protocolApplied"]
}
