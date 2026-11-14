package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ProcedureBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewProcedureBuilder() *ProcedureBuilder {
	b := &ProcedureBuilder{ResourceType: "Procedure", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Procedure"
	return b
}

func (b *ProcedureBuilder) SetId(id string) *ProcedureBuilder {
	b.Data["id"] = id
	return b
}

func (b *ProcedureBuilder) AddIdentifier(identifier *datatype.Identifier) *ProcedureBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ProcedureBuilder) SetStatus(status string) *ProcedureBuilder {
	b.Data["status"] = status
	return b
}

func (b *ProcedureBuilder) SetSubject(reference string) *ProcedureBuilder {
	b.Data["subject"] = datatype.Reference{}.ToArray()
	return b
}

func (b *ProcedureBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ProcedureBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}
