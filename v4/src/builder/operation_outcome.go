package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type OperationOutcomeBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewOperationOutcomeBuilder() *OperationOutcomeBuilder {
	b := &OperationOutcomeBuilder{ResourceType: "OperationOutcome", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "OperationOutcome"
	return b
}

func (b *OperationOutcomeBuilder) SetId(id string) *OperationOutcomeBuilder {
	b.Data["id"] = id
	return b
}

func (b *OperationOutcomeBuilder) AddIdentifier(identifier *datatype.Identifier) *OperationOutcomeBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *OperationOutcomeBuilder) SetStatus(status string) *OperationOutcomeBuilder {
	b.Data["status"] = status
	return b
}

func (b *OperationOutcomeBuilder) SetSubject(reference string, display ...string) *OperationOutcomeBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *OperationOutcomeBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *OperationOutcomeBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *OperationOutcomeBuilder) SetCode(system, code, display string) *OperationOutcomeBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *OperationOutcomeBuilder) SetEncounter(reference string) *OperationOutcomeBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
