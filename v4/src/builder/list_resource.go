package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ListResourceBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewListBuilder() *ListResourceBuilder { return NewListResourceBuilder() }

type ListBuilder = ListResourceBuilder

func NewListResourceBuilder() *ListResourceBuilder {
	b := &ListResourceBuilder{ResourceType: "List", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "List"
	return b
}

func (b *ListResourceBuilder) SetId(id string) *ListResourceBuilder {
	b.Data["id"] = id
	return b
}

func (b *ListResourceBuilder) AddIdentifier(identifier *datatype.Identifier) *ListResourceBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ListResourceBuilder) SetStatus(status string) *ListResourceBuilder {
	b.Data["status"] = status
	return b
}

func (b *ListResourceBuilder) SetSubject(reference string, display ...string) *ListResourceBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *ListResourceBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ListResourceBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *ListResourceBuilder) SetCode(system, code, display string) *ListResourceBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetDate sets date, as the PHP builder does.
func (b *ListResourceBuilder) SetDate(date string) *ListResourceBuilder {
	b.Data["date"] = date
	return b
}
