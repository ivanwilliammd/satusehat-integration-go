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

func (b *EnrollmentRequestBuilder) SetSubject(reference string, display ...string) *EnrollmentRequestBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *EnrollmentRequestBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *EnrollmentRequestBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *EnrollmentRequestBuilder) SetCode(system, code, display string) *EnrollmentRequestBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *EnrollmentRequestBuilder) SetEncounter(reference string) *EnrollmentRequestBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
