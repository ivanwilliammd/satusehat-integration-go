package builder

import (
    "encoding/json"
)

type MedicinalProductIndicationBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewMedicinalProductIndicationBuilder() *MedicinalProductIndicationBuilder {
    b := &MedicinalProductIndicationBuilder{ResourceType: "MedicinalProductIndication", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "MedicinalProductIndication"
    return b
}

func (b *MedicinalProductIndicationBuilder) setId(id string) *MedicinalProductIndicationBuilder {
    b.Data["id"] = id
    return b
}

func (b *MedicinalProductIndicationBuilder) addIdentifier(system, value string) *MedicinalProductIndicationBuilder {
    ids, _ := b.Data["identifier"].([]map[string]string)
    ids = append(ids, map[string]string{"system": system, "value": value})
    b.Data["identifier"] = ids
    return b
}

func (b *MedicinalProductIndicationBuilder) setStatus(status string) *MedicinalProductIndicationBuilder {
    b.Data["status"] = status
    return b
}

func (b *MedicinalProductIndicationBuilder) setSubject(reference string, display ...string) *MedicinalProductIndicationBuilder {
    ref := map[string]interface{}{"reference": reference}
    if len(display) > 0 && display[0] != "" {
        ref["display"] = display[0]
    }
    b.Data["subject"] = ref
    return b
}

func (b *MedicinalProductIndicationBuilder) setDisease(system, code, display string) *MedicinalProductIndicationBuilder {
    b.Data["disease"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *MedicinalProductIndicationBuilder) setComorbidity(reference string, display ...string) *MedicinalProductIndicationBuilder {
    ref := map[string]interface{}{"reference": reference}
    if len(display) > 0 && display[0] != "" {
        ref["display"] = display[0]
    }
    b.Data["comorbidity"] = ref
    return b
}

func (b *MedicinalProductIndicationBuilder) Build() map[string]interface{} {
    clean := make(map[string]interface{})
    for k, v := range b.Data {
        if v != nil {
            clean[k] = v
        }
    }
    return clean
}

func (b *MedicinalProductIndicationBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Build())
}
