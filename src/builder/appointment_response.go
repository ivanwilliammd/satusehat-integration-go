package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type AppointmentResponseBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewAppointmentResponseBuilder() *AppointmentResponseBuilder {
    b := &AppointmentResponseBuilder{ResourceType: "AppointmentResponse", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "AppointmentResponse"
    return b
}

func (b *AppointmentResponseBuilder) setId(id string) *AppointmentResponseBuilder {
    b.Data["id"] = id
    return b
}

func (b *AppointmentResponseBuilder) addIdentifier(identifier *datatype.Identifier) *AppointmentResponseBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *AppointmentResponseBuilder) setStatus(status string) *AppointmentResponseBuilder {
    b.Data["status"] = status
    return b
}

func (b *AppointmentResponseBuilder) setSubject(reference string) *AppointmentResponseBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *AppointmentResponseBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *AppointmentResponseBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
