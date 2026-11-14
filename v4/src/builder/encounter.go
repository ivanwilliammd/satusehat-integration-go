package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type EncounterBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewEncounterBuilder() *EncounterBuilder {
	b := &EncounterBuilder{ResourceType: "Encounter", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Encounter"
	return b
}

func (b *EncounterBuilder) SetId(id string) *EncounterBuilder {
	b.Data["id"] = id
	return b
}

func (b *EncounterBuilder) AddIdentifier(identifier *datatype.Identifier) *EncounterBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *EncounterBuilder) SetStatus(status string) *EncounterBuilder {
	b.Data["status"] = status
	return b
}

func (b *EncounterBuilder) SetSubject(reference string) *EncounterBuilder {
	b.Data["subject"] = datatype.Reference{}.ToArray()
	return b
}

func (b *EncounterBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *EncounterBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}
