package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type ResearchSubjectBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewResearchSubjectBuilder() *ResearchSubjectBuilder {
    b := &ResearchSubjectBuilder{ResourceType: "ResearchSubject", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "ResearchSubject"
    return b
}

func (b *ResearchSubjectBuilder) setId(id string) *ResearchSubjectBuilder {
    b.Data["id"] = id
    return b
}

func (b *ResearchSubjectBuilder) addIdentifier(identifier *datatype.Identifier) *ResearchSubjectBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *ResearchSubjectBuilder) setStatus(status string) *ResearchSubjectBuilder {
    b.Data["status"] = status
    return b
}

func (b *ResearchSubjectBuilder) setSubject(reference string) *ResearchSubjectBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *ResearchSubjectBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *ResearchSubjectBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
