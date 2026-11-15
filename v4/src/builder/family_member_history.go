package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type FamilyMemberHistoryBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewFamilyMemberHistoryBuilder() *FamilyMemberHistoryBuilder {
	b := &FamilyMemberHistoryBuilder{ResourceType: "FamilyMemberHistory", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "FamilyMemberHistory"
	return b
}

func (b *FamilyMemberHistoryBuilder) SetId(id string) *FamilyMemberHistoryBuilder {
	b.Data["id"] = id
	return b
}

func (b *FamilyMemberHistoryBuilder) AddIdentifier(identifier *datatype.Identifier) *FamilyMemberHistoryBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *FamilyMemberHistoryBuilder) SetStatus(status string) *FamilyMemberHistoryBuilder {
	b.Data["status"] = status
	return b
}

func (b *FamilyMemberHistoryBuilder) SetSubject(reference string, display ...string) *FamilyMemberHistoryBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *FamilyMemberHistoryBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *FamilyMemberHistoryBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetPatient sets patient from a Reference, as the PHP builder does.
func (b *FamilyMemberHistoryBuilder) SetPatient(patient *datatype.Reference) *FamilyMemberHistoryBuilder {
	b.Data["patient"] = patient.ToArray()
	return b
}

// SetDate sets date, as the PHP builder does.
func (b *FamilyMemberHistoryBuilder) SetDate(date string) *FamilyMemberHistoryBuilder {
	b.Data["date"] = date
	return b
}

// SetName sets name when one is given, as the PHP builder does.
func (b *FamilyMemberHistoryBuilder) SetName(name string) *FamilyMemberHistoryBuilder {
	if name != "" {
		b.Data["name"] = name
	}
	return b
}

// SetInstantiatesCanonical sets the instantiatesCanonical reference.
func (b *FamilyMemberHistoryBuilder) SetInstantiatesCanonical(instantiatesCanonical string) *FamilyMemberHistoryBuilder {
	b.Data["instantiatesCanonical"] = instantiatesCanonical
	return b
}

// SetDataAbsentReason sets dataAbsentReason from a CodeableConcept.
func (b *FamilyMemberHistoryBuilder) SetDataAbsentReason(reason *datatype.CodeableConcept) *FamilyMemberHistoryBuilder {
	b.Data["dataAbsentReason"] = reason.ToArray()
	return b
}

// SetRelationship sets relationship from a CodeableConcept.
func (b *FamilyMemberHistoryBuilder) SetRelationship(relationship *datatype.CodeableConcept) *FamilyMemberHistoryBuilder {
	b.Data["relationship"] = relationship.ToArray()
	return b
}

// SetSex sets sex from a CodeableConcept.
func (b *FamilyMemberHistoryBuilder) SetSex(sex *datatype.CodeableConcept) *FamilyMemberHistoryBuilder {
	b.Data["sex"] = sex.ToArray()
	return b
}

// SetBornPeriod sets bornPeriod. The canonical builder takes a CodeableConcept here, and so does this.
func (b *FamilyMemberHistoryBuilder) SetBornPeriod(bornPeriod *datatype.CodeableConcept) *FamilyMemberHistoryBuilder {
	b.Data["bornPeriod"] = bornPeriod.ToArray()
	return b
}

// SetBornDate sets bornDate.
func (b *FamilyMemberHistoryBuilder) SetBornDate(bornDate string) *FamilyMemberHistoryBuilder {
	b.Data["bornDate"] = bornDate
	return b
}

// SetBornString sets bornString.
func (b *FamilyMemberHistoryBuilder) SetBornString(bornString string) *FamilyMemberHistoryBuilder {
	b.Data["bornString"] = bornString
	return b
}

// SetAgeAge writes ageAge and ageUnit, the two keys the canonical builder writes for an age.
func (b *FamilyMemberHistoryBuilder) SetAgeAge(value int, unit string) *FamilyMemberHistoryBuilder {
	b.Data["ageAge"] = value
	b.Data["ageUnit"] = unit
	return b
}

// SetAgeRange writes the value/low/high shape the canonical builder uses for an age range.
func (b *FamilyMemberHistoryBuilder) SetAgeRange(value int, low, high string) *FamilyMemberHistoryBuilder {
	b.Data["ageRange"] = map[string]interface{}{
		"value": value,
		"low":   map[string]interface{}{"value": low},
		"high":  map[string]interface{}{"value": high},
	}
	return b
}

// SetAgeString sets ageString.
func (b *FamilyMemberHistoryBuilder) SetAgeString(ageString string) *FamilyMemberHistoryBuilder {
	b.Data["ageString"] = ageString
	return b
}

// SetDeceasedBoolean sets deceasedBoolean.
func (b *FamilyMemberHistoryBuilder) SetDeceasedBoolean(deceased bool) *FamilyMemberHistoryBuilder {
	b.Data["deceasedBoolean"] = deceased
	return b
}

// SetDeceasedAge writes deceasedAge and deceasedAgeUnit, as the canonical builder does.
func (b *FamilyMemberHistoryBuilder) SetDeceasedAge(value int, unit string) *FamilyMemberHistoryBuilder {
	b.Data["deceasedAge"] = value
	b.Data["deceasedAgeUnit"] = unit
	return b
}

// SetDeceasedRange writes the value/low/high shape the canonical builder uses.
func (b *FamilyMemberHistoryBuilder) SetDeceasedRange(value int, low, high string) *FamilyMemberHistoryBuilder {
	b.Data["deceasedRange"] = map[string]interface{}{
		"value": value,
		"low":   map[string]interface{}{"value": low},
		"high":  map[string]interface{}{"value": high},
	}
	return b
}

// SetDeceasedDate sets deceasedDate.
func (b *FamilyMemberHistoryBuilder) SetDeceasedDate(deceasedDate string) *FamilyMemberHistoryBuilder {
	b.Data["deceasedDate"] = deceasedDate
	return b
}

// SetDeceasedString sets deceasedString.
func (b *FamilyMemberHistoryBuilder) SetDeceasedString(deceasedString string) *FamilyMemberHistoryBuilder {
	b.Data["deceasedString"] = deceasedString
	return b
}

// SetDeceasedCodeableConcept sets deceased from a CodeableConcept.
func (b *FamilyMemberHistoryBuilder) SetDeceasedCodeableConcept(deceased *datatype.CodeableConcept) *FamilyMemberHistoryBuilder {
	b.Data["deceasedCodeableConcept"] = deceased.ToArray()
	return b
}

// SetReasonCode sets reasonCode, replacing any previous value; AddReasonCode appends.
func (b *FamilyMemberHistoryBuilder) SetReasonCode(reasonCode *datatype.CodeableConcept) *FamilyMemberHistoryBuilder {
	b.Data["reasonCode"] = reasonCode.ToArray()
	return b
}

// AddReasonCode appends a reason coding.
func (b *FamilyMemberHistoryBuilder) AddReasonCode(reasonCode *datatype.CodeableConcept) *FamilyMemberHistoryBuilder {
	pushPath(b.Data, "reasonCode", reasonCode.ToArray())
	return b
}

// AddReasonReference appends a reason reference.
func (b *FamilyMemberHistoryBuilder) AddReasonReference(reasonReference *datatype.Reference) *FamilyMemberHistoryBuilder {
	pushPath(b.Data, "reasonReference", reasonReference.ToArray())
	return b
}

// SetNote appends a note with plain text, as the canonical builder does.
func (b *FamilyMemberHistoryBuilder) SetNote(note string) *FamilyMemberHistoryBuilder {
	pushPath(b.Data, "note", map[string]interface{}{"text": note})
	return b
}

// AddCondition appends a condition with its optional onset, outcome and death contribution.
func (b *FamilyMemberHistoryBuilder) AddCondition(code *datatype.CodeableConcept, onsetString *string, outcome *datatype.CodeableConcept, contributedToDeath *bool) *FamilyMemberHistoryBuilder {
	condition := map[string]interface{}{"code": code.ToArray()}
	if onsetString != nil {
		condition["onsetString"] = *onsetString
	}
	if outcome != nil {
		condition["outcome"] = outcome.ToArray()
	}
	if contributedToDeath != nil {
		condition["contributedToDeath"] = *contributedToDeath
	}
	pushPath(b.Data, "condition", condition)
	return b
}
