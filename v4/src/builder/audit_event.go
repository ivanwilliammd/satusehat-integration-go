package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type AuditEventBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewAuditEventBuilder() *AuditEventBuilder {
	b := &AuditEventBuilder{ResourceType: "AuditEvent", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "AuditEvent"
	return b
}

func (b *AuditEventBuilder) SetId(id string) *AuditEventBuilder {
	b.Data["id"] = id
	return b
}

func (b *AuditEventBuilder) AddIdentifier(identifier *datatype.Identifier) *AuditEventBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *AuditEventBuilder) SetStatus(status string) *AuditEventBuilder {
	b.Data["status"] = status
	return b
}

func (b *AuditEventBuilder) SetSubject(reference string, display ...string) *AuditEventBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *AuditEventBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *AuditEventBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *AuditEventBuilder) SetCode(system, code, display string) *AuditEventBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *AuditEventBuilder) SetEncounter(reference string) *AuditEventBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
