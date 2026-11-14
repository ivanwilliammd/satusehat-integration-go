package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type PractitionerRoleBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewPractitionerRoleBuilder() *PractitionerRoleBuilder {
	b := &PractitionerRoleBuilder{ResourceType: "PractitionerRole", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "PractitionerRole"
	return b
}

func (b *PractitionerRoleBuilder) SetId(id string) *PractitionerRoleBuilder {
	b.Data["id"] = id
	return b
}

func (b *PractitionerRoleBuilder) AddIdentifier(identifier *datatype.Identifier) *PractitionerRoleBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *PractitionerRoleBuilder) SetStatus(status string) *PractitionerRoleBuilder {
	b.Data["status"] = status
	return b
}

func (b *PractitionerRoleBuilder) SetSubject(reference string, display ...string) *PractitionerRoleBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *PractitionerRoleBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *PractitionerRoleBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}
