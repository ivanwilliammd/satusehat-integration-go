package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type AppointmentBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewAppointmentBuilder() *AppointmentBuilder {
	b := &AppointmentBuilder{ResourceType: "Appointment", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Appointment"
	return b
}

func (b *AppointmentBuilder) SetId(id string) *AppointmentBuilder {
	b.Data["id"] = id
	return b
}

func (b *AppointmentBuilder) AddIdentifier(identifier *datatype.Identifier) *AppointmentBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *AppointmentBuilder) SetStatus(status string) *AppointmentBuilder {
	b.Data["status"] = status
	return b
}

func (b *AppointmentBuilder) SetSubject(reference string, display ...string) *AppointmentBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *AppointmentBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *AppointmentBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *AppointmentBuilder) SetCode(system, code, display string) *AppointmentBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *AppointmentBuilder) SetEncounter(reference string) *AppointmentBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
