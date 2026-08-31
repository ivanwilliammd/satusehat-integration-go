package builder

import (
    "encoding/json"
)

type ExplanationOfBenefitBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewExplanationOfBenefitBuilder() *ExplanationOfBenefitBuilder {
    b := &ExplanationOfBenefitBuilder{ResourceType: "ExplanationOfBenefit", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "ExplanationOfBenefit"
    return b
}

func (b *ExplanationOfBenefitBuilder) setId(id string) *ExplanationOfBenefitBuilder {
    b.Data["id"] = id
    return b
}

func (b *ExplanationOfBenefitBuilder) addIdentifier(system, value string) *ExplanationOfBenefitBuilder {
    ids, _ := b.Data["identifier"].([]map[string]string)
    ids = append(ids, map[string]string{"system": system, "value": value})
    b.Data["identifier"] = ids
    return b
}

func (b *ExplanationOfBenefitBuilder) setStatus(status string) *ExplanationOfBenefitBuilder {
    b.Data["status"] = status
    return b
}

func (b *ExplanationOfBenefitBuilder) setType(system, code, display string) *ExplanationOfBenefitBuilder {
    b.Data["type"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *ExplanationOfBenefitBuilder) setUse(system, code, display string) *ExplanationOfBenefitBuilder {
    b.Data["use"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *ExplanationOfBenefitBuilder) setPatient(reference string, display ...string) *ExplanationOfBenefitBuilder {
    ref := map[string]interface{}{"reference": reference}
    if len(display) > 0 && display[0] != "" {
        ref["display"] = display[0]
    }
    b.Data["patient"] = ref
    return b
}

func (b *ExplanationOfBenefitBuilder) setCreated(value string) *ExplanationOfBenefitBuilder {
    b.Data["created"] = value
    return b
}

func (b *ExplanationOfBenefitBuilder) setProvider(reference string, display ...string) *ExplanationOfBenefitBuilder {
    ref := map[string]interface{}{"reference": reference}
    if len(display) > 0 && display[0] != "" {
        ref["display"] = display[0]
    }
    b.Data["provider"] = ref
    return b
}

func (b *ExplanationOfBenefitBuilder) Build() map[string]interface{} {
    clean := make(map[string]interface{})
    for k, v := range b.Data {
        if v != nil {
            clean[k] = v
        }
    }
    return clean
}

func (b *ExplanationOfBenefitBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Build())
}
