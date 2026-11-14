package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type FlagBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewFlagBuilder() *FlagBuilder {
	b := &FlagBuilder{ResourceType: "Flag", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Flag"
	return b
}

func (b *FlagBuilder) SetId(id string) *FlagBuilder {
	b.Data["id"] = id
	return b
}

func (b *FlagBuilder) AddIdentifier(identifier *datatype.Identifier) *FlagBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *FlagBuilder) SetStatus(status string) *FlagBuilder {
	b.Data["status"] = status
	return b
}

func (b *FlagBuilder) SetSubject(reference string, display ...string) *FlagBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *FlagBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *FlagBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *FlagBuilder) SetCode(system, code, display string) *FlagBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *FlagBuilder) SetEncounter(reference string) *FlagBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
