package parser

// AllergyIntoleranceParser extracts fields from a FHIR AllergyIntolerance resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/AllergyIntolerance.json, copied unchanged to
// src/parser/testdata/AllergyIntolerance.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type AllergyIntoleranceParser struct {
	payload map[string]any
}

// NewAllergyIntoleranceParser creates a new AllergyIntoleranceParser instance.
func NewAllergyIntoleranceParser(payload map[string]any) *AllergyIntoleranceParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &AllergyIntoleranceParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *AllergyIntoleranceParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *AllergyIntoleranceParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetIdentifier returns the identifier element unchanged.
func (p *AllergyIntoleranceParser) GetIdentifier() any {
	return p.payload["identifier"]
}

// GetClinicalStatus returns the clinicalStatus element unchanged.
func (p *AllergyIntoleranceParser) GetClinicalStatus() any {
	return p.payload["clinicalStatus"]
}

// GetVerificationStatus returns the verificationStatus element unchanged.
func (p *AllergyIntoleranceParser) GetVerificationStatus() any {
	return p.payload["verificationStatus"]
}

// GetCategory returns the category element unchanged.
func (p *AllergyIntoleranceParser) GetCategory() any {
	return p.payload["category"]
}

// GetCode returns the code element unchanged.
func (p *AllergyIntoleranceParser) GetCode() any {
	return p.payload["code"]
}

// GetPatient returns the patient element unchanged.
func (p *AllergyIntoleranceParser) GetPatient() any {
	return p.payload["patient"]
}

// GetEncounter returns the encounter element unchanged.
func (p *AllergyIntoleranceParser) GetEncounter() any {
	return p.payload["encounter"]
}

// GetRecordedDate returns the recordedDate element unchanged.
func (p *AllergyIntoleranceParser) GetRecordedDate() any {
	return p.payload["recordedDate"]
}

// GetRecorder returns the recorder element unchanged.
func (p *AllergyIntoleranceParser) GetRecorder() any {
	return p.payload["recorder"]
}
