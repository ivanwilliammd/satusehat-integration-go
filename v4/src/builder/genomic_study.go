package builder

import (
	"encoding/json"
	"fmt"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type GenomicStudyBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewGenomicStudyBuilder() *GenomicStudyBuilder {
	b := &GenomicStudyBuilder{ResourceType: "GenomicStudy", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "GenomicStudy"
	return b
}

func (b *GenomicStudyBuilder) SetId(id string) *GenomicStudyBuilder {
	b.Data["id"] = id
	return b
}

func (b *GenomicStudyBuilder) AddIdentifier(identifier *datatype.Identifier) *GenomicStudyBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *GenomicStudyBuilder) SetStatus(status string) *GenomicStudyBuilder {
	b.Data["status"] = status
	return b
}

func (b *GenomicStudyBuilder) SetSubject(reference string, display ...string) *GenomicStudyBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *GenomicStudyBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *GenomicStudyBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetEncounter sets encounter from a Reference, as the PHP builder does.
func (b *GenomicStudyBuilder) SetEncounter(encounter *datatype.Reference) *GenomicStudyBuilder {
	b.Data["encounter"] = encounter.ToArray()
	return b
}

// AddNote accepts an annotation or a plain string, as the PHP builder does; a string becomes {text}.
func (b *GenomicStudyBuilder) AddNote(note interface{}) *GenomicStudyBuilder {
	entri := map[string]interface{}{}
	switch v := note.(type) {
	case string:
		entri["text"] = v
	case *datatype.Annotation:
		entri = v.ToArray()
	case datatype.Annotation:
		entri = v.ToArray()
	default:
		entri["text"] = fmt.Sprint(v)
	}
	daftar, _ := b.Data["note"].([]interface{})
	b.Data["note"] = append(daftar, entri)
	return b
}

// SetDescription sets description, as the PHP builder does.
func (b *GenomicStudyBuilder) SetDescription(description string) *GenomicStudyBuilder {
	b.Data["description"] = description
	return b
}

// SetMetaProfile appends a profile to meta.profile. The canonical builder pushes the path "meta/profile",
// which setPath turns into the nested object meta.profile instead of a literal key.
func (b *GenomicStudyBuilder) SetMetaProfile(profile string) *GenomicStudyBuilder {
	pushPath(b.Data, "meta/profile", profile)
	return b
}

// AddType appends a type coding.
func (b *GenomicStudyBuilder) AddType(studyType *datatype.CodeableConcept) *GenomicStudyBuilder {
	pushPath(b.Data, "type", studyType.ToArray())
	return b
}

// SetStartDate sets startDate.
func (b *GenomicStudyBuilder) SetStartDate(startDate string) *GenomicStudyBuilder {
	b.Data["startDate"] = startDate
	return b
}

// AddBasedOn appends a basedOn reference.
func (b *GenomicStudyBuilder) AddBasedOn(basedOn *datatype.Reference) *GenomicStudyBuilder {
	pushPath(b.Data, "basedOn", basedOn.ToArray())
	return b
}

// SetReferrer sets referrer from a reference.
func (b *GenomicStudyBuilder) SetReferrer(referrer *datatype.Reference) *GenomicStudyBuilder {
	b.Data["referrer"] = referrer.ToArray()
	return b
}

// AddInterpreter appends an interpreter reference.
func (b *GenomicStudyBuilder) AddInterpreter(interpreter *datatype.Reference) *GenomicStudyBuilder {
	pushPath(b.Data, "interpreter", interpreter.ToArray())
	return b
}

// AddReason appends a reason wrapped in a concept, as the canonical builder writes it.
func (b *GenomicStudyBuilder) AddReason(reason *datatype.CodeableConcept) *GenomicStudyBuilder {
	pushPath(b.Data, "reason", map[string]interface{}{"concept": reason.ToArray()})
	return b
}

// AddAnalysis appends an analysis block the caller has already shaped.
func (b *GenomicStudyBuilder) AddAnalysis(analysis map[string]interface{}) *GenomicStudyBuilder {
	pushPath(b.Data, "analysis", analysis)
	return b
}

// AddPerformer appends a performer with the participation role, defaulting to PERF as the canonical builder
// does.
func (b *GenomicStudyBuilder) AddPerformer(actor *datatype.Reference, role string) *GenomicStudyBuilder {
	if role == "" {
		role = "PERF"
	}
	pushPath(b.Data, "performer", map[string]interface{}{
		"actor": actor.ToArray(),
		"role": map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{
				"system":  "http://terminology.hl7.org/3.1.0/CodeSystem-v3-ParticipationType.html",
				"code":    role,
				"display": "Performer",
			}},
		},
	})
	return b
}
