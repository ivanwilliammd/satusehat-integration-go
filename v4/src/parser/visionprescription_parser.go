package parser

// VisionPrescriptionParser extracts fields from a FHIR VisionPrescription resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/VisionPrescription.json, copied unchanged to
// src/parser/testdata/VisionPrescription.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type VisionPrescriptionParser struct {
	payload map[string]any
}

// NewVisionPrescriptionParser creates a new VisionPrescriptionParser instance.
func NewVisionPrescriptionParser(payload map[string]any) *VisionPrescriptionParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &VisionPrescriptionParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *VisionPrescriptionParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *VisionPrescriptionParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetStatus returns the status element unchanged.
func (p *VisionPrescriptionParser) GetStatus() any {
	return p.payload["status"]
}

// GetCreated returns the created element unchanged.
func (p *VisionPrescriptionParser) GetCreated() any {
	return p.payload["created"]
}

// GetPatient returns the patient element unchanged.
func (p *VisionPrescriptionParser) GetPatient() any {
	return p.payload["patient"]
}

// GetEncounter returns the encounter element unchanged.
func (p *VisionPrescriptionParser) GetEncounter() any {
	return p.payload["encounter"]
}

// GetDateWritten returns the dateWritten element unchanged.
func (p *VisionPrescriptionParser) GetDateWritten() any {
	return p.payload["dateWritten"]
}

// GetPrescriber returns the prescriber element unchanged.
func (p *VisionPrescriptionParser) GetPrescriber() any {
	return p.payload["prescriber"]
}

// GetLensSpecification returns the lensSpecification element unchanged.
func (p *VisionPrescriptionParser) GetLensSpecification() any {
	return p.payload["lensSpecification"]
}
