package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type FamilyMemberHistoryBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewFamilyMemberHistoryBuilder() *FamilyMemberHistoryBuilder {
    b := &FamilyMemberHistoryBuilder{ResourceType: "FamilyMemberHistory", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "FamilyMemberHistory"
    return b
}

func (b *FamilyMemberHistoryBuilder) setId(id string) *FamilyMemberHistoryBuilder {
    b.Data["id"] = id
    return b
}

func (b *FamilyMemberHistoryBuilder) addIdentifier(identifier *datatype.Identifier) *FamilyMemberHistoryBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *FamilyMemberHistoryBuilder) setStatus(status string) *FamilyMemberHistoryBuilder {
    b.Data["status"] = status
    return b
}

func (b *FamilyMemberHistoryBuilder) setSubject(reference string) *FamilyMemberHistoryBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *FamilyMemberHistoryBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *FamilyMemberHistoryBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
