package parser

// PatientParser extracts fields from a FHIR Patient resource map.
type PatientParser struct {
	patient map[string]any
}

// NewPatientParser creates a new PatientParser instance.
func NewPatientParser(patient map[string]any) *PatientParser {
	if patient == nil {
		patient = map[string]any{}
	}
	return &PatientParser{patient: patient}
}

func (p *PatientParser) GetIdentifiers() any {
	return p.patient["identifier"]
}

func (p *PatientParser) GetName() any {
	return p.patient["name"]
}

func (p *PatientParser) GetTelecom() any {
	return p.patient["telecom"]
}

func (p *PatientParser) GetGender() string {
	return strVal(p.patient["gender"])
}

func (p *PatientParser) GetBirthDate() string {
	return strVal(p.patient["birthDate"])
}

func (p *PatientParser) GetDeceased() any {
	return p.patient["deceasedBoolean"]
}

func (p *PatientParser) GetAddress() any {
	return p.patient["address"]
}

func (p *PatientParser) GetMaritalStatus() any {
	return p.patient["maritalStatus"]
}

func (p *PatientParser) GetMultipleBirth() any {
	if v, ok := p.patient["multipleBirthBoolean"]; ok && v != nil {
		return v
	}
	return p.patient["multipleBirthInteger"]
}

func (p *PatientParser) GetEmergencyContact() any {
	return p.patient["contact"]
}

func (p *PatientParser) GetCommunication() any {
	return p.patient["communication"]
}

func (p *PatientParser) GetExtension() any {
	return p.patient["extension"]
}
