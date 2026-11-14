package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type CommunicationBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewCommunicationBuilder() *CommunicationBuilder {
	b := &CommunicationBuilder{ResourceType: "Communication", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Communication"
	return b
}

func (b *CommunicationBuilder) SetId(id string) *CommunicationBuilder {
	b.Data["id"] = id
	return b
}

func (b *CommunicationBuilder) AddIdentifier(identifier *datatype.Identifier) *CommunicationBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *CommunicationBuilder) SetStatus(status string) *CommunicationBuilder {
	b.Data["status"] = status
	return b
}

func (b *CommunicationBuilder) SetSubject(reference string, display ...string) *CommunicationBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *CommunicationBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *CommunicationBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *CommunicationBuilder) SetCode(system, code, display string) *CommunicationBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *CommunicationBuilder) SetEncounter(reference string) *CommunicationBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
