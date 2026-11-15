package builder

import (
	"encoding/json"
	"fmt"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type MolecularSequenceBuilder struct {
	err          error
	ResourceType string
	Data         map[string]interface{}
}

func NewMolecularSequenceBuilder() *MolecularSequenceBuilder {
	b := &MolecularSequenceBuilder{ResourceType: "MolecularSequence", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "MolecularSequence"
	return b
}

func (b *MolecularSequenceBuilder) SetId(id string) *MolecularSequenceBuilder {
	b.Data["id"] = id
	return b
}

func (b *MolecularSequenceBuilder) AddIdentifier(identifier *datatype.Identifier) *MolecularSequenceBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *MolecularSequenceBuilder) SetStatus(status string) *MolecularSequenceBuilder {
	b.Data["status"] = status
	return b
}

func (b *MolecularSequenceBuilder) SetSubject(reference string, display ...string) *MolecularSequenceBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

// Build returns the payload. The canonical builder throws when coordinateSystem is missing; here that is
// recorded as an error, because a Go method cannot refuse to return a map.
func (b *MolecularSequenceBuilder) Build() map[string]interface{} {
	if _, ada := b.Data["coordinateSystem"]; !ada {
		b.err = fmt.Errorf("MolecularSequence.coordinateSystem wajib diisi")
	}
	return b.Data
}

func (b *MolecularSequenceBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetPatient sets patient from a Reference, as the PHP builder does.
func (b *MolecularSequenceBuilder) SetPatient(patient *datatype.Reference) *MolecularSequenceBuilder {
	b.Data["patient"] = patient.ToArray()
	return b
}

// SetType sets type and rejects a value the canonical builder rejects, rather than writing it.
func (b *MolecularSequenceBuilder) SetType(kind string) *MolecularSequenceBuilder {
	switch kind {
	case "aa", "dna", "rna":
		b.Data["type"] = kind
	default:
		b.err = fmt.Errorf("invalid type: must be one of aa, dna, rna")
	}
	return b
}

// Err returns the first validation error, as the canonical builder throws on the same condition.
func (b *MolecularSequenceBuilder) Err() error {
	return b.err
}

// SetText sets the narrative block.
func (b *MolecularSequenceBuilder) SetText(status, div string) *MolecularSequenceBuilder {
	b.Data["text"] = map[string]interface{}{"status": status, "div": div}
	return b
}

// SetCoordinateSystem sets coordinateSystem. The canonical builder refuses to build without it, so Build
// records that as an error here; see Build.
func (b *MolecularSequenceBuilder) SetCoordinateSystem(coordinateSystem int) *MolecularSequenceBuilder {
	b.Data["coordinateSystem"] = coordinateSystem
	return b
}

// SetSpecimen sets specimen from a reference.
func (b *MolecularSequenceBuilder) SetSpecimen(specimen *datatype.Reference) *MolecularSequenceBuilder {
	b.Data["specimen"] = specimen.ToArray()
	return b
}

// SetDevice sets device from a reference.
func (b *MolecularSequenceBuilder) SetDevice(device *datatype.Reference) *MolecularSequenceBuilder {
	b.Data["device"] = device.ToArray()
	return b
}

// SetPerformer sets performer from a reference.
func (b *MolecularSequenceBuilder) SetPerformer(performer *datatype.Reference) *MolecularSequenceBuilder {
	b.Data["performer"] = performer.ToArray()
	return b
}

// SetReferenceSeq sets referenceSeq with its optional strand and window.
func (b *MolecularSequenceBuilder) SetReferenceSeq(referenceSeqId *datatype.CodeableConcept, strand *string, windowStart, windowEnd *int, orientation *string) *MolecularSequenceBuilder {
	refSeq := map[string]interface{}{"referenceSeqId": referenceSeqId.ToArray()}
	if strand != nil {
		refSeq["strand"] = *strand
	}
	if windowStart != nil {
		refSeq["windowStart"] = *windowStart
	}
	if windowEnd != nil {
		refSeq["windowEnd"] = *windowEnd
	}
	if orientation != nil {
		refSeq["orientation"] = *orientation
	}
	b.Data["referenceSeq"] = refSeq
	return b
}

// AddVariant appends a variant with its alleles and optional pointer.
func (b *MolecularSequenceBuilder) AddVariant(start, end int, observedAllele, referenceAllele string, variantPointer *datatype.Reference) *MolecularSequenceBuilder {
	variant := map[string]interface{}{
		"start":           start,
		"end":             end,
		"observedAllele":  observedAllele,
		"referenceAllele": referenceAllele,
	}
	if variantPointer != nil {
		variant["variantPointer"] = variantPointer.ToArray()
	}
	pushPath(b.Data, "variant", variant)
	return b
}
