package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type SubstanceProteinBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewSubstanceProteinBuilder() *SubstanceProteinBuilder {
	b := &SubstanceProteinBuilder{ResourceType: "SubstanceProtein", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "SubstanceProtein"
	return b
}

func (b *SubstanceProteinBuilder) SetId(id string) *SubstanceProteinBuilder {
	b.Data["id"] = id
	return b
}

func (b *SubstanceProteinBuilder) AddIdentifier(identifier *datatype.Identifier) *SubstanceProteinBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *SubstanceProteinBuilder) SetStatus(status string) *SubstanceProteinBuilder {
	b.Data["status"] = status
	return b
}

func (b *SubstanceProteinBuilder) SetSubject(reference string) *SubstanceProteinBuilder {
	b.Data["subject"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

func (b *SubstanceProteinBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *SubstanceProteinBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}
