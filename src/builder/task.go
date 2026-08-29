package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type TaskBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewTaskBuilder() *TaskBuilder {
    b := &TaskBuilder{ResourceType: "Task", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Task"
    return b
}

func (b *TaskBuilder) setId(id string) *TaskBuilder {
    b.Data["id"] = id
    return b
}

func (b *TaskBuilder) addIdentifier(identifier *datatype.Identifier) *TaskBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *TaskBuilder) setStatus(status string) *TaskBuilder {
    b.Data["status"] = status
    return b
}

func (b *TaskBuilder) setSubject(reference string) *TaskBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *TaskBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *TaskBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
