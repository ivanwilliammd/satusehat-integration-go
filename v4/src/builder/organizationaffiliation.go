package builder

import (
    "encoding/json"
)

type OrganizationAffiliationBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewOrganizationAffiliationBuilder() *OrganizationAffiliationBuilder {
    b := &OrganizationAffiliationBuilder{ResourceType: "OrganizationAffiliation", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "OrganizationAffiliation"
    return b
}

func (b *OrganizationAffiliationBuilder) setId(id string) *OrganizationAffiliationBuilder {
    b.Data["id"] = id
    return b
}

func (b *OrganizationAffiliationBuilder) addIdentifier(system, value string) *OrganizationAffiliationBuilder {
    ids, _ := b.Data["identifier"].([]map[string]string)
    ids = append(ids, map[string]string{"system": system, "value": value})
    b.Data["identifier"] = ids
    return b
}

func (b *OrganizationAffiliationBuilder) setStatus(status string) *OrganizationAffiliationBuilder {
    b.Data["status"] = status
    return b
}

func (b *OrganizationAffiliationBuilder) setOrganization(reference string, display ...string) *OrganizationAffiliationBuilder {
    ref := map[string]interface{}{"reference": reference}
    if len(display) > 0 && display[0] != "" {
        ref["display"] = display[0]
    }
    b.Data["organization"] = ref
    return b
}

func (b *OrganizationAffiliationBuilder) setParticipatingOrganization(reference string, display ...string) *OrganizationAffiliationBuilder {
    ref := map[string]interface{}{"reference": reference}
    if len(display) > 0 && display[0] != "" {
        ref["display"] = display[0]
    }
    b.Data["participatingOrganization"] = ref
    return b
}

func (b *OrganizationAffiliationBuilder) setNetwork(system, code, display string) *OrganizationAffiliationBuilder {
    b.Data["network"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *OrganizationAffiliationBuilder) setCode(system, code, display string) *OrganizationAffiliationBuilder {
    b.Data["code"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *OrganizationAffiliationBuilder) Build() map[string]interface{} {
    clean := make(map[string]interface{})
    for k, v := range b.Data {
        if v != nil {
            clean[k] = v
        }
    }
    return clean
}

func (b *OrganizationAffiliationBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Build())
}
