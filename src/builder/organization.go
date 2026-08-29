package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type OrganizationBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewOrganizationBuilder() *OrganizationBuilder {
    b := &OrganizationBuilder{ResourceType: "Organization", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Organization"
    return b
}

func (b *OrganizationBuilder) setId(id string) *OrganizationBuilder {
    b.Data["id"] = id
    return b
}

func (b *OrganizationBuilder) addIdentifier(identifier *datatype.Identifier) *OrganizationBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *OrganizationBuilder) setStatus(status string) *OrganizationBuilder {
    b.Data["status"] = status
    return b
}

func (b *OrganizationBuilder) setSubject(reference string) *OrganizationBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *OrganizationBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *OrganizationBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
