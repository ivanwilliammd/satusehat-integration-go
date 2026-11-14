package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ProvenanceBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewProvenanceBuilder() *ProvenanceBuilder {
	b := &ProvenanceBuilder{ResourceType: "Provenance", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Provenance"
	return b
}

func (b *ProvenanceBuilder) SetId(id string) *ProvenanceBuilder {
	b.Data["id"] = id
	return b
}

func (b *ProvenanceBuilder) AddIdentifier(identifier *datatype.Identifier) *ProvenanceBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ProvenanceBuilder) SetStatus(status string) *ProvenanceBuilder {
	b.Data["status"] = status
	return b
}

func (b *ProvenanceBuilder) SetSubject(reference string, display ...string) *ProvenanceBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *ProvenanceBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ProvenanceBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *ProvenanceBuilder) SetCode(system, code, display string) *ProvenanceBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *ProvenanceBuilder) SetEncounter(reference string) *ProvenanceBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
