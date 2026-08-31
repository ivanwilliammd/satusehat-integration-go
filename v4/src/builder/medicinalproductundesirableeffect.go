package builder

import (
    "encoding/json"
)

type MedicinalProductUndesirableEffectBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewMedicinalProductUndesirableEffectBuilder() *MedicinalProductUndesirableEffectBuilder {
    b := &MedicinalProductUndesirableEffectBuilder{ResourceType: "MedicinalProductUndesirableEffect", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "MedicinalProductUndesirableEffect"
    return b
}

func (b *MedicinalProductUndesirableEffectBuilder) setId(id string) *MedicinalProductUndesirableEffectBuilder {
    b.Data["id"] = id
    return b
}

func (b *MedicinalProductUndesirableEffectBuilder) addIdentifier(system, value string) *MedicinalProductUndesirableEffectBuilder {
    ids, _ := b.Data["identifier"].([]map[string]string)
    ids = append(ids, map[string]string{"system": system, "value": value})
    b.Data["identifier"] = ids
    return b
}

func (b *MedicinalProductUndesirableEffectBuilder) setStatus(status string) *MedicinalProductUndesirableEffectBuilder {
    b.Data["status"] = status
    return b
}

func (b *MedicinalProductUndesirableEffectBuilder) setSubject(reference string, display ...string) *MedicinalProductUndesirableEffectBuilder {
    ref := map[string]interface{}{"reference": reference}
    if len(display) > 0 && display[0] != "" {
        ref["display"] = display[0]
    }
    b.Data["subject"] = ref
    return b
}

func (b *MedicinalProductUndesirableEffectBuilder) setSymptomConditionEffect(system, code, display string) *MedicinalProductUndesirableEffectBuilder {
    b.Data["symptomConditionEffect"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *MedicinalProductUndesirableEffectBuilder) setClassification(system, code, display string) *MedicinalProductUndesirableEffectBuilder {
    b.Data["classification"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *MedicinalProductUndesirableEffectBuilder) Build() map[string]interface{} {
    clean := make(map[string]interface{})
    for k, v := range b.Data {
        if v != nil {
            clean[k] = v
        }
    }
    return clean
}

func (b *MedicinalProductUndesirableEffectBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Build())
}
