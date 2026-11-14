package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type MediaBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewMediaBuilder() *MediaBuilder {
	b := &MediaBuilder{ResourceType: "Media", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Media"
	return b
}

func (b *MediaBuilder) SetId(id string) *MediaBuilder {
	b.Data["id"] = id
	return b
}

func (b *MediaBuilder) AddIdentifier(identifier *datatype.Identifier) *MediaBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *MediaBuilder) SetStatus(status string) *MediaBuilder {
	b.Data["status"] = status
	return b
}

func (b *MediaBuilder) SetSubject(reference string, display ...string) *MediaBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *MediaBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *MediaBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *MediaBuilder) SetCode(system, code, display string) *MediaBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *MediaBuilder) SetEncounter(reference string) *MediaBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
