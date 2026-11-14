package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type SlotBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewSlotBuilder() *SlotBuilder {
	b := &SlotBuilder{ResourceType: "Slot", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Slot"
	return b
}

func (b *SlotBuilder) SetId(id string) *SlotBuilder {
	b.Data["id"] = id
	return b
}

func (b *SlotBuilder) AddIdentifier(identifier *datatype.Identifier) *SlotBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *SlotBuilder) SetStatus(status string) *SlotBuilder {
	b.Data["status"] = status
	return b
}

func (b *SlotBuilder) SetSubject(reference string, display ...string) *SlotBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *SlotBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *SlotBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *SlotBuilder) SetCode(system, code, display string) *SlotBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *SlotBuilder) SetEncounter(reference string) *SlotBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
