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

func (b *ContractBuilder) SetSubject(reference string, display ...string) *ContractBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *ContractBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ContractBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *ContractBuilder) SetCode(system, code, display string) *ContractBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *ContractBuilder) SetEncounter(reference string) *ContractBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
