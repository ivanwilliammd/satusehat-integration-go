package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type EnrollmentRequestBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewEnrollmentRequestBuilder() *EnrollmentRequestBuilder {
	b := &EnrollmentRequestBuilder{ResourceType: "EnrollmentRequest", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "EnrollmentRequest"
	return b
}

func (b *EnrollmentRequestBuilder) SetId(id string) *EnrollmentRequestBuilder {
	b.Data["id"] = id
	return b
}

func (b *EnrollmentRequestBuilder) AddIdentifier(identifier *datatype.Identifier) *EnrollmentRequestBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *EnrollmentRequestBuilder) SetStatus(status string) *EnrollmentRequestBuilder {
	b.Data["status"] = status
	return b
}

func (b *EnrollmentRequestBuilder) SetSubject(reference string) *EnrollmentRequestBuilder {
	b.Data["subject"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

func (b *EnrollmentRequestBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *EnrollmentRequestBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}
