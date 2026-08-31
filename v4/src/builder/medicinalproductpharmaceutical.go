package builder

import (
    "encoding/json"
)

type MedicinalProductPharmaceuticalBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewMedicinalProductPharmaceuticalBuilder() *MedicinalProductPharmaceuticalBuilder {
    b := &MedicinalProductPharmaceuticalBuilder{ResourceType: "MedicinalProductPharmaceutical", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "MedicinalProductPharmaceutical"
    return b
}

func (b *MedicinalProductPharmaceuticalBuilder) setId(id string) *MedicinalProductPharmaceuticalBuilder {
    b.Data["id"] = id
    return b
}

func (b *MedicinalProductPharmaceuticalBuilder) addIdentifier(system, value string) *MedicinalProductPharmaceuticalBuilder {
    ids, _ := b.Data["identifier"].([]map[string]string)
    ids = append(ids, map[string]string{"system": system, "value": value})
    b.Data["identifier"] = ids
    return b
}

func (b *MedicinalProductPharmaceuticalBuilder) setStatus(status string) *MedicinalProductPharmaceuticalBuilder {
    b.Data["status"] = status
    return b
}

func (b *MedicinalProductPharmaceuticalBuilder) setRoute(system, code, display string) *MedicinalProductPharmaceuticalBuilder {
    b.Data["route"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *MedicinalProductPharmaceuticalBuilder) setDoseForm(system, code, display string) *MedicinalProductPharmaceuticalBuilder {
    b.Data["doseForm"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *MedicinalProductPharmaceuticalBuilder) setIngredient(system, code, display string) *MedicinalProductPharmaceuticalBuilder {
    b.Data["ingredient"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *MedicinalProductPharmaceuticalBuilder) Build() map[string]interface{} {
    clean := make(map[string]interface{})
    for k, v := range b.Data {
        if v != nil {
            clean[k] = v
        }
    }
    return clean
}

func (b *MedicinalProductPharmaceuticalBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Build())
}
