package builder

import (
    "encoding/json"
)

type MedicinalProductManufacturedBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewMedicinalProductManufacturedBuilder() *MedicinalProductManufacturedBuilder {
    b := &MedicinalProductManufacturedBuilder{ResourceType: "MedicinalProductManufactured", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "MedicinalProductManufactured"
    return b
}

func (b *MedicinalProductManufacturedBuilder) setId(id string) *MedicinalProductManufacturedBuilder {
    b.Data["id"] = id
    return b
}

func (b *MedicinalProductManufacturedBuilder) addIdentifier(system, value string) *MedicinalProductManufacturedBuilder {
    ids, _ := b.Data["identifier"].([]map[string]string)
    ids = append(ids, map[string]string{"system": system, "value": value})
    b.Data["identifier"] = ids
    return b
}

func (b *MedicinalProductManufacturedBuilder) setStatus(status string) *MedicinalProductManufacturedBuilder {
    b.Data["status"] = status
    return b
}

func (b *MedicinalProductManufacturedBuilder) setQuantity(value string) *MedicinalProductManufacturedBuilder {
    b.Data["quantity"] = value
    return b
}

func (b *MedicinalProductManufacturedBuilder) setManufacturer(reference string, display ...string) *MedicinalProductManufacturedBuilder {
    ref := map[string]interface{}{"reference": reference}
    if len(display) > 0 && display[0] != "" {
        ref["display"] = display[0]
    }
    b.Data["manufacturer"] = ref
    return b
}

func (b *MedicinalProductManufacturedBuilder) setIngredient(system, code, display string) *MedicinalProductManufacturedBuilder {
    b.Data["ingredient"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *MedicinalProductManufacturedBuilder) Build() map[string]interface{} {
    clean := make(map[string]interface{})
    for k, v := range b.Data {
        if v != nil {
            clean[k] = v
        }
    }
    return clean
}

func (b *MedicinalProductManufacturedBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Build())
}
