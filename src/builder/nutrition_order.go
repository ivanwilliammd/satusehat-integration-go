package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type NutritionOrderBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewNutritionOrderBuilder() *NutritionOrderBuilder {
    b := &NutritionOrderBuilder{ResourceType: "NutritionOrder", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "NutritionOrder"
    return b
}

func (b *NutritionOrderBuilder) setId(id string) *NutritionOrderBuilder {
    b.Data["id"] = id
    return b
}

func (b *NutritionOrderBuilder) addIdentifier(identifier *datatype.Identifier) *NutritionOrderBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *NutritionOrderBuilder) setStatus(status string) *NutritionOrderBuilder {
    b.Data["status"] = status
    return b
}

func (b *NutritionOrderBuilder) setSubject(reference string) *NutritionOrderBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *NutritionOrderBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *NutritionOrderBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
