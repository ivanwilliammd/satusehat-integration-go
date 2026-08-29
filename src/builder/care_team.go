package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type CareTeamBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewCareTeamBuilder() *CareTeamBuilder {
    b := &CareTeamBuilder{ResourceType: "CareTeam", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "CareTeam"
    return b
}

func (b *CareTeamBuilder) setId(id string) *CareTeamBuilder {
    b.Data["id"] = id
    return b
}

func (b *CareTeamBuilder) addIdentifier(identifier *datatype.Identifier) *CareTeamBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *CareTeamBuilder) setStatus(status string) *CareTeamBuilder {
    b.Data["status"] = status
    return b
}

func (b *CareTeamBuilder) setSubject(reference string) *CareTeamBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *CareTeamBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *CareTeamBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
