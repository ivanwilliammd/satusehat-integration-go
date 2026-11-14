package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type SearchParameterBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewSearchParameterBuilder() *SearchParameterBuilder {
	b := &SearchParameterBuilder{ResourceType: "SearchParameter", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "SearchParameter"
	return b
}

func (b *SearchParameterBuilder) SetId(id string) *SearchParameterBuilder {
	b.Data["id"] = id
	return b
}

func (b *SearchParameterBuilder) AddIdentifier(identifier *datatype.Identifier) *SearchParameterBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *SearchParameterBuilder) SetStatus(status string) *SearchParameterBuilder {
	b.Data["status"] = status
	return b
}

func (b *SearchParameterBuilder) SetSubject(reference string, display ...string) *SearchParameterBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *SearchParameterBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *SearchParameterBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *SearchParameterBuilder) SetCode(system, code, display string) *SearchParameterBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *SearchParameterBuilder) SetEncounter(reference string) *SearchParameterBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
