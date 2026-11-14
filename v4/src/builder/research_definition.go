package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ResearchDefinitionBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewResearchDefinitionBuilder() *ResearchDefinitionBuilder {
	b := &ResearchDefinitionBuilder{ResourceType: "ResearchDefinition", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "ResearchDefinition"
	return b
}

func (b *ResearchDefinitionBuilder) SetId(id string) *ResearchDefinitionBuilder {
	b.Data["id"] = id
	return b
}

func (b *ResearchDefinitionBuilder) AddIdentifier(identifier *datatype.Identifier) *ResearchDefinitionBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ResearchDefinitionBuilder) SetStatus(status string) *ResearchDefinitionBuilder {
	b.Data["status"] = status
	return b
}

func (b *ResearchDefinitionBuilder) SetSubject(reference string, display ...string) *ResearchDefinitionBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *ResearchDefinitionBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ResearchDefinitionBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *ResearchDefinitionBuilder) SetCode(system, code, display string) *ResearchDefinitionBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *ResearchDefinitionBuilder) SetEncounter(reference string) *ResearchDefinitionBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
