package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type SubscriptionBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewSubscriptionBuilder() *SubscriptionBuilder {
	b := &SubscriptionBuilder{ResourceType: "Subscription", Data: make(map[string]interface{})}
	return b
}

func (b *SubscriptionBuilder) SetId(id string) *SubscriptionBuilder { b.Data["id"] = id; return b }
func (b *SubscriptionBuilder) AddIdentifier(id *datatype.Identifier) *SubscriptionBuilder {
	if _, ok := b.Data["identifier"]; !ok { b.Data["identifier"] = make([]interface{}, 0) }
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), id.ToArray())
	return b
}
func (b *SubscriptionBuilder) SetStatus(status string) *SubscriptionBuilder { b.Data["status"] = status; return b }
func (b *SubscriptionBuilder) SetTopic(topic string) *SubscriptionBuilder { b.Data["topic"] = topic; return b }
func (b *SubscriptionBuilder) SetContact(system string, value string) *SubscriptionBuilder {
	b.Data["contact"] = []interface{}{map[string]interface{}{"system": system, "value": value}}
	return b
}
func (b *SubscriptionBuilder) SetChannel(channelType string, endpoint string) *SubscriptionBuilder {
	b.Data["channel"] = map[string]interface{}{"type": channelType, "endpoint": endpoint}
	return b
}
func (b *SubscriptionBuilder) SetReason(reason string) *SubscriptionBuilder { b.Data["reason"] = reason; return b }
func (b *SubscriptionBuilder) SetError(error string) *SubscriptionBuilder { b.Data["error"] = map[string]interface{}{"message": error}; return b }
func (b *SubscriptionBuilder) Build() map[string]interface{} { return b.Data }
