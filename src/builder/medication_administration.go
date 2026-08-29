package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type MedicationAdministrationBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewMedicationAdministrationBuilder() *MedicationAdministrationBuilder {
    b := &MedicationAdministrationBuilder{ResourceType: "MedicationAdministration", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "MedicationAdministration"
    return b
}

func (b *MedicationAdministrationBuilder) setId(id string) *MedicationAdministrationBuilder {
    b.Data["id"] = id
    return b
}

func (b *MedicationAdministrationBuilder) addIdentifier(identifier *datatype.Identifier) *MedicationAdministrationBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *MedicationAdministrationBuilder) setStatus(status string) *MedicationAdministrationBuilder {
    b.Data["status"] = status
    return b
}

func (b *MedicationAdministrationBuilder) setSubject(reference string) *MedicationAdministrationBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *MedicationAdministrationBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *MedicationAdministrationBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
