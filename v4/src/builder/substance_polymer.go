package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type SubstancePolymerBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewSubstancePolymerBuilder() *SubstancePolymerBuilder {
	b := &SubstancePolymerBuilder{ResourceType: "SubstancePolymer", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "SubstancePolymer"
	return b
}

func (b *SubstancePolymerBuilder) SetId(id string) *SubstancePolymerBuilder {
	b.Data["id"] = id
	return b
}

func (b *SubstancePolymerBuilder) AddIdentifier(identifier *datatype.Identifier) *SubstancePolymerBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *SubstancePolymerBuilder) SetStatus(status string) *SubstancePolymerBuilder {
	b.Data["status"] = status
	return b
}

func (b *SubstancePolymerBuilder) SetSubject(reference string, display ...string) *SubstancePolymerBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *SubstancePolymerBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *SubstancePolymerBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *SubstancePolymerBuilder) SetCode(system, code, display string) *SubstancePolymerBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *SubstancePolymerBuilder) SetEncounter(reference string) *SubstancePolymerBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
