package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ExampleScenarioBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewExampleScenarioBuilder() *ExampleScenarioBuilder {
	b := &ExampleScenarioBuilder{ResourceType: "ExampleScenario", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "ExampleScenario"
	return b
}

func (b *ExampleScenarioBuilder) SetId(id string) *ExampleScenarioBuilder {
	b.Data["id"] = id
	return b
}

func (b *ExampleScenarioBuilder) AddIdentifier(identifier *datatype.Identifier) *ExampleScenarioBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ExampleScenarioBuilder) SetStatus(status string) *ExampleScenarioBuilder {
	b.Data["status"] = status
	return b
}

func (b *ExampleScenarioBuilder) SetSubject(reference string, display ...string) *ExampleScenarioBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *ExampleScenarioBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ExampleScenarioBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *ExampleScenarioBuilder) SetCode(system, code, display string) *ExampleScenarioBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *ExampleScenarioBuilder) SetEncounter(reference string) *ExampleScenarioBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
