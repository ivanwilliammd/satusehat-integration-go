package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type EvidenceBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewEvidenceBuilder() *EvidenceBuilder {
	b := &EvidenceBuilder{ResourceType: "Evidence", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Evidence"
	return b
}

func (b *EvidenceBuilder) SetId(id string) *EvidenceBuilder {
	b.Data["id"] = id
	return b
}

func (b *EvidenceBuilder) AddIdentifier(identifier *datatype.Identifier) *EvidenceBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *EvidenceBuilder) SetStatus(status string) *EvidenceBuilder {
	b.Data["status"] = status
	return b
}

func (b *EvidenceBuilder) SetSubject(reference string, display ...string) *EvidenceBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *EvidenceBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *EvidenceBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *EvidenceBuilder) SetCode(system, code, display string) *EvidenceBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *EvidenceBuilder) SetEncounter(reference string) *EvidenceBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
