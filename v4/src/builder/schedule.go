package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ScheduleBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewScheduleBuilder() *ScheduleBuilder {
	b := &ScheduleBuilder{ResourceType: "Schedule", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Schedule"
	return b
}

func (b *ScheduleBuilder) SetId(id string) *ScheduleBuilder {
	b.Data["id"] = id
	return b
}

func (b *ScheduleBuilder) AddIdentifier(identifier *datatype.Identifier) *ScheduleBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ScheduleBuilder) SetStatus(status string) *ScheduleBuilder {
	b.Data["status"] = status
	return b
}

func (b *ScheduleBuilder) SetSubject(reference string, display ...string) *ScheduleBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *ScheduleBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ScheduleBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *ScheduleBuilder) SetCode(system, code, display string) *ScheduleBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *ScheduleBuilder) SetEncounter(reference string) *ScheduleBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
