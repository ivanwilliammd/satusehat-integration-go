package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type PersonBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewPersonBuilder() *PersonBuilder {
	b := &PersonBuilder{ResourceType: "Person", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Person"
	return b
}

func (b *PersonBuilder) SetId(id string) *PersonBuilder {
	b.Data["id"] = id
	return b
}

func (b *PersonBuilder) AddIdentifier(identifier *datatype.Identifier) *PersonBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *PersonBuilder) SetStatus(status string) *PersonBuilder {
	b.Data["status"] = status
	return b
}

func (b *PersonBuilder) SetSubject(reference string, display ...string) *PersonBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *PersonBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *PersonBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *PersonBuilder) SetCode(system, code, display string) *PersonBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *PersonBuilder) SetEncounter(reference string) *PersonBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
