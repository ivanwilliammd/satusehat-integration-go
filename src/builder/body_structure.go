package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type BodyStructureBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewBodyStructureBuilder() *BodyStructureBuilder {
	b := &BodyStructureBuilder{ResourceType: "BodyStructure", Data: make(map[string]interface{})}
	return b
}

func (b *BodyStructureBuilder) SetId(id string) *BodyStructureBuilder { b.Data["id"] = id; return b }

func (b *BodyStructureBuilder) AddIdentifier(id *datatype.Identifier) *BodyStructureBuilder {
	if _, ok := b.Data["identifier"]; !ok { b.Data["identifier"] = make([]interface{}, 0) }
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), id.ToArray())
	return b
}

func (b *BodyStructureBuilder) SetLocation(location *datatype.CodeableConcept) *BodyStructureBuilder { b.Data["location"] = location.ToArray(); return b }
func (b *BodyStructureBuilder) SetMorphology(morph *datatype.CodeableConcept) *BodyStructureBuilder { b.Data["morphology"] = morph.ToArray(); return b }
func (b *BodyStructureBuilder) SetDescription(description string) *BodyStructureBuilder { b.Data["description"] = description; return b }
func (b *BodyStructureBuilder) SetActive(active bool) *BodyStructureBuilder { b.Data["active"] = active; return b }
func (b *BodyStructureBuilder) SetPatient(patientRef string) *BodyStructureBuilder { b.Data["patient"] = map[string]interface{}{"reference": patientRef}; return b }

func (b *BodyStructureBuilder) Build() map[string]interface{} { return b.Data }
