package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ExampleScenarioBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewExampleScenarioBuilder() *ExampleScenarioBuilder {
    b := &ExampleScenarioBuilder{ResourceType: "ExampleScenario", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "ExampleScenario"
    return b
}

func (b *ExampleScenarioBuilder) setId(id string) *ExampleScenarioBuilder {
    b.Data["id"] = id
    return b
}

func (b *ExampleScenarioBuilder) addIdentifier(identifier *datatype.Identifier) *ExampleScenarioBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *ExampleScenarioBuilder) setStatus(status string) *ExampleScenarioBuilder {
    b.Data["status"] = status
    return b
}

func (b *ExampleScenarioBuilder) setSubject(reference string) *ExampleScenarioBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *ExampleScenarioBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *ExampleScenarioBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
