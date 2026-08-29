package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type AuditEventBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewAuditEventBuilder() *AuditEventBuilder {
    b := &AuditEventBuilder{ResourceType: "AuditEvent", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "AuditEvent"
    return b
}

func (b *AuditEventBuilder) setId(id string) *AuditEventBuilder {
    b.Data["id"] = id
    return b
}

func (b *AuditEventBuilder) addIdentifier(identifier *datatype.Identifier) *AuditEventBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *AuditEventBuilder) setStatus(status string) *AuditEventBuilder {
    b.Data["status"] = status
    return b
}

func (b *AuditEventBuilder) setSubject(reference string) *AuditEventBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *AuditEventBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *AuditEventBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
