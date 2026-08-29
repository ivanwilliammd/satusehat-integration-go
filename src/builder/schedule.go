package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ScheduleBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewScheduleBuilder() *ScheduleBuilder {
    b := &ScheduleBuilder{ResourceType: "Schedule", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Schedule"
    return b
}

func (b *ScheduleBuilder) setId(id string) *ScheduleBuilder {
    b.Data["id"] = id
    return b
}

func (b *ScheduleBuilder) addIdentifier(identifier *datatype.Identifier) *ScheduleBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *ScheduleBuilder) setStatus(status string) *ScheduleBuilder {
    b.Data["status"] = status
    return b
}

func (b *ScheduleBuilder) setSubject(reference string) *ScheduleBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *ScheduleBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *ScheduleBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
