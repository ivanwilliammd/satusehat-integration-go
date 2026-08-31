package builder

import (
    "encoding/json"
)

type ActivityDefinitionBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewActivityDefinitionBuilder() *ActivityDefinitionBuilder {
    b := &ActivityDefinitionBuilder{ResourceType: "ActivityDefinition", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "ActivityDefinition"
    return b
}

func (b *ActivityDefinitionBuilder) setId(id string) *ActivityDefinitionBuilder {
    b.Data["id"] = id
    return b
}

func (b *ActivityDefinitionBuilder) addIdentifier(system, value string) *ActivityDefinitionBuilder {
    ids, _ := b.Data["identifier"].([]map[string]string)
    ids = append(ids, map[string]string{"system": system, "value": value})
    b.Data["identifier"] = ids
    return b
}

func (b *ActivityDefinitionBuilder) setStatus(status string) *ActivityDefinitionBuilder {
    b.Data["status"] = status
    return b
}

func (b *ActivityDefinitionBuilder) setDescription(value string) *ActivityDefinitionBuilder {
    b.Data["description"] = value
    return b
}

func (b *ActivityDefinitionBuilder) setKind(system, code, display string) *ActivityDefinitionBuilder {
    b.Data["kind"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *ActivityDefinitionBuilder) setCode(system, code, display string) *ActivityDefinitionBuilder {
    b.Data["code"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *ActivityDefinitionBuilder) setAuthor(reference string, display ...string) *ActivityDefinitionBuilder {
    ref := map[string]interface{}{"reference": reference}
    if len(display) > 0 && display[0] != "" {
        ref["display"] = display[0]
    }
    b.Data["author"] = ref
    return b
}

func (b *ActivityDefinitionBuilder) setTiming(value string) *ActivityDefinitionBuilder {
    b.Data["timing"] = value
    return b
}

func (b *ActivityDefinitionBuilder) Build() map[string]interface{} {
    clean := make(map[string]interface{})
    for k, v := range b.Data {
        if v != nil {
            clean[k] = v
        }
    }
    return clean
}

func (b *ActivityDefinitionBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Build())
}
