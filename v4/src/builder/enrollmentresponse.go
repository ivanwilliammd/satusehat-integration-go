package builder

import (
	"encoding/json"
)

type EnrollmentResponseBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewEnrollmentResponseBuilder() *EnrollmentResponseBuilder {
	b := &EnrollmentResponseBuilder{ResourceType: "EnrollmentResponse", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "EnrollmentResponse"
	return b
}

func (b *EnrollmentResponseBuilder) SetId(id string) *EnrollmentResponseBuilder {
	b.Data["id"] = id
	return b
}

func (b *EnrollmentResponseBuilder) AddIdentifier(system, value string) *EnrollmentResponseBuilder {
	ids, _ := b.Data["identifier"].([]map[string]string)
	ids = append(ids, map[string]string{"system": system, "value": value})
	b.Data["identifier"] = ids
	return b
}

func (b *EnrollmentResponseBuilder) SetStatus(status string) *EnrollmentResponseBuilder {
	b.Data["status"] = status
	return b
}

func (b *EnrollmentResponseBuilder) SetRequest(reference string, display ...string) *EnrollmentResponseBuilder {
	ref := map[string]interface{}{"reference": reference}
	if len(display) > 0 && display[0] != "" {
		ref["display"] = display[0]
	}
	b.Data["request"] = ref
	return b
}

func (b *EnrollmentResponseBuilder) SetOutcome(system, code, display string) *EnrollmentResponseBuilder {
	b.Data["outcome"] = map[string]interface{}{
		"coding": []map[string]string{{"system": system, "code": code, "display": display}},
	}
	return b
}

func (b *EnrollmentResponseBuilder) SetDisposition(value string) *EnrollmentResponseBuilder {
	b.Data["disposition"] = value
	return b
}

func (b *EnrollmentResponseBuilder) SetCreated(value string) *EnrollmentResponseBuilder {
	b.Data["created"] = value
	return b
}

func (b *EnrollmentResponseBuilder) Build() map[string]interface{} {
	clean := make(map[string]interface{})
	for k, v := range b.Data {
		if v != nil {
			clean[k] = v
		}
	}
	return clean
}

func (b *EnrollmentResponseBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Build())
}
