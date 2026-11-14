package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type CommunicationRequestBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewCommunicationRequestBuilder() *CommunicationRequestBuilder {
	b := &CommunicationRequestBuilder{ResourceType: "CommunicationRequest", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "CommunicationRequest"
	return b
}

func (b *CommunicationRequestBuilder) SetId(id string) *CommunicationRequestBuilder {
	b.Data["id"] = id
	return b
}

func (b *CommunicationRequestBuilder) AddIdentifier(identifier *datatype.Identifier) *CommunicationRequestBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *CommunicationRequestBuilder) SetStatus(status string) *CommunicationRequestBuilder {
	b.Data["status"] = status
	return b
}

func (b *CommunicationRequestBuilder) SetSubject(reference string, display ...string) *CommunicationRequestBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *CommunicationRequestBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *CommunicationRequestBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *CommunicationRequestBuilder) SetCode(system, code, display string) *CommunicationRequestBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *CommunicationRequestBuilder) SetEncounter(reference string) *CommunicationRequestBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
