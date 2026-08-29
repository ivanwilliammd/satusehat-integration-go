package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type SearchParameterBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewSearchParameterBuilder() *SearchParameterBuilder {
    b := &SearchParameterBuilder{ResourceType: "SearchParameter", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "SearchParameter"
    return b
}

func (b *SearchParameterBuilder) setId(id string) *SearchParameterBuilder {
    b.Data["id"] = id
    return b
}

func (b *SearchParameterBuilder) addIdentifier(identifier *datatype.Identifier) *SearchParameterBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *SearchParameterBuilder) setStatus(status string) *SearchParameterBuilder {
    b.Data["status"] = status
    return b
}

func (b *SearchParameterBuilder) setSubject(reference string) *SearchParameterBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *SearchParameterBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *SearchParameterBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
