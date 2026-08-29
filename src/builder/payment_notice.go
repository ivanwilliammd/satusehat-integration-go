package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type PaymentNoticeBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewPaymentNoticeBuilder() *PaymentNoticeBuilder {
    b := &PaymentNoticeBuilder{ResourceType: "PaymentNotice", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "PaymentNotice"
    return b
}

func (b *PaymentNoticeBuilder) setId(id string) *PaymentNoticeBuilder {
    b.Data["id"] = id
    return b
}

func (b *PaymentNoticeBuilder) addIdentifier(identifier *datatype.Identifier) *PaymentNoticeBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *PaymentNoticeBuilder) setStatus(status string) *PaymentNoticeBuilder {
    b.Data["status"] = status
    return b
}

func (b *PaymentNoticeBuilder) setSubject(reference string) *PaymentNoticeBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *PaymentNoticeBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *PaymentNoticeBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
