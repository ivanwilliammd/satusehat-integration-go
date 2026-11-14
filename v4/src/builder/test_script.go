package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type TestScriptBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewTestScriptBuilder() *TestScriptBuilder {
	b := &TestScriptBuilder{ResourceType: "TestScript", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "TestScript"
	return b
}

func (b *TestScriptBuilder) SetId(id string) *TestScriptBuilder {
	b.Data["id"] = id
	return b
}

func (b *TestScriptBuilder) AddIdentifier(identifier *datatype.Identifier) *TestScriptBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *TestScriptBuilder) SetStatus(status string) *TestScriptBuilder {
	b.Data["status"] = status
	return b
}

func (b *TestScriptBuilder) SetSubject(reference string, display ...string) *TestScriptBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *TestScriptBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *TestScriptBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *TestScriptBuilder) SetCode(system, code, display string) *TestScriptBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *TestScriptBuilder) SetEncounter(reference string) *TestScriptBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
