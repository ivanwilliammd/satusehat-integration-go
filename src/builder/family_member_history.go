package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type FamilyMemberHistoryBuilder struct {
	Data map[string]interface{}
}

func NewFamilyMemberHistoryBuilder() *FamilyMemberHistoryBuilder {
	return &FamilyMemberHistoryBuilder{Data: make(map[string]interface{})}
}

func (b *FamilyMemberHistoryBuilder) SetId(id string) *FamilyMemberHistoryBuilder {
	b.Data["id"] = id
	return b
}

func (b *FamilyMemberHistoryBuilder) SetStatus(status string) *FamilyMemberHistoryBuilder {
	b.Data["status"] = status
	return b
}

func (b *FamilyMemberHistoryBuilder) SetPatient(ref *datatype.Reference) *FamilyMemberHistoryBuilder {
	b.Data["patient"] = ref.ToArray()
	return b
}

func (b *FamilyMemberHistoryBuilder) SetRelationship(cc *datatype.CodeableConcept) *FamilyMemberHistoryBuilder {
	b.Data["relationship"] = cc.ToArray()
	return b
}

func (b *FamilyMemberHistoryBuilder) SetCode(cc *datatype.CodeableConcept) *FamilyMemberHistoryBuilder {
	b.Data["code"] = cc.ToArray()
	return b
}

func (b *FamilyMemberHistoryBuilder) SetDate(date string) *FamilyMemberHistoryBuilder {
	b.Data["date"] = date
	return b
}

func (b *FamilyMemberHistoryBuilder) Build() map[string]interface{} {
	dt := make(map[string]interface{})
	for k, v := range b.Data {
		if v == nil || v == "" {
			continue
		}
		dt[k] = v
	}
	dt["resourceType"] = "FamilyMemberHistory"
	return dt
}
