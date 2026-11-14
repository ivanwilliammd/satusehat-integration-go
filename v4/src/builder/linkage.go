package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type LinkageBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewLinkageBuilder() *LinkageBuilder {
	b := &LinkageBuilder{ResourceType: "Linkage", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Linkage"
	return b
}

func (b *LinkageBuilder) SetId(id string) *LinkageBuilder {
	b.Data["id"] = id
	return b
}

func (b *LinkageBuilder) AddIdentifier(identifier *datatype.Identifier) *LinkageBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *LinkageBuilder) SetStatus(status string) *LinkageBuilder {
	b.Data["status"] = status
	return b
}

func (b *LinkageBuilder) SetSubject(reference string, display ...string) *LinkageBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *LinkageBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *LinkageBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *LinkageBuilder) SetCode(system, code, display string) *LinkageBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *LinkageBuilder) SetEncounter(reference string) *LinkageBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
