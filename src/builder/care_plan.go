package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type CarePlanBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewCarePlanBuilder() *CarePlanBuilder {
    b := &CarePlanBuilder{ResourceType: "CarePlan", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "CarePlan"
    return b
}

func (b *CarePlanBuilder) setId(id string) *CarePlanBuilder {
    b.Data["id"] = id
    return b
}

func (b *CarePlanBuilder) addIdentifier(identifier *datatype.Identifier) *CarePlanBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *CarePlanBuilder) setStatus(status string) *CarePlanBuilder {
    b.Data["status"] = status
    return b
}

func (b *CarePlanBuilder) setSubject(reference string) *CarePlanBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *CarePlanBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *CarePlanBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
