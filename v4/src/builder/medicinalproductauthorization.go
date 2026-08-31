package builder

import (
    "encoding/json"
)

type MedicinalProductAuthorizationBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewMedicinalProductAuthorizationBuilder() *MedicinalProductAuthorizationBuilder {
    b := &MedicinalProductAuthorizationBuilder{ResourceType: "MedicinalProductAuthorization", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "MedicinalProductAuthorization"
    return b
}

func (b *MedicinalProductAuthorizationBuilder) setId(id string) *MedicinalProductAuthorizationBuilder {
    b.Data["id"] = id
    return b
}

func (b *MedicinalProductAuthorizationBuilder) addIdentifier(system, value string) *MedicinalProductAuthorizationBuilder {
    ids, _ := b.Data["identifier"].([]map[string]string)
    ids = append(ids, map[string]string{"system": system, "value": value})
    b.Data["identifier"] = ids
    return b
}

func (b *MedicinalProductAuthorizationBuilder) setStatus(status string) *MedicinalProductAuthorizationBuilder {
    b.Data["status"] = status
    return b
}

func (b *MedicinalProductAuthorizationBuilder) setCountry(system, code, display string) *MedicinalProductAuthorizationBuilder {
    b.Data["country"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *MedicinalProductAuthorizationBuilder) setRegulator(reference string, display ...string) *MedicinalProductAuthorizationBuilder {
    ref := map[string]interface{}{"reference": reference}
    if len(display) > 0 && display[0] != "" {
        ref["display"] = display[0]
    }
    b.Data["regulator"] = ref
    return b
}

func (b *MedicinalProductAuthorizationBuilder) setValidityPeriod(value string) *MedicinalProductAuthorizationBuilder {
    b.Data["validityPeriod"] = value
    return b
}

func (b *MedicinalProductAuthorizationBuilder) Build() map[string]interface{} {
    clean := make(map[string]interface{})
    for k, v := range b.Data {
        if v != nil {
            clean[k] = v
        }
    }
    return clean
}

func (b *MedicinalProductAuthorizationBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Build())
}
