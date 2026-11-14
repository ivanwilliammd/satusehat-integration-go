package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type GuidanceResponseBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewGuidanceResponseBuilder() *GuidanceResponseBuilder {
	b := &GuidanceResponseBuilder{ResourceType: "GuidanceResponse", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "GuidanceResponse"
	return b
}

func (b *GuidanceResponseBuilder) SetId(id string) *GuidanceResponseBuilder {
	b.Data["id"] = id
	return b
}

func (b *GuidanceResponseBuilder) AddIdentifier(identifier *datatype.Identifier) *GuidanceResponseBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *GuidanceResponseBuilder) SetStatus(status string) *GuidanceResponseBuilder {
	b.Data["status"] = status
	return b
}

func (b *GuidanceResponseBuilder) SetSubject(reference string) *GuidanceResponseBuilder {
	b.Data["subject"] = datatype.Reference{}.ToArray()
	return b
}

func (b *GuidanceResponseBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *GuidanceResponseBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}
