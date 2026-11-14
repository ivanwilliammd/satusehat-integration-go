package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ContractBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewContractBuilder() *ContractBuilder {
	b := &ContractBuilder{ResourceType: "Contract", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Contract"
	return b
}

func (b *ContractBuilder) SetId(id string) *ContractBuilder {
	b.Data["id"] = id
	return b
}

func (b *ContractBuilder) AddIdentifier(identifier *datatype.Identifier) *ContractBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ContractBuilder) SetStatus(status string) *ContractBuilder {
	b.Data["status"] = status
	return b
}

func (b *ContractBuilder) SetSubject(reference string) *ContractBuilder {
	b.Data["subject"] = datatype.Reference{}.ToArray()
	return b
}

func (b *ContractBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ContractBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}
