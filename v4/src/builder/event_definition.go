package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type EventDefinitionBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewEventDefinitionBuilder() *EventDefinitionBuilder {
	b := &EventDefinitionBuilder{ResourceType: "EventDefinition", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "EventDefinition"
	return b
}

func (b *EventDefinitionBuilder) SetId(id string) *EventDefinitionBuilder {
	b.Data["id"] = id
	return b
}

func (b *EventDefinitionBuilder) AddIdentifier(identifier *datatype.Identifier) *EventDefinitionBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *EventDefinitionBuilder) SetStatus(status string) *EventDefinitionBuilder {
	b.Data["status"] = status
	return b
}

func (b *EventDefinitionBuilder) SetSubject(reference string, display ...string) *EventDefinitionBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *EventDefinitionBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *EventDefinitionBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *EventDefinitionBuilder) SetCode(system, code, display string) *EventDefinitionBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *EventDefinitionBuilder) SetEncounter(reference string) *EventDefinitionBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
