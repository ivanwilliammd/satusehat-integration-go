package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type AdverseEventBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewAdverseEventBuilder() *AdverseEventBuilder {
	b := &AdverseEventBuilder{ResourceType: "AdverseEvent", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "AdverseEvent"
	return b
}

func (b *AdverseEventBuilder) SetId(id string) *AdverseEventBuilder {
	b.Data["id"] = id
	return b
}

func (b *AdverseEventBuilder) AddIdentifier(identifier *datatype.Identifier) *AdverseEventBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *AdverseEventBuilder) SetStatus(status string) *AdverseEventBuilder {
	b.Data["status"] = status
	return b
}

func (b *AdverseEventBuilder) SetSubject(reference string, display ...string) *AdverseEventBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *AdverseEventBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *AdverseEventBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *AdverseEventBuilder) SetCode(system, code, display string) *AdverseEventBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *AdverseEventBuilder) SetEncounter(reference string) *AdverseEventBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
