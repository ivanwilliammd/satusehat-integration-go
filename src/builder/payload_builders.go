package builder

type BaseBuilder struct {
	Data map[string]interface{}
}

func (b *BaseBuilder) ToMap() map[string]interface{} {
	return b.Data
}

type PatientBuilder struct{ BaseBuilder }
func NewPatient() *PatientBuilder {
	return &PatientBuilder{BaseBuilder: BaseBuilder{Data: map[string]interface{}{"resourceType": "Patient"}}}
}
func (b *PatientBuilder) SetNik(nik string) *PatientBuilder {
	b.Data["identifier"] = []map[string]interface{}{
		{"system": "https://fhir.kemkes.go.id/id/nik", "value": nik},
	}
	return b
}
func (b *PatientBuilder) SetName(name string) *PatientBuilder {
	b.Data["name"] = []map[string]interface{}{{"use": "official", "text": name}}
	return b
}
func (b *PatientBuilder) SetGender(gender string) *PatientBuilder {
	b.Data["gender"] = gender
	return b
}
func (b *PatientBuilder) SetBirthDate(birthDate string) *PatientBuilder {
	b.Data["birthDate"] = birthDate
	return b
}

type PractitionerBuilder struct{ BaseBuilder }
func NewPractitioner() *PractitionerBuilder {
	return &PractitionerBuilder{BaseBuilder: BaseBuilder{Data: map[string]interface{}{"resourceType": "Practitioner"}}}
}
func (b *PractitionerBuilder) SetNik(nik string) *PractitionerBuilder {
	b.Data["identifier"] = []map[string]interface{}{
		{"system": "https://fhir.kemkes.go.id/id/nik", "value": nik},
	}
	return b
}
func (b *PractitionerBuilder) SetName(name string) *PractitionerBuilder {
	b.Data["name"] = []map[string]interface{}{{"use": "official", "text": name}}
	return b
}

type OrganizationBuilder struct{ BaseBuilder }
func NewOrganization() *OrganizationBuilder {
	return &OrganizationBuilder{BaseBuilder: BaseBuilder{Data: map[string]interface{}{"resourceType": "Organization"}}}
}
func (b *OrganizationBuilder) SetId(id string) *OrganizationBuilder {
	b.Data["id"] = id
	return b
}
func (b *OrganizationBuilder) SetName(name string) *OrganizationBuilder {
	b.Data["name"] = name
	return b
}

type LocationBuilder struct{ BaseBuilder }
func NewLocation() *LocationBuilder {
	return &LocationBuilder{BaseBuilder: BaseBuilder{Data: map[string]interface{}{"resourceType": "Location"}}}
}
func (b *LocationBuilder) SetName(name string) *LocationBuilder {
	b.Data["name"] = name
	return b
}
func (b *LocationBuilder) SetManagingOrganization(orgRef string) *LocationBuilder {
	b.Data["managingOrganization"] = map[string]interface{}{"reference": orgRef}
	return b
}

type EncounterBuilder struct{ BaseBuilder }
func NewEncounter() *EncounterBuilder {
	return &EncounterBuilder{BaseBuilder: BaseBuilder{Data: map[string]interface{}{"resourceType": "Encounter"}}}
}
func (b *EncounterBuilder) SetStatus(status string) *EncounterBuilder {
	b.Data["status"] = status
	return b
}
func (b *EncounterBuilder) SetSubject(patientRef string, display string) *EncounterBuilder {
	b.Data["subject"] = map[string]interface{}{"reference": patientRef, "display": display}
	return b
}

type ConditionBuilder struct{ BaseBuilder }
func NewCondition() *ConditionBuilder {
	return &ConditionBuilder{BaseBuilder: BaseBuilder{Data: map[string]interface{}{"resourceType": "Condition"}}}
}
func (b *ConditionBuilder) SetClinicalStatus(code string) *ConditionBuilder {
	b.Data["clinicalStatus"] = map[string]interface{}{
		"coding": []map[string]interface{}{
			{"system": "http://terminology.hl7.org/CodeSystem/condition-clinical", "code": code},
		},
	}
	return b
}
func (b *ConditionBuilder) SetSubject(patientRef string, display string) *ConditionBuilder {
	b.Data["subject"] = map[string]interface{}{"reference": patientRef, "display": display}
	return b
}
