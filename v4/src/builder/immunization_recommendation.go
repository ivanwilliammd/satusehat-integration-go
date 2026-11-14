package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ImmunizationRecommendationBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewImmunizationRecommendationBuilder() *ImmunizationRecommendationBuilder {
	b := &ImmunizationRecommendationBuilder{ResourceType: "ImmunizationRecommendation", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "ImmunizationRecommendation"
	return b
}

func (b *ImmunizationRecommendationBuilder) SetId(id string) *ImmunizationRecommendationBuilder {
	b.Data["id"] = id
	return b
}

func (b *ImmunizationRecommendationBuilder) AddIdentifier(identifier *datatype.Identifier) *ImmunizationRecommendationBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ImmunizationRecommendationBuilder) SetStatus(status string) *ImmunizationRecommendationBuilder {
	b.Data["status"] = status
	return b
}

func (b *ImmunizationRecommendationBuilder) SetSubject(reference string, display ...string) *ImmunizationRecommendationBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *ImmunizationRecommendationBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ImmunizationRecommendationBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *ImmunizationRecommendationBuilder) SetCode(system, code, display string) *ImmunizationRecommendationBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *ImmunizationRecommendationBuilder) SetEncounter(reference string) *ImmunizationRecommendationBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
