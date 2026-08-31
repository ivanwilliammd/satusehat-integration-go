package builder

import (
    "encoding/json"
)

type MedicinalProductPackagedBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewMedicinalProductPackagedBuilder() *MedicinalProductPackagedBuilder {
    b := &MedicinalProductPackagedBuilder{ResourceType: "MedicinalProductPackaged", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "MedicinalProductPackaged"
    return b
}

func (b *MedicinalProductPackagedBuilder) setId(id string) *MedicinalProductPackagedBuilder {
    b.Data["id"] = id
    return b
}

func (b *MedicinalProductPackagedBuilder) addIdentifier(system, value string) *MedicinalProductPackagedBuilder {
    ids, _ := b.Data["identifier"].([]map[string]string)
    ids = append(ids, map[string]string{"system": system, "value": value})
    b.Data["identifier"] = ids
    return b
}

func (b *MedicinalProductPackagedBuilder) setStatus(status string) *MedicinalProductPackagedBuilder {
    b.Data["status"] = status
    return b
}

func (b *MedicinalProductPackagedBuilder) setPackageItem(system, code, display string) *MedicinalProductPackagedBuilder {
    b.Data["packageItem"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *MedicinalProductPackagedBuilder) setDescription(value string) *MedicinalProductPackagedBuilder {
    b.Data["description"] = value
    return b
}

func (b *MedicinalProductPackagedBuilder) setManufacturer(reference string, display ...string) *MedicinalProductPackagedBuilder {
    ref := map[string]interface{}{"reference": reference}
    if len(display) > 0 && display[0] != "" {
        ref["display"] = display[0]
    }
    b.Data["manufacturer"] = ref
    return b
}

func (b *MedicinalProductPackagedBuilder) Build() map[string]interface{} {
    clean := make(map[string]interface{})
    for k, v := range b.Data {
        if v != nil {
            clean[k] = v
        }
    }
    return clean
}

func (b *MedicinalProductPackagedBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Build())
}
