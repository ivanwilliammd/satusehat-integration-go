package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type EpisodeOfCareBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewEpisodeOfCareBuilder() *EpisodeOfCareBuilder {
	b := &EpisodeOfCareBuilder{ResourceType: "EpisodeOfCare", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "EpisodeOfCare"
	return b
}

func (b *EpisodeOfCareBuilder) SetId(id string) *EpisodeOfCareBuilder {
	b.Data["id"] = id
	return b
}

func (b *EpisodeOfCareBuilder) AddIdentifier(identifier *datatype.Identifier) *EpisodeOfCareBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *EpisodeOfCareBuilder) SetStatus(status string) *EpisodeOfCareBuilder {
	b.Data["status"] = status
	return b
}

func (b *EpisodeOfCareBuilder) SetSubject(reference string, display ...string) *EpisodeOfCareBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *EpisodeOfCareBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *EpisodeOfCareBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetPatient sets patient from a Reference, as the PHP builder does.
func (b *EpisodeOfCareBuilder) SetPatient(patient *datatype.Reference) *EpisodeOfCareBuilder {
	b.Data["patient"] = patient.ToArray()
	return b
}
