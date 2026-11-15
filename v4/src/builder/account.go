package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type AccountBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewAccountBuilder() *AccountBuilder {
	b := &AccountBuilder{ResourceType: "Account", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Account"
	return b
}

func (b *AccountBuilder) SetId(id string) *AccountBuilder {
	b.Data["id"] = id
	return b
}

func (b *AccountBuilder) AddIdentifier(identifier *datatype.Identifier) *AccountBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *AccountBuilder) SetStatus(status string) *AccountBuilder {
	b.Data["status"] = status
	return b
}

func (b *AccountBuilder) SetSubject(reference string, display ...string) *AccountBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *AccountBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *AccountBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetType sets type with one coding and a text, as the PHP builder does.
func (b *AccountBuilder) SetType(system, code, display, text string) *AccountBuilder {
	b.Data["type"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
		"text":   text,
	}
	return b
}

// AddCoverage adds a coverage with its priority, as the PHP builder does.
func (b *AccountBuilder) AddCoverage(reference string, priority int) *AccountBuilder {
	daftar, _ := b.Data["coverage"].([]interface{})
	b.Data["coverage"] = append(daftar, map[string]interface{}{
		"coverage": map[string]interface{}{"reference": reference},
		"priority": priority,
	})
	return b
}

// SetDescription sets description, as the PHP builder does.
func (b *AccountBuilder) SetDescription(description string) *AccountBuilder {
	b.Data["description"] = description
	return b
}

// SetName sets name, as the PHP builder does.
func (b *AccountBuilder) SetName(name string) *AccountBuilder {
	b.Data["name"] = name
	return b
}

// AddSubject appends a subject reference with an optional display.
func (b *AccountBuilder) AddSubject(reference string, display *string) *AccountBuilder {
	subject := map[string]interface{}{"reference": reference}
	if display != nil {
		subject["display"] = *display
	}
	pushPath(b.Data, "subject", subject)
	return b
}

// SetServicePeriod sets servicePeriod as one block.
func (b *AccountBuilder) SetServicePeriod(start, end string) *AccountBuilder {
	b.Data["servicePeriod"] = map[string]interface{}{"start": start, "end": end}
	return b
}

// SetOwner sets owner by reference.
func (b *AccountBuilder) SetOwner(reference string) *AccountBuilder {
	b.Data["owner"] = map[string]interface{}{"reference": reference}
	return b
}
