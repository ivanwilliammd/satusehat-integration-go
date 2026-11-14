package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type MeasureBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewMeasureBuilder() *MeasureBuilder {
	b := &MeasureBuilder{ResourceType: "Measure", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Measure"
	return b
}

func (b *MeasureBuilder) SetId(id string) *MeasureBuilder {
	b.Data["id"] = id
	return b
}

func (b *MeasureBuilder) AddIdentifier(identifier *datatype.Identifier) *MeasureBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *MeasureBuilder) SetStatus(status string) *MeasureBuilder {
	b.Data["status"] = status
	return b
}

func (b *MeasureBuilder) SetSubject(reference string) *MeasureBuilder {
	b.Data["subject"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

func (b *MeasureBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *MeasureBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}
