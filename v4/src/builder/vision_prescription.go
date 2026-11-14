package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type VisionPrescriptionBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewVisionPrescriptionBuilder() *VisionPrescriptionBuilder {
	b := &VisionPrescriptionBuilder{ResourceType: "VisionPrescription", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "VisionPrescription"
	return b
}

func (b *VisionPrescriptionBuilder) SetId(id string) *VisionPrescriptionBuilder {
	b.Data["id"] = id
	return b
}

func (b *VisionPrescriptionBuilder) AddIdentifier(identifier *datatype.Identifier) *VisionPrescriptionBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *VisionPrescriptionBuilder) SetStatus(status string) *VisionPrescriptionBuilder {
	b.Data["status"] = status
	return b
}

func (b *VisionPrescriptionBuilder) SetSubject(reference string, display ...string) *VisionPrescriptionBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *VisionPrescriptionBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *VisionPrescriptionBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *VisionPrescriptionBuilder) SetCode(system, code, display string) *VisionPrescriptionBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *VisionPrescriptionBuilder) SetEncounter(reference string) *VisionPrescriptionBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
