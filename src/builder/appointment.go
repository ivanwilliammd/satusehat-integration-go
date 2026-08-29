package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type AppointmentBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewAppointmentBuilder() *AppointmentBuilder {
    b := &AppointmentBuilder{ResourceType: "Appointment", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Appointment"
    return b
}

func (b *AppointmentBuilder) setId(id string) *AppointmentBuilder {
    b.Data["id"] = id
    return b
}

func (b *AppointmentBuilder) addIdentifier(identifier *datatype.Identifier) *AppointmentBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *AppointmentBuilder) setStatus(status string) *AppointmentBuilder {
    b.Data["status"] = status
    return b
}

func (b *AppointmentBuilder) setSubject(reference string) *AppointmentBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *AppointmentBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *AppointmentBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
