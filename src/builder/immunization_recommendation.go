package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ImmunizationRecommendationBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewImmunizationRecommendationBuilder() *ImmunizationRecommendationBuilder {
    b := &ImmunizationRecommendationBuilder{ResourceType: "ImmunizationRecommendation", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "ImmunizationRecommendation"
    return b
}

func (b *ImmunizationRecommendationBuilder) setId(id string) *ImmunizationRecommendationBuilder {
    b.Data["id"] = id
    return b
}

func (b *ImmunizationRecommendationBuilder) addIdentifier(identifier *datatype.Identifier) *ImmunizationRecommendationBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *ImmunizationRecommendationBuilder) setStatus(status string) *ImmunizationRecommendationBuilder {
    b.Data["status"] = status
    return b
}

func (b *ImmunizationRecommendationBuilder) setSubject(reference string) *ImmunizationRecommendationBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *ImmunizationRecommendationBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *ImmunizationRecommendationBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
