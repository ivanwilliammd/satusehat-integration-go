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

func (b *BinaryBuilder) SetSubject(reference string) *BinaryBuilder {
	b.Data["subject"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

func (b *BinaryBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *BinaryBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}
