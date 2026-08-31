package builder

import (
    "encoding/json"
)

type MedicinalProductIngredientBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewMedicinalProductIngredientBuilder() *MedicinalProductIngredientBuilder {
    b := &MedicinalProductIngredientBuilder{ResourceType: "MedicinalProductIngredient", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "MedicinalProductIngredient"
    return b
}

func (b *MedicinalProductIngredientBuilder) setId(id string) *MedicinalProductIngredientBuilder {
    b.Data["id"] = id
    return b
}

func (b *MedicinalProductIngredientBuilder) addIdentifier(system, value string) *MedicinalProductIngredientBuilder {
    ids, _ := b.Data["identifier"].([]map[string]string)
    ids = append(ids, map[string]string{"system": system, "value": value})
    b.Data["identifier"] = ids
    return b
}

func (b *MedicinalProductIngredientBuilder) setStatus(status string) *MedicinalProductIngredientBuilder {
    b.Data["status"] = status
    return b
}

func (b *MedicinalProductIngredientBuilder) setRole(system, code, display string) *MedicinalProductIngredientBuilder {
    b.Data["role"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *MedicinalProductIngredientBuilder) setSubstance(value string) *MedicinalProductIngredientBuilder {
    b.Data["substance"] = value
    return b
}

func (b *MedicinalProductIngredientBuilder) setQuantity(value string) *MedicinalProductIngredientBuilder {
    b.Data["quantity"] = value
    return b
}

func (b *MedicinalProductIngredientBuilder) Build() map[string]interface{} {
    clean := make(map[string]interface{})
    for k, v := range b.Data {
        if v != nil {
            clean[k] = v
        }
    }
    return clean
}

func (b *MedicinalProductIngredientBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Build())
}
