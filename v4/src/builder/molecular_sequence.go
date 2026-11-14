package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type MolecularSequenceBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewMolecularSequenceBuilder() *MolecularSequenceBuilder {
	b := &MolecularSequenceBuilder{ResourceType: "MolecularSequence", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "MolecularSequence"
	return b
}

func (b *MolecularSequenceBuilder) SetId(id string) *MolecularSequenceBuilder {
	b.Data["id"] = id
	return b
}

func (b *MolecularSequenceBuilder) AddIdentifier(identifier *datatype.Identifier) *MolecularSequenceBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *MolecularSequenceBuilder) SetStatus(status string) *MolecularSequenceBuilder {
	b.Data["status"] = status
	return b
}

func (b *MolecularSequenceBuilder) SetSubject(reference string, display ...string) *MolecularSequenceBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *MolecularSequenceBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *MolecularSequenceBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetPatient sets patient from a Reference, as the PHP builder does.
func (b *MolecularSequenceBuilder) SetPatient(patient *datatype.Reference) *MolecularSequenceBuilder {
	b.Data["patient"] = patient.ToArray()
	return b
}
