package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type EvidenceVariableBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewEvidenceVariableBuilder() *EvidenceVariableBuilder {
	b := &EvidenceVariableBuilder{ResourceType: "EvidenceVariable", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "EvidenceVariable"
	return b
}

func (b *EvidenceVariableBuilder) SetId(id string) *EvidenceVariableBuilder {
	b.Data["id"] = id
	return b
}

func (b *EvidenceVariableBuilder) AddIdentifier(identifier *datatype.Identifier) *EvidenceVariableBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *EvidenceVariableBuilder) SetStatus(status string) *EvidenceVariableBuilder {
	b.Data["status"] = status
	return b
}

func (b *EvidenceVariableBuilder) SetSubject(reference string) *EvidenceVariableBuilder {
	b.Data["subject"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

func (b *EvidenceVariableBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *EvidenceVariableBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}
