package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type EvidenceVariableBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewEvidenceVariableBuilder() *EvidenceVariableBuilder {
	b := &EvidenceVariableBuilder{ResourceType: "EvidenceVariable", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "EvidenceVariable"
	return b
}

func (b *EvidenceVariableBuilder) SetId(id string) *EvidenceVariableBuilder {
	b.Data["id"] = id
	return b
}

func (b *EvidenceVariableBuilder) AddIdentifier(identifier *datatype.Identifier) *EvidenceVariableBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *EvidenceVariableBuilder) SetStatus(status string) *EvidenceVariableBuilder {
	b.Data["status"] = status
	return b
}

func (b *EvidenceVariableBuilder) SetSubject(reference string, display ...string) *EvidenceVariableBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *EvidenceVariableBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *EvidenceVariableBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *EvidenceVariableBuilder) SetCode(system, code, display string) *EvidenceVariableBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *EvidenceVariableBuilder) SetEncounter(reference string) *EvidenceVariableBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
