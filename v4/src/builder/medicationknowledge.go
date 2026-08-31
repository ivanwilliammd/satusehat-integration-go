package builder

import (
    "encoding/json"
)

type MedicationKnowledgeBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewMedicationKnowledgeBuilder() *MedicationKnowledgeBuilder {
    b := &MedicationKnowledgeBuilder{ResourceType: "MedicationKnowledge", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "MedicationKnowledge"
    return b
}

func (b *MedicationKnowledgeBuilder) setId(id string) *MedicationKnowledgeBuilder {
    b.Data["id"] = id
    return b
}

func (b *MedicationKnowledgeBuilder) addIdentifier(system, value string) *MedicationKnowledgeBuilder {
    ids, _ := b.Data["identifier"].([]map[string]string)
    ids = append(ids, map[string]string{"system": system, "value": value})
    b.Data["identifier"] = ids
    return b
}

func (b *MedicationKnowledgeBuilder) setStatus(status string) *MedicationKnowledgeBuilder {
    b.Data["status"] = status
    return b
}

func (b *MedicationKnowledgeBuilder) setCode(system, code, display string) *MedicationKnowledgeBuilder {
    b.Data["code"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *MedicationKnowledgeBuilder) setManufacturer(reference string, display ...string) *MedicationKnowledgeBuilder {
    ref := map[string]interface{}{"reference": reference}
    if len(display) > 0 && display[0] != "" {
        ref["display"] = display[0]
    }
    b.Data["manufacturer"] = ref
    return b
}

func (b *MedicationKnowledgeBuilder) setDoseForm(system, code, display string) *MedicationKnowledgeBuilder {
    b.Data["doseForm"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *MedicationKnowledgeBuilder) Build() map[string]interface{} {
    clean := make(map[string]interface{})
    for k, v := range b.Data {
        if v != nil {
            clean[k] = v
        }
    }
    return clean
}

func (b *MedicationKnowledgeBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Build())
}
