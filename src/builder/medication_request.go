package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type MedicationRequestBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewMedicationRequestBuilder() *MedicationRequestBuilder {
    b := &MedicationRequestBuilder{ResourceType: "MedicationRequest", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "MedicationRequest"
    return b
}

func (b *MedicationRequestBuilder) setId(id string) *MedicationRequestBuilder {
    b.Data["id"] = id
    return b
}

func (b *MedicationRequestBuilder) addIdentifier(identifier *datatype.Identifier) *MedicationRequestBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *MedicationRequestBuilder) setStatus(status string) *MedicationRequestBuilder {
    b.Data["status"] = status
    return b
}

func (b *MedicationRequestBuilder) setSubject(reference string) *MedicationRequestBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *MedicationRequestBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *MedicationRequestBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
