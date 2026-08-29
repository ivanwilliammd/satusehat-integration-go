package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type MedicationDispenseBuilder struct {
	Data map[string]interface{}
}

func NewMedicationDispenseBuilder() *MedicationDispenseBuilder {
	return &MedicationDispenseBuilder{Data: make(map[string]interface{})}
}

func (b *MedicationDispenseBuilder) SetId(id string) *MedicationDispenseBuilder {
	b.Data["id"] = id
	return b
}

func (b *MedicationDispenseBuilder) SetStatus(status string) *MedicationDispenseBuilder {
	b.Data["status"] = status
	return b
}

func (b *MedicationDispenseBuilder) SetMedication(cc *datatype.CodeableConcept) *MedicationDispenseBuilder {
	b.Data["medication"] = cc.ToArray()
	return b
}

func (b *MedicationDispenseBuilder) SetSubject(ref *datatype.Reference) *MedicationDispenseBuilder {
	b.Data["subject"] = ref.ToArray()
	return b
}

func (b *MedicationDispenseBuilder) SetEncounter(ref *datatype.Reference) *MedicationDispenseBuilder {
	b.Data["encounter"] = ref.ToArray()
	return b
}

func (b *MedicationDispenseBuilder) SetWhenPrepared(dt string) *MedicationDispenseBuilder {
	b.Data["whenPrepared"] = dt
	return b
}

func (b *MedicationDispenseBuilder) SetWhenHandedOver(dt string) *MedicationDispenseBuilder {
	b.Data["whenHandedOver"] = dt
	return b
}

func (b *MedicationDispenseBuilder) SetDestination(ref *datatype.Reference) *MedicationDispenseBuilder {
	b.Data["destination"] = ref.ToArray()
	return b
}

func (b *MedicationDispenseBuilder) Build() map[string]interface{} {
	dt := make(map[string]interface{})
	for k, v := range b.Data {
		if v == nil || v == "" {
			continue
		}
		dt[k] = v
	}
	dt["resourceType"] = "MedicationDispense"
	return dt
}
