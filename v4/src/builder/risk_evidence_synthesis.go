package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type RiskEvidenceSynthesisBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewRiskEvidenceSynthesisBuilder() *RiskEvidenceSynthesisBuilder {
	b := &RiskEvidenceSynthesisBuilder{ResourceType: "RiskEvidenceSynthesis", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "RiskEvidenceSynthesis"
	return b
}

func (b *RiskEvidenceSynthesisBuilder) SetId(id string) *RiskEvidenceSynthesisBuilder {
	b.Data["id"] = id
	return b
}

func (b *RiskEvidenceSynthesisBuilder) AddIdentifier(identifier *datatype.Identifier) *RiskEvidenceSynthesisBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *RiskEvidenceSynthesisBuilder) SetStatus(status string) *RiskEvidenceSynthesisBuilder {
	b.Data["status"] = status
	return b
}

func (b *RiskEvidenceSynthesisBuilder) SetSubject(reference string, display ...string) *RiskEvidenceSynthesisBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *RiskEvidenceSynthesisBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *RiskEvidenceSynthesisBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *RiskEvidenceSynthesisBuilder) SetCode(system, code, display string) *RiskEvidenceSynthesisBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *RiskEvidenceSynthesisBuilder) SetEncounter(reference string) *RiskEvidenceSynthesisBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
