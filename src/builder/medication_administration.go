package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type MedicationAdministrationBuilder struct {
	Data map[string]interface{}
}

func NewMedicationAdministrationBuilder() *MedicationAdministrationBuilder {
	return &MedicationAdministrationBuilder{Data: make(map[string]interface{})}
}

func (b *MedicationAdministrationBuilder) SetId(id string) *MedicationAdministrationBuilder {
	b.Data["id"] = id
	return b
}

func (b *MedicationAdministrationBuilder) SetStatus(status string) *MedicationAdministrationBuilder {
	b.Data["status"] = status
	return b
}

func (b *MedicationAdministrationBuilder) SetMedication(cc *datatype.CodeableConcept) *MedicationAdministrationBuilder {
	b.Data["medication"] = cc.ToArray()
	return b
}

func (b *MedicationAdministrationBuilder) SetSubject(ref *datatype.Reference) *MedicationAdministrationBuilder {
	b.Data["subject"] = ref.ToArray()
	return b
}

func (b *MedicationAdministrationBuilder) SetEncounter(ref *datatype.Reference) *MedicationAdministrationBuilder {
	b.Data["encounter"] = ref.ToArray()
	return b
}

func (b *MedicationAdministrationBuilder) SetEffectiveDateTime(dt string) *MedicationAdministrationBuilder {
	b.Data["effectiveDateTime"] = dt
	return b
}

func (b *MedicationAdministrationBuilder) SetRequester(ref *datatype.Reference) *MedicationAdministrationBuilder {
	b.Data["requester"] = ref.ToArray()
	return b
}

func (b *MedicationAdministrationBuilder) Build() map[string]interface{} {
	dt := make(map[string]interface{})
	for k, v := range b.Data {
		if v == nil || v == "" {
			continue
		}
		dt[k] = v
	}
	dt["resourceType"] = "MedicationAdministration"
	return dt
}
