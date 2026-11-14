package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type LocationBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewLocationBuilder() *LocationBuilder {
	b := &LocationBuilder{ResourceType: "Location", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Location"
	return b
}

func (b *LocationBuilder) SetId(id string) *LocationBuilder {
	b.Data["id"] = id
	return b
}

func (b *LocationBuilder) AddIdentifier(identifier *datatype.Identifier) *LocationBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *LocationBuilder) SetStatus(status string) *LocationBuilder {
	b.Data["status"] = status
	return b
}

func (b *LocationBuilder) SetSubject(reference string) *LocationBuilder {
	b.Data["subject"] = datatype.Reference{}.ToArray()
	return b
}

func (b *LocationBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *LocationBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}
