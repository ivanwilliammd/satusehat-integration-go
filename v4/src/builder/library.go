package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type LibraryBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewLibraryBuilder() *LibraryBuilder {
	b := &LibraryBuilder{ResourceType: "Library", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Library"
	return b
}

func (b *LibraryBuilder) SetId(id string) *LibraryBuilder {
	b.Data["id"] = id
	return b
}

func (b *LibraryBuilder) AddIdentifier(identifier *datatype.Identifier) *LibraryBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *LibraryBuilder) SetStatus(status string) *LibraryBuilder {
	b.Data["status"] = status
	return b
}

func (b *LibraryBuilder) SetSubject(reference string, display ...string) *LibraryBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *LibraryBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *LibraryBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *LibraryBuilder) SetCode(system, code, display string) *LibraryBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *LibraryBuilder) SetEncounter(reference string) *LibraryBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
