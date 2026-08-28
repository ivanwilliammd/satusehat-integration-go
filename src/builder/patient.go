package builder

type PatientBuilder struct {
	data map[string]interface{}
}

func NewPatient() *PatientBuilder {
	return &PatientBuilder{
		data: map[string]interface{}{
			"resourceType": "Patient",
		},
	}
}

func (b *PatientBuilder) SetNik(nik string) *PatientBuilder {
	// Add identifier for NIK
	return b
}

func (b *PatientBuilder) SetName(name string) *PatientBuilder {
	b.data["name"] = []map[string]interface{}{
		{"use": "official", "text": name},
	}
	return b
}

func (b *PatientBuilder) ToMap() map[string]interface{} {
	return b.data
}
