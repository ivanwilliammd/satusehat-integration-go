package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type AppointmentResponseBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewAppointmentResponseBuilder() *AppointmentResponseBuilder {
	b := &AppointmentResponseBuilder{ResourceType: "AppointmentResponse", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "AppointmentResponse"
	return b
}

func (b *AppointmentResponseBuilder) SetId(id string) *AppointmentResponseBuilder {
	b.Data["id"] = id
	return b
}

func (b *AppointmentResponseBuilder) AddIdentifier(identifier *datatype.Identifier) *AppointmentResponseBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *AppointmentResponseBuilder) SetStatus(status string) *AppointmentResponseBuilder {
	b.Data["status"] = status
	return b
}

func (b *AppointmentResponseBuilder) SetSubject(reference string, display ...string) *AppointmentResponseBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *AppointmentResponseBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *AppointmentResponseBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *AppointmentResponseBuilder) SetCode(system, code, display string) *AppointmentResponseBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *AppointmentResponseBuilder) SetEncounter(reference string) *AppointmentResponseBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
