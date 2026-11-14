package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ResearchElementDefinitionBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewResearchElementDefinitionBuilder() *ResearchElementDefinitionBuilder {
	b := &ResearchElementDefinitionBuilder{ResourceType: "ResearchElementDefinition", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "ResearchElementDefinition"
	return b
}

func (b *ResearchElementDefinitionBuilder) SetId(id string) *ResearchElementDefinitionBuilder {
	b.Data["id"] = id
	return b
}

func (b *ResearchElementDefinitionBuilder) AddIdentifier(identifier *datatype.Identifier) *ResearchElementDefinitionBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ResearchElementDefinitionBuilder) SetStatus(status string) *ResearchElementDefinitionBuilder {
	b.Data["status"] = status
	return b
}

func (b *ResearchElementDefinitionBuilder) SetSubject(reference string, display ...string) *ResearchElementDefinitionBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *ResearchElementDefinitionBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ResearchElementDefinitionBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *ResearchElementDefinitionBuilder) SetCode(system, code, display string) *ResearchElementDefinitionBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *ResearchElementDefinitionBuilder) SetEncounter(reference string) *ResearchElementDefinitionBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
