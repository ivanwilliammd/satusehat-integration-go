package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type GoalBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewGoalBuilder() *GoalBuilder {
    b := &GoalBuilder{ResourceType: "Goal", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Goal"
    return b
}

func (b *GoalBuilder) setId(id string) *GoalBuilder {
    b.Data["id"] = id
    return b
}

func (b *GoalBuilder) addIdentifier(identifier *datatype.Identifier) *GoalBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *GoalBuilder) setStatus(status string) *GoalBuilder {
    b.Data["status"] = status
    return b
}

func (b *GoalBuilder) setSubject(reference string) *GoalBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *GoalBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *GoalBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
