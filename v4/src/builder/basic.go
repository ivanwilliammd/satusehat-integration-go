package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type BasicBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewBasicBuilder() *BasicBuilder {
	b := &BasicBuilder{ResourceType: "Basic", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Basic"
	return b
}

func (b *BasicBuilder) SetId(id string) *BasicBuilder {
	b.Data["id"] = id
	return b
}

func (b *BasicBuilder) AddIdentifier(identifier *datatype.Identifier) *BasicBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *BasicBuilder) SetStatus(status string) *BasicBuilder {
	b.Data["status"] = status
	return b
}

func (b *BasicBuilder) SetSubject(reference string, display ...string) *BasicBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *BasicBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *BasicBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *BasicBuilder) SetCode(system, code, display string) *BasicBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *BasicBuilder) SetEncounter(reference string) *BasicBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
