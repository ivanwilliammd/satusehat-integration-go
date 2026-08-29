package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type VisionPrescriptionBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewVisionPrescriptionBuilder() *VisionPrescriptionBuilder {
    b := &VisionPrescriptionBuilder{ResourceType: "VisionPrescription", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "VisionPrescription"
    return b
}

func (b *VisionPrescriptionBuilder) setId(id string) *VisionPrescriptionBuilder {
    b.Data["id"] = id
    return b
}

func (b *VisionPrescriptionBuilder) addIdentifier(identifier *datatype.Identifier) *VisionPrescriptionBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *VisionPrescriptionBuilder) setStatus(status string) *VisionPrescriptionBuilder {
    b.Data["status"] = status
    return b
}

func (b *VisionPrescriptionBuilder) setSubject(reference string) *VisionPrescriptionBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *VisionPrescriptionBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *VisionPrescriptionBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
