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
func (b *PatientBuilder) SetName(name string) *PatientBuilder {
	b.Data["name"] = []map[string]interface{}{{"use": "official", "text": name}}
	return b
}

type PractitionerBuilder struct{ BaseBuilder }
func NewPractitioner() *PractitionerBuilder {
	return &PractitionerBuilder{BaseBuilder: BaseBuilder{Data: map[string]interface{}{"resourceType": "Practitioner"}}}
}
func (b *PractitionerBuilder) SetName(name string) *PractitionerBuilder {
	b.Data["name"] = []map[string]interface{}{{"use": "official", "text": name}}
	return b
}

type OrganizationBuilder struct{ BaseBuilder }
func NewOrganization() *OrganizationBuilder {
	return &OrganizationBuilder{BaseBuilder: BaseBuilder{Data: map[string]interface{}{"resourceType": "Organization"}}}
}

type LocationBuilder struct{ BaseBuilder }
func NewLocation() *LocationBuilder {
	return &LocationBuilder{BaseBuilder: BaseBuilder{Data: map[string]interface{}{"resourceType": "Location"}}}
}

type EncounterBuilder struct{ BaseBuilder }
func NewEncounter() *EncounterBuilder {
	return &EncounterBuilder{BaseBuilder: BaseBuilder{Data: map[string]interface{}{"resourceType": "Encounter"}}}
}

type ConditionBuilder struct{ BaseBuilder }
func NewCondition() *ConditionBuilder {
	return &ConditionBuilder{BaseBuilder: BaseBuilder{Data: map[string]interface{}{"resourceType": "Condition"}}}
}
