package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ConsentBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewConsentBuilder() *ConsentBuilder {
	b := &ConsentBuilder{ResourceType: "Consent", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Consent"
	return b
}

func (b *ConsentBuilder) SetId(id string) *ConsentBuilder {
	b.Data["id"] = id
	return b
}

func (b *ConsentBuilder) AddIdentifier(identifier *datatype.Identifier) *ConsentBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ConsentBuilder) SetStatus(status string) *ConsentBuilder {
	b.Data["status"] = status
	return b
}

func (b *ConsentBuilder) SetSubject(reference string, display ...string) *ConsentBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *ConsentBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ConsentBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *ConsentBuilder) SetCode(system, code, display string) *ConsentBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *ConsentBuilder) SetEncounter(reference string) *ConsentBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
