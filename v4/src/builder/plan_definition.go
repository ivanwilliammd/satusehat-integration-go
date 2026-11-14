package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type PlanDefinitionBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewPlanDefinitionBuilder() *PlanDefinitionBuilder {
	b := &PlanDefinitionBuilder{ResourceType: "PlanDefinition", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "PlanDefinition"
	return b
}

func (b *PlanDefinitionBuilder) SetId(id string) *PlanDefinitionBuilder {
	b.Data["id"] = id
	return b
}

func (b *PlanDefinitionBuilder) AddIdentifier(identifier *datatype.Identifier) *PlanDefinitionBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *PlanDefinitionBuilder) SetStatus(status string) *PlanDefinitionBuilder {
	b.Data["status"] = status
	return b
}

func (b *PlanDefinitionBuilder) SetSubject(reference string, display ...string) *PlanDefinitionBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *PlanDefinitionBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *PlanDefinitionBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *PlanDefinitionBuilder) SetCode(system, code, display string) *PlanDefinitionBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *PlanDefinitionBuilder) SetEncounter(reference string) *PlanDefinitionBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
