package builder

import (
    "encoding/json"
)

type InsurancePlanBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewInsurancePlanBuilder() *InsurancePlanBuilder {
    b := &InsurancePlanBuilder{ResourceType: "InsurancePlan", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "InsurancePlan"
    return b
}

func (b *InsurancePlanBuilder) setId(id string) *InsurancePlanBuilder {
    b.Data["id"] = id
    return b
}

func (b *InsurancePlanBuilder) addIdentifier(system, value string) *InsurancePlanBuilder {
    ids, _ := b.Data["identifier"].([]map[string]string)
    ids = append(ids, map[string]string{"system": system, "value": value})
    b.Data["identifier"] = ids
    return b
}

func (b *InsurancePlanBuilder) setStatus(status string) *InsurancePlanBuilder {
    b.Data["status"] = status
    return b
}

func (b *InsurancePlanBuilder) setName(value string) *InsurancePlanBuilder {
    b.Data["name"] = value
    return b
}

func (b *InsurancePlanBuilder) setType(system, code, display string) *InsurancePlanBuilder {
    b.Data["type"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *InsurancePlanBuilder) setAdministeredBy(reference string, display ...string) *InsurancePlanBuilder {
    ref := map[string]interface{}{"reference": reference}
    if len(display) > 0 && display[0] != "" {
        ref["display"] = display[0]
    }
    b.Data["administeredBy"] = ref
    return b
}

func (b *InsurancePlanBuilder) setCoverage(system, code, display string) *InsurancePlanBuilder {
    b.Data["coverage"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *InsurancePlanBuilder) Build() map[string]interface{} {
    clean := make(map[string]interface{})
    for k, v := range b.Data {
        if v != nil {
            clean[k] = v
        }
    }
    return clean
}

func (b *InsurancePlanBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Build())
}
