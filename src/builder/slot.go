package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type SlotBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewSlotBuilder() *SlotBuilder {
    b := &SlotBuilder{ResourceType: "Slot", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Slot"
    return b
}

func (b *SlotBuilder) setId(id string) *SlotBuilder {
    b.Data["id"] = id
    return b
}

func (b *SlotBuilder) addIdentifier(identifier *datatype.Identifier) *SlotBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *SlotBuilder) setStatus(status string) *SlotBuilder {
    b.Data["status"] = status
    return b
}

func (b *SlotBuilder) setSubject(reference string) *SlotBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *SlotBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *SlotBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
