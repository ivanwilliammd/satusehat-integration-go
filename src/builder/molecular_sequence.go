package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type MolecularSequenceBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewMolecularSequenceBuilder() *MolecularSequenceBuilder {
    b := &MolecularSequenceBuilder{ResourceType: "MolecularSequence", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "MolecularSequence"
    return b
}

func (b *MolecularSequenceBuilder) setId(id string) *MolecularSequenceBuilder {
    b.Data["id"] = id
    return b
}

func (b *MolecularSequenceBuilder) addIdentifier(identifier *datatype.Identifier) *MolecularSequenceBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *MolecularSequenceBuilder) setStatus(status string) *MolecularSequenceBuilder {
    b.Data["status"] = status
    return b
}

func (b *MolecularSequenceBuilder) setSubject(reference string) *MolecularSequenceBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *MolecularSequenceBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *MolecularSequenceBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
