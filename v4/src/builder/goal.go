package builder

import (
	"encoding/json"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

// GoalBuilder builds a FHIR R4 Goal resource.
//
// Method-for-method counterpart of PayloadBuilderGoal in the PHP SDK. Targets are appended as
// separate detail[x] entries, so several targets with different shapes can coexist, and statusReason
// is written as a plain string because that is what PHP does (FHIR declares a CodeableConcept there).
//
// @link https://www.hl7.org/fhir/goal.html
type GoalBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewGoalBuilder creates an empty Goal with its resourceType set.
func NewGoalBuilder() *GoalBuilder {
	b := &GoalBuilder{ResourceType: "Goal", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

func (b *GoalBuilder) set(k string, v interface{}) *GoalBuilder {
	b.Data[k] = v
	return b
}

func (b *GoalBuilder) push(k string, v interface{}) *GoalBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

// SetId sets the logical id.
func (b *GoalBuilder) SetId(id string) *GoalBuilder { return b.set("id", id) }

// AddIdentifier appends an identifier.
func (b *GoalBuilder) AddIdentifier(identifier *datatype.Identifier) *GoalBuilder {
	return b.push("identifier", identifier.ToArray())
}

// SetLifecycleStatus sets lifecycleStatus (proposed | planned | accepted | active | on-hold | completed | cancelled | entered-in-error | rejected).
func (b *GoalBuilder) SetLifecycleStatus(status string) *GoalBuilder {
	return b.set("lifecycleStatus", status)
}

// SetDescription sets description.
func (b *GoalBuilder) SetDescription(description *datatype.CodeableConcept) *GoalBuilder {
	return b.set("description", description.ToArray())
}

// SetSubject sets subject.
func (b *GoalBuilder) SetSubject(subject *datatype.Reference) *GoalBuilder {
	return b.set("subject", subject.ToArray())
}

// SetStartDateTime sets startDateTime.
func (b *GoalBuilder) SetStartDateTime(v string) *GoalBuilder { return b.set("startDateTime", v) }

// SetStartAge sets startAge.
func (b *GoalBuilder) SetStartAge(v *datatype.Age) *GoalBuilder {
	return b.set("startAge", v.ToArray())
}

// SetStartPeriod sets startPeriod.
func (b *GoalBuilder) SetStartPeriod(v *datatype.Period) *GoalBuilder {
	return b.set("startPeriod", v.ToArray())
}

// AddTargetDetailQuantity appends a target carrying detailQuantity.
func (b *GoalBuilder) AddTargetDetailQuantity(detail *datatype.Quantity) *GoalBuilder {
	return b.push("target", map[string]interface{}{"detailQuantity": detail.ToArray()})
}

// AddTargetDetailCodeableConcept appends a target carrying detailCodeableConcept.
func (b *GoalBuilder) AddTargetDetailCodeableConcept(detail *datatype.CodeableConcept) *GoalBuilder {
	return b.push("target", map[string]interface{}{"detailCodeableConcept": detail.ToArray()})
}

// AddTargetRange appends a target carrying detailRange.
func (b *GoalBuilder) AddTargetRange(detail *datatype.Range) *GoalBuilder {
	return b.push("target", map[string]interface{}{"detailRange": detail.ToArray()})
}

// SetStatusReason sets statusReason as a plain string, as PHP does.
func (b *GoalBuilder) SetStatusReason(statusReason string) *GoalBuilder {
	return b.set("statusReason", statusReason)
}

// AddNote appends a note, either as a typed Annotation or as plain text.
func (b *GoalBuilder) AddNote(note interface{}) *GoalBuilder {
	switch v := note.(type) {
	case *datatype.Annotation:
		return b.push("note", v.ToArray())
	case string:
		return b.push("note", map[string]interface{}{"text": v})
	}
	return b
}

// AddOutcomeReference appends an outcomeReference.
func (b *GoalBuilder) AddOutcomeReference(ref *datatype.Reference) *GoalBuilder {
	return b.push("outcomeReference", ref.ToArray())
}

// Build returns the payload.
func (b *GoalBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *GoalBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
