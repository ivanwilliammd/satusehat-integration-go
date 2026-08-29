package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ImagingStudyBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewImagingStudyBuilder() *ImagingStudyBuilder {
    b := &ImagingStudyBuilder{ResourceType: "ImagingStudy", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "ImagingStudy"
    return b
}

func (b *ImagingStudyBuilder) setId(id string) *ImagingStudyBuilder {
    b.Data["id"] = id
    return b
}

func (b *ImagingStudyBuilder) addIdentifier(identifier *datatype.Identifier) *ImagingStudyBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *ImagingStudyBuilder) setStatus(status string) *ImagingStudyBuilder {
    b.Data["status"] = status
    return b
}

func (b *ImagingStudyBuilder) setSubject(reference string) *ImagingStudyBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *ImagingStudyBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *ImagingStudyBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
