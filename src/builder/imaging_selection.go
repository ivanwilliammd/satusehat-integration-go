package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type ImagingSelectionBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewImagingSelectionBuilder() *ImagingSelectionBuilder {
    b := &ImagingSelectionBuilder{ResourceType: "ImagingSelection", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "ImagingSelection"
    return b
}

func (b *ImagingSelectionBuilder) setId(id string) *ImagingSelectionBuilder {
    b.Data["id"] = id
    return b
}

func (b *ImagingSelectionBuilder) addIdentifier(identifier *datatype.Identifier) *ImagingSelectionBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *ImagingSelectionBuilder) setStatus(status string) *ImagingSelectionBuilder {
    b.Data["status"] = status
    return b
}

func (b *ImagingSelectionBuilder) setSubject(reference string) *ImagingSelectionBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *ImagingSelectionBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *ImagingSelectionBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
