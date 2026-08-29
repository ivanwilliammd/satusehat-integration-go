package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type BundleBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewBundleBuilder() *BundleBuilder {
    b := &BundleBuilder{ResourceType: "Bundle", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Bundle"
    return b
}

func (b *BundleBuilder) setId(id string) *BundleBuilder {
    b.Data["id"] = id
    return b
}

func (b *BundleBuilder) addIdentifier(identifier *datatype.Identifier) *BundleBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *BundleBuilder) setStatus(status string) *BundleBuilder {
    b.Data["status"] = status
    return b
}

func (b *BundleBuilder) setSubject(reference string) *BundleBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *BundleBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *BundleBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
