package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type SubscriptionBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewSubscriptionBuilder() *SubscriptionBuilder {
	b := &SubscriptionBuilder{ResourceType: "Subscription", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Subscription"
	return b
}

func (b *SubscriptionBuilder) SetId(id string) *SubscriptionBuilder {
	b.Data["id"] = id
	return b
}

func (b *SubscriptionBuilder) AddIdentifier(identifier *datatype.Identifier) *SubscriptionBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *SubscriptionBuilder) SetStatus(status string) *SubscriptionBuilder {
	b.Data["status"] = status
	return b
}

func (b *SubscriptionBuilder) SetSubject(reference string) *SubscriptionBuilder {
	b.Data["subject"] = datatype.Reference{}.ToArray()
	return b
}

func (b *SubscriptionBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *SubscriptionBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}
