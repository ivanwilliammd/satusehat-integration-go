package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type PatientBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewPatientBuilder() *PatientBuilder {
	b := &PatientBuilder{ResourceType: "Patient", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Patient"
	return b
}

func (b *PatientBuilder) SetId(id string) *PatientBuilder {
	b.Data["id"] = id
	return b
}

func (b *PatientBuilder) AddIdentifier(identifier *datatype.Identifier) *PatientBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *PatientBuilder) SetStatus(status string) *PatientBuilder {
	b.Data["status"] = status
	return b
}

func (b *PatientBuilder) SetSubject(reference string) *PatientBuilder {
	b.Data["subject"] = datatype.Reference{}.ToArray()
	return b
}

func (b *PatientBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *PatientBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}
