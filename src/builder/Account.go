package builder

import (
	"encoding/json"

	"github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type AccountBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewAccountBuilder() *AccountBuilder {
	b := &AccountBuilder{ResourceType: "Account", Data: make(map[string]interface{})}
	return b
}

func (b *AccountBuilder) setMetaProfile(profile string) *AccountBuilder {
	if _, ok := b.Data["meta"]; !ok {
		b.Data["meta"] = make(map[string]interface{})
	}
	b.Data["meta"].(map[string]interface{})["profile"] = profile
	return b
}

func (b *AccountBuilder) setId(id string) *AccountBuilder {
	b.Data["id"] = id
	return b
}

func (b *AccountBuilder) addIdentifier(identifier *datatype.Identifier) *AccountBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), identifier.ToArray())
	return b
}

func (b *AccountBuilder) setStatus(status string) *AccountBuilder {
	b.Data["status"] = status
	return b
}

func (b *AccountBuilder) setType(system, code, display, text string) *AccountBuilder {
	b.Data["type"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
		"text":   text,
	}
	return b
}

func (b *AccountBuilder) setName(name string) *AccountBuilder {
	b.Data["name"] = name
	return b
}

func (b *AccountBuilder) addSubject(reference string, display ...string) *AccountBuilder {
	subject := map[string]interface{}{"reference": reference}
	if len(display) > 0 && display[0] != "" {
		subject["display"] = display[0]
	}
	b.Data["subject"] = append(b.Data["subject"].([]interface{}), subject)
	return b
}

func (b *AccountBuilder) setServicePeriod(start, end string) *AccountBuilder {
	b.Data["servicePeriod"] = map[string]interface{}{"start": start, "end": end}
	return b
}

func (b *AccountBuilder) addCoverage(reference string, priority int) *AccountBuilder {
	b.Data["coverage"] = append(b.Data["coverage"].([]interface{}), map[string]interface{}{
		"coverage": map[string]interface{}{"reference": reference},
		"priority": priority,
	})
	return b
}

func (b *AccountBuilder) setOwner(reference string) *AccountBuilder {
	b.Data["owner"] = map[string]interface{}{"reference": reference}
	return b
}

func (b *AccountBuilder) setDescription(description string) *AccountBuilder {
	b.Data["description"] = description
	return b
}

func (b *AccountBuilder) setActive(active bool) *AccountBuilder {
	b.Data["active"] = active
	return b
}

func (b *AccountBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *AccountBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}
