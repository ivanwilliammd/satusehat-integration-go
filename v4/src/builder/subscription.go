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

func (b *SubscriptionBuilder) SetSubject(reference string, display ...string) *SubscriptionBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *SubscriptionBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *SubscriptionBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *SubscriptionBuilder) SetCode(system, code, display string) *SubscriptionBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *SubscriptionBuilder) SetEncounter(reference string) *SubscriptionBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
