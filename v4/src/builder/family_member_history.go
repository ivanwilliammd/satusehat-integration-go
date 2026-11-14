package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type FamilyMemberHistoryBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewFamilyMemberHistoryBuilder() *FamilyMemberHistoryBuilder {
	b := &FamilyMemberHistoryBuilder{ResourceType: "FamilyMemberHistory", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "FamilyMemberHistory"
	return b
}

func (b *FamilyMemberHistoryBuilder) SetId(id string) *FamilyMemberHistoryBuilder {
	b.Data["id"] = id
	return b
}

func (b *FamilyMemberHistoryBuilder) AddIdentifier(identifier *datatype.Identifier) *FamilyMemberHistoryBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *FamilyMemberHistoryBuilder) SetStatus(status string) *FamilyMemberHistoryBuilder {
	b.Data["status"] = status
	return b
}

func (b *FamilyMemberHistoryBuilder) SetSubject(reference string, display ...string) *FamilyMemberHistoryBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *FamilyMemberHistoryBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *FamilyMemberHistoryBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetPatient sets patient from a Reference, as the PHP builder does.
func (b *FamilyMemberHistoryBuilder) SetPatient(patient *datatype.Reference) *FamilyMemberHistoryBuilder {
	b.Data["patient"] = patient.ToArray()
	return b
}

// SetDate sets date, as the PHP builder does.
func (b *FamilyMemberHistoryBuilder) SetDate(date string) *FamilyMemberHistoryBuilder {
	b.Data["date"] = date
	return b
}
