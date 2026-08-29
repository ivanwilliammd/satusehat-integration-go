package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ImagingManifestBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewImagingManifestBuilder() *ImagingManifestBuilder {
    b := &ImagingManifestBuilder{ResourceType: "ImagingManifest", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "ImagingManifest"
    return b
}

func (b *ImagingManifestBuilder) setId(id string) *ImagingManifestBuilder {
    b.Data["id"] = id
    return b
}

func (b *ImagingManifestBuilder) addIdentifier(identifier *datatype.Identifier) *ImagingManifestBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *ImagingManifestBuilder) setStatus(status string) *ImagingManifestBuilder {
    b.Data["status"] = status
    return b
}

func (b *ImagingManifestBuilder) setSubject(reference string) *ImagingManifestBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *ImagingManifestBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *ImagingManifestBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
