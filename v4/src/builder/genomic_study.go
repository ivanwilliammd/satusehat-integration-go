package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type GenomicStudyBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewGenomicStudyBuilder() *GenomicStudyBuilder {
	b := &GenomicStudyBuilder{ResourceType: "GenomicStudy", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "GenomicStudy"
	return b
}

func (b *GenomicStudyBuilder) SetId(id string) *GenomicStudyBuilder {
	b.Data["id"] = id
	return b
}

func (b *GenomicStudyBuilder) AddIdentifier(identifier *datatype.Identifier) *GenomicStudyBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *GenomicStudyBuilder) SetStatus(status string) *GenomicStudyBuilder {
	b.Data["status"] = status
	return b
}

func (b *GenomicStudyBuilder) SetSubject(reference string, display ...string) *GenomicStudyBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *GenomicStudyBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *GenomicStudyBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetEncounter sets encounter from a Reference, as the PHP builder does.
func (b *GenomicStudyBuilder) SetEncounter(encounter *datatype.Reference) *GenomicStudyBuilder {
	b.Data["encounter"] = encounter.ToArray()
	return b
}
