package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type RequestGroupBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewRequestGroupBuilder() *RequestGroupBuilder {
	b := &RequestGroupBuilder{ResourceType: "RequestGroup", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "RequestGroup"
	return b
}

func (b *RequestGroupBuilder) SetId(id string) *RequestGroupBuilder {
	b.Data["id"] = id
	return b
}

func (b *RequestGroupBuilder) AddIdentifier(identifier *datatype.Identifier) *RequestGroupBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *RequestGroupBuilder) SetStatus(status string) *RequestGroupBuilder {
	b.Data["status"] = status
	return b
}

func (b *RequestGroupBuilder) SetSubject(reference string, display ...string) *RequestGroupBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *RequestGroupBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *RequestGroupBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *RequestGroupBuilder) SetCode(system, code, display string) *RequestGroupBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *RequestGroupBuilder) SetEncounter(reference string) *RequestGroupBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
