package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type EpisodeOfCareBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewEpisodeOfCareBuilder() *EpisodeOfCareBuilder {
    b := &EpisodeOfCareBuilder{ResourceType: "EpisodeOfCare", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "EpisodeOfCare"
    return b
}

func (b *EpisodeOfCareBuilder) setId(id string) *EpisodeOfCareBuilder {
    b.Data["id"] = id
    return b
}

func (b *EpisodeOfCareBuilder) addIdentifier(identifier *datatype.Identifier) *EpisodeOfCareBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *EpisodeOfCareBuilder) setStatus(status string) *EpisodeOfCareBuilder {
    b.Data["status"] = status
    return b
}

func (b *EpisodeOfCareBuilder) setSubject(reference string) *EpisodeOfCareBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *EpisodeOfCareBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *EpisodeOfCareBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
