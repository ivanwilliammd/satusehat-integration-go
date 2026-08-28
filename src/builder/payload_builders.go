package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type BaseBuilder struct {
	Data map[string]interface{}
}

func (b *BaseBuilder) ToMap() map[string]interface{} { return b.Data }
func (b *BaseBuilder) SetId(id string) { b.Data["id"] = id }
func (b *BaseBuilder) SetActive(active bool) { b.Data["active"] = active }

type PatientBuilder struct{ BaseBuilder }
func NewPatient() *PatientBuilder {
	return &PatientBuilder{BaseBuilder: BaseBuilder{Data: map[string]interface{}{"resourceType": "Patient"}}}
}
func (b *PatientBuilder) AddIdentifier(id datatype.Identifier) *PatientBuilder {
	b.Data["identifier"] = append(b.GetSlice("identifier"), id.Value)
	return b
}
func (b *PatientBuilder) SetGender(gender string) *PatientBuilder { b.Data["gender"] = gender; return b }
func (b *PatientBuilder) SetBirthDate(dt string) *PatientBuilder { b.Data["birthDate"] = dt; return b }
func (b *PatientBuilder) AddName(name datatype.HumanName) *PatientBuilder {
	b.Data["name"] = append(b.GetSlice("name"), name.ToArray())
	return b
}
func (b *PatientBuilder) AddAddress(addr datatype.Address) *PatientBuilder {
	b.Data["address"] = append(b.GetSlice("address"), addr.ToArray())
	return b
}
func (b *PatientBuilder) AddTelecom(tc datatype.ContactPoint) *PatientBuilder {
	b.Data["telecom"] = append(b.GetSlice("telecom"), tc.ToArray())
	return b
}

func (b *BaseBuilder) GetSlice(key string) []interface{} {
	if val, ok := b.Data[key].([]interface{}); ok {
		return val
	}
	return []interface{}{}
}

type PractitionerBuilder struct{ BaseBuilder }
func NewPractitioner() *PractitionerBuilder {
	return &PractitionerBuilder{BaseBuilder: BaseBuilder{Data: map[string]interface{}{"resourceType": "Practitioner"}}}
}

type OrganizationBuilder struct{ BaseBuilder }
func NewOrganization() *OrganizationBuilder {
	return &OrganizationBuilder{BaseBuilder: BaseBuilder{Data: map[string]interface{}{"resourceType": "Organization"}}}
}
func (b *OrganizationBuilder) SetName(name string) *OrganizationBuilder { b.Data["name"] = name; return b }

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
