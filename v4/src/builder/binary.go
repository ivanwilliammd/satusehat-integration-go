package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type BinaryBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewBinaryBuilder() *BinaryBuilder {
	b := &BinaryBuilder{ResourceType: "Binary", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Binary"
	return b
}

func (b *BinaryBuilder) SetId(id string) *BinaryBuilder {
	b.Data["id"] = id
	return b
}

func (b *BinaryBuilder) AddIdentifier(identifier *datatype.Identifier) *BinaryBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *BinaryBuilder) SetStatus(status string) *BinaryBuilder {
	b.Data["status"] = status
	return b
}

func (b *BinaryBuilder) SetSubject(reference string, display ...string) *BinaryBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *BinaryBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *BinaryBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *BinaryBuilder) SetCode(system, code, display string) *BinaryBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *BinaryBuilder) SetEncounter(reference string) *BinaryBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
