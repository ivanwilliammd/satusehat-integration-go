package builder

import (
    "encoding/json"
)

type ObservationDefinitionBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewObservationDefinitionBuilder() *ObservationDefinitionBuilder {
    b := &ObservationDefinitionBuilder{ResourceType: "ObservationDefinition", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "ObservationDefinition"
    return b
}

func (b *ObservationDefinitionBuilder) setId(id string) *ObservationDefinitionBuilder {
    b.Data["id"] = id
    return b
}

func (b *ObservationDefinitionBuilder) addIdentifier(system, value string) *ObservationDefinitionBuilder {
    ids, _ := b.Data["identifier"].([]map[string]string)
    ids = append(ids, map[string]string{"system": system, "value": value})
    b.Data["identifier"] = ids
    return b
}

func (b *ObservationDefinitionBuilder) setStatus(status string) *ObservationDefinitionBuilder {
    b.Data["status"] = status
    return b
}

func (b *ObservationDefinitionBuilder) setCode(system, code, display string) *ObservationDefinitionBuilder {
    b.Data["code"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *ObservationDefinitionBuilder) setCategory(system, code, display string) *ObservationDefinitionBuilder {
    b.Data["category"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *ObservationDefinitionBuilder) setMethod(system, code, display string) *ObservationDefinitionBuilder {
    b.Data["method"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *ObservationDefinitionBuilder) Build() map[string]interface{} {
    clean := make(map[string]interface{})
    for k, v := range b.Data {
        if v != nil {
            clean[k] = v
        }
    }
    return clean
}

func (b *ObservationDefinitionBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Build())
}
