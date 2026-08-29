package builder

type VisionPrescriptionBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewVisionPrescriptionBuilder() *VisionPrescriptionBuilder {
	b := &VisionPrescriptionBuilder{ResourceType: "VisionPrescription", Data: make(map[string]interface{})}
	return b
}

func (b *VisionPrescriptionBuilder) SetId(id string) *VisionPrescriptionBuilder { b.Data["id"] = id; return b }
func (b *VisionPrescriptionBuilder) SetStatus(status string) *VisionPrescriptionBuilder { b.Data["status"] = status; return b }
func (b *VisionPrescriptionBuilder) SetCreated(created string) *VisionPrescriptionBuilder { b.Data["created"] = created; return b }
func (b *VisionPrescriptionBuilder) SetPatient(patientRef string) *VisionPrescriptionBuilder { b.Data["patient"] = map[string]interface{}{"reference": patientRef}; return b }
func (b *VisionPrescriptionBuilder) SetEncounter(encRef string) *VisionPrescriptionBuilder { b.Data["encounter"] = map[string]interface{}{"reference": encRef}; return b }
func (b *VisionPrescriptionBuilder) SetPrescriber(prescriberRef string) *VisionPrescriptionBuilder { b.Data["prescriber"] = map[string]interface{}{"reference": prescriberRef}; return b }
func (b *VisionPrescriptionBuilder) Build() map[string]interface{} { return b.Data }
