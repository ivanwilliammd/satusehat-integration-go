package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type MedicationStatementBuilder struct {
	Data map[string]interface{}
}

func NewMedicationStatementBuilder() *MedicationStatementBuilder {
	return &MedicationStatementBuilder{Data: make(map[string]interface{})}
}

func (b *MedicationStatementBuilder) SetId(id string) *MedicationStatementBuilder {
	b.Data["id"] = id
	return b
}

func (b *MedicationStatementBuilder) SetStatus(status string) *MedicationStatementBuilder {
	b.Data["status"] = status
	return b
}

func (b *MedicationStatementBuilder) SetMedication(cc *datatype.CodeableConcept) *MedicationStatementBuilder {
	b.Data["medication"] = cc.ToArray()
	return b
}

func (b *MedicationStatementBuilder) SetSubject(ref *datatype.Reference) *MedicationStatementBuilder {
	b.Data["subject"] = ref.ToArray()
	return b
}

func (b *MedicationStatementBuilder) SetEffectiveDateTime(dt string) *MedicationStatementBuilder {
	b.Data["effectiveDateTime"] = dt
	return b
}

func (b *MedicationStatementBuilder) SetDateAsserted(dt string) *MedicationStatementBuilder {
	b.Data["dateAsserted"] = dt
	return b
}

func (b *MedicationStatementBuilder) SetInformationSource(ref *datatype.Reference) *MedicationStatementBuilder {
	b.Data["informationSource"] = ref.ToArray()
	return b
}

func (b *MedicationStatementBuilder) Build() map[string]interface{} {
	dt := make(map[string]interface{})
	for k, v := range b.Data {
		if v == nil || v == "" {
			continue
		}
		dt[k] = v
	}
	dt["resourceType"] = "MedicationStatement"
	return dt
}
