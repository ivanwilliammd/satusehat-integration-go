package builder

type MolecularSequenceBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewMolecularSequenceBuilder() *MolecularSequenceBuilder {
	b := &MolecularSequenceBuilder{ResourceType: "MolecularSequence", Data: make(map[string]interface{})}
	return b
}

func (b *MolecularSequenceBuilder) SetId(id string) *MolecularSequenceBuilder { b.Data["id"] = id; return b }
func (b *MolecularSequenceBuilder) SetType(type_ string) *MolecularSequenceBuilder { b.Data["type"] = type_; return b }
func (b *MolecularSequenceBuilder) SetPatient(patientRef string) *MolecularSequenceBuilder { b.Data["patient"] = map[string]interface{}{"reference": patientRef}; return b }
func (b *MolecularSequenceBuilder) SetSpecimen(specimenRef string) *MolecularSequenceBuilder { b.Data["specimen"] = map[string]interface{}{"reference": specimenRef}; return b }
func (b *MolecularSequenceBuilder) Build() map[string]interface{} { return b.Data }
