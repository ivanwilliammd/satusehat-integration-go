package parser

// MolecularSequenceParser extracts fields from a FHIR MolecularSequence resource map.
type MolecularSequenceParser struct {
	molecularSequence map[string]any
}

// NewMolecularSequenceParser creates a new MolecularSequenceParser instance.
func NewMolecularSequenceParser(molecularSequence map[string]any) *MolecularSequenceParser {
	if molecularSequence == nil {
		molecularSequence = map[string]any{}
	}
	return &MolecularSequenceParser{molecularSequence: molecularSequence}
}

func (p *MolecularSequenceParser) GetText() any {
	return p.molecularSequence["text"]
}

func (p *MolecularSequenceParser) GetType() string {
	return strVal(p.molecularSequence["type"])
}

func (p *MolecularSequenceParser) GetCoordinateSystem() any {
	return p.molecularSequence["coordinateSystem"]
}

func (p *MolecularSequenceParser) GetPatientReference() string {
	return removePrefix(nestedStr(p.molecularSequence, "patient", "reference"), "Patient/")
}

func (p *MolecularSequenceParser) GetSpecimenReference() string {
	return removePrefix(nestedStr(p.molecularSequence, "specimen", "reference"), "Specimen/")
}

func (p *MolecularSequenceParser) GetDeviceReference() string {
	return removePrefix(nestedStr(p.molecularSequence, "device", "reference"), "Device/")
}

func (p *MolecularSequenceParser) GetPerformerReference() string {
	return removePrefix(nestedStr(p.molecularSequence, "performer", "reference"), "Organization/")
}

func (p *MolecularSequenceParser) GetReferenceSeq() any {
	return p.molecularSequence["referenceSeq"]
}

func (p *MolecularSequenceParser) GetVariants() any {
	return p.molecularSequence["variant"]
}
