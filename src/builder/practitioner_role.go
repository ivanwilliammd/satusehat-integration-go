package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type PractitionerRoleBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewPractitionerRoleBuilder() *PractitionerRoleBuilder {
    b := &PractitionerRoleBuilder{ResourceType: "PractitionerRole", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "PractitionerRole"
    return b
}

func (b *PractitionerRoleBuilder) setId(id string) *PractitionerRoleBuilder {
    b.Data["id"] = id
    return b
}

func (b *PractitionerRoleBuilder) addIdentifier(identifier *datatype.Identifier) *PractitionerRoleBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *PractitionerRoleBuilder) setStatus(status string) *PractitionerRoleBuilder {
    b.Data["status"] = status
    return b
}

func (b *PractitionerRoleBuilder) setSubject(reference string) *PractitionerRoleBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *PractitionerRoleBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *PractitionerRoleBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
