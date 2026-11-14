package builder

import (
	"encoding/json"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

// CarePlanBuilder builds a FHIR R4 CarePlan resource.
//
// Method-for-method counterpart of PayloadBuilderCarePlan in the PHP SDK. addActivity takes the
// detail block as a ready-made map and wraps it in {detail: …}, because PHP does not shape it.
//
// @link https://www.hl7.org/fhir/careplan.html
type CarePlanBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewCarePlanBuilder creates an empty CarePlan with its resourceType set.
func NewCarePlanBuilder() *CarePlanBuilder {
	b := &CarePlanBuilder{ResourceType: "CarePlan", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

func (b *CarePlanBuilder) set(k string, v interface{}) *CarePlanBuilder {
	b.Data[k] = v
	return b
}

func (b *CarePlanBuilder) push(k string, v interface{}) *CarePlanBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

// SetId sets the logical id.
func (b *CarePlanBuilder) SetId(id string) *CarePlanBuilder { return b.set("id", id) }

// AddIdentifier appends an identifier.
func (b *CarePlanBuilder) AddIdentifier(identifier *datatype.Identifier) *CarePlanBuilder {
	return b.push("identifier", identifier.ToArray())
}

// SetStatus sets status (draft | active | on-hold | revoked | completed | entered-in-error | unknown).
func (b *CarePlanBuilder) SetStatus(status string) *CarePlanBuilder { return b.set("status", status) }

// SetIntent sets intent.
func (b *CarePlanBuilder) SetIntent(intent string) *CarePlanBuilder { return b.set("intent", intent) }

// AddCategory appends a category.
func (b *CarePlanBuilder) AddCategory(category *datatype.CodeableConcept) *CarePlanBuilder {
	return b.push("category", category.ToArray())
}

// SetTitle sets title.
func (b *CarePlanBuilder) SetTitle(title string) *CarePlanBuilder { return b.set("title", title) }

// SetDescription sets description.
func (b *CarePlanBuilder) SetDescription(description string) *CarePlanBuilder {
	return b.set("description", description)
}

// SetSubject sets subject.
func (b *CarePlanBuilder) SetSubject(subject *datatype.Reference) *CarePlanBuilder {
	return b.set("subject", subject.ToArray())
}

// SetEncounter sets encounter.
func (b *CarePlanBuilder) SetEncounter(encounter *datatype.Reference) *CarePlanBuilder {
	return b.set("encounter", encounter.ToArray())
}

// SetPeriod sets period.
func (b *CarePlanBuilder) SetPeriod(period *datatype.Period) *CarePlanBuilder {
	return b.set("period", period.ToArray())
}

// SetAuthor sets author.
func (b *CarePlanBuilder) SetAuthor(author *datatype.Reference) *CarePlanBuilder {
	return b.set("author", author.ToArray())
}

// AddContributor appends a contributor.
func (b *CarePlanBuilder) AddContributor(contributor *datatype.Reference) *CarePlanBuilder {
	return b.push("contributor", contributor.ToArray())
}

// AddAddressee appends an addressee.
func (b *CarePlanBuilder) AddAddressee(addressee *datatype.Reference) *CarePlanBuilder {
	return b.push("addressee", addressee.ToArray())
}

// AddSupportingInfo appends a supportingInfo reference.
func (b *CarePlanBuilder) AddSupportingInfo(info *datatype.Reference) *CarePlanBuilder {
	return b.push("supportingInfo", info.ToArray())
}

// AddGoal appends a goal reference.
func (b *CarePlanBuilder) AddGoal(goal *datatype.Reference) *CarePlanBuilder {
	return b.push("goal", goal.ToArray())
}

// AddActivity appends an activity, wrapping the given detail block in {detail: …}.
func (b *CarePlanBuilder) AddActivity(detail map[string]interface{}) *CarePlanBuilder {
	return b.push("activity", map[string]interface{}{"detail": detail})
}

// Build returns the payload.
func (b *CarePlanBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *CarePlanBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
