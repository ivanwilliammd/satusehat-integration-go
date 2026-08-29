package builder

type MessageHeaderBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewMessageHeaderBuilder() *MessageHeaderBuilder {
	b := &MessageHeaderBuilder{ResourceType: "MessageHeader", Data: make(map[string]interface{})}
	return b
}

func (b *MessageHeaderBuilder) SetId(id string) *MessageHeaderBuilder { b.Data["id"] = id; return b }
func (b *MessageHeaderBuilder) SetEventUri(uri string) *MessageHeaderBuilder { b.Data["eventUri"] = uri; return b }
func (b *MessageHeaderBuilder) SetDestination(name string, uri string) *MessageHeaderBuilder {
	b.Data["destination"] = []interface{}{map[string]interface{}{"name": name, "endpoint": uri}}
	return b
}
func (b *MessageHeaderBuilder) SetSender(ref string) *MessageHeaderBuilder { b.Data["sender"] = map[string]interface{}{"reference": ref}; return b }
func (b *MessageHeaderBuilder) SetSource(name string, uri string) *MessageHeaderBuilder {
	b.Data["source"] = map[string]interface{}{"name": name, "endpoint": uri}
	return b
}
func (b *MessageHeaderBuilder) SetFocus(ref string) *MessageHeaderBuilder { b.Data["focus"] = []interface{}{map[string]interface{}{"reference": ref}}; return b }
func (b *MessageHeaderBuilder) Build() map[string]interface{} { return b.Data }
