package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type DetectedIssueBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewDetectedIssueBuilder() *DetectedIssueBuilder {
	b := &DetectedIssueBuilder{ResourceType: "DetectedIssue", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "DetectedIssue"
	return b
}

func (b *DetectedIssueBuilder) SetId(id string) *DetectedIssueBuilder {
	b.Data["id"] = id
	return b
}

func (b *DetectedIssueBuilder) AddIdentifier(identifier *datatype.Identifier) *DetectedIssueBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *DetectedIssueBuilder) SetStatus(status string) *DetectedIssueBuilder {
	b.Data["status"] = status
	return b
}

func (b *DetectedIssueBuilder) SetSubject(reference string, display ...string) *DetectedIssueBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *DetectedIssueBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *DetectedIssueBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *DetectedIssueBuilder) SetCode(system, code, display string) *DetectedIssueBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *DetectedIssueBuilder) SetEncounter(reference string) *DetectedIssueBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
