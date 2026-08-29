package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type TestScriptBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewTestScriptBuilder() *TestScriptBuilder {
    b := &TestScriptBuilder{ResourceType: "TestScript", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "TestScript"
    return b
}

func (b *TestScriptBuilder) setId(id string) *TestScriptBuilder {
    b.Data["id"] = id
    return b
}

func (b *TestScriptBuilder) addIdentifier(identifier *datatype.Identifier) *TestScriptBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *TestScriptBuilder) setStatus(status string) *TestScriptBuilder {
    b.Data["status"] = status
    return b
}

func (b *TestScriptBuilder) setSubject(reference string) *TestScriptBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *TestScriptBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *TestScriptBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
