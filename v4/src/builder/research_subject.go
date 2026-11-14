package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ResearchSubjectBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewResearchSubjectBuilder() *ResearchSubjectBuilder {
	b := &ResearchSubjectBuilder{ResourceType: "ResearchSubject", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "ResearchSubject"
	return b
}

func (b *ResearchSubjectBuilder) SetId(id string) *ResearchSubjectBuilder {
	b.Data["id"] = id
	return b
}

func (b *ResearchSubjectBuilder) AddIdentifier(identifier *datatype.Identifier) *ResearchSubjectBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ResearchSubjectBuilder) SetStatus(status string) *ResearchSubjectBuilder {
	b.Data["status"] = status
	return b
}

func (b *ResearchSubjectBuilder) SetSubject(reference string, display ...string) *ResearchSubjectBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *ResearchSubjectBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ResearchSubjectBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *ResearchSubjectBuilder) SetCode(system, code, display string) *ResearchSubjectBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *ResearchSubjectBuilder) SetEncounter(reference string) *ResearchSubjectBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
