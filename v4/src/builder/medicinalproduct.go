package builder

import (
    "encoding/json"
)

type MedicinalProductBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewMedicinalProductBuilder() *MedicinalProductBuilder {
    b := &MedicinalProductBuilder{ResourceType: "MedicinalProduct", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "MedicinalProduct"
    return b
}

func (b *MedicinalProductBuilder) setId(id string) *MedicinalProductBuilder {
    b.Data["id"] = id
    return b
}

func (b *MedicinalProductBuilder) addIdentifier(system, value string) *MedicinalProductBuilder {
    ids, _ := b.Data["identifier"].([]map[string]string)
    ids = append(ids, map[string]string{"system": system, "value": value})
    b.Data["identifier"] = ids
    return b
}

func (b *MedicinalProductBuilder) setStatus(status string) *MedicinalProductBuilder {
    b.Data["status"] = status
    return b
}

func (b *MedicinalProductBuilder) setName(value string) *MedicinalProductBuilder {
    b.Data["name"] = value
    return b
}

func (b *MedicinalProductBuilder) setManufacturer(reference string, display ...string) *MedicinalProductBuilder {
    ref := map[string]interface{}{"reference": reference}
    if len(display) > 0 && display[0] != "" {
        ref["display"] = display[0]
    }
    b.Data["manufacturer"] = ref
    return b
}

func (b *MedicinalProductBuilder) setProductClass(system, code, display string) *MedicinalProductBuilder {
    b.Data["productClass"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *MedicinalProductBuilder) Build() map[string]interface{} {
    clean := make(map[string]interface{})
    for k, v := range b.Data {
        if v != nil {
            clean[k] = v
        }
    }
    return clean
}

func (b *MedicinalProductBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Build())
}
