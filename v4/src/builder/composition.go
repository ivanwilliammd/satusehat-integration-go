package builder

import (
	"encoding/json"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

// CompositionBuilder builds a FHIR R4 Composition resource.
//
// Method-for-method counterpart of PayloadBuilderComposition in the PHP SDK. Sections are passed as
// ready-made maps: the PHP builder does not shape them, so neither does this port.
//
// @link https://www.hl7.org/fhir/composition.html
type CompositionBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewCompositionBuilder creates an empty Composition with its resourceType set.
func NewCompositionBuilder() *CompositionBuilder {
	b := &CompositionBuilder{ResourceType: "Composition", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

func (b *CompositionBuilder) set(k string, v interface{}) *CompositionBuilder {
	b.Data[k] = v
	return b
}

func (b *CompositionBuilder) push(k string, v interface{}) *CompositionBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

// SetId sets the logical id.
func (b *CompositionBuilder) SetId(id string) *CompositionBuilder { return b.set("id", id) }

// AddIdentifier appends an identifier.
func (b *CompositionBuilder) AddIdentifier(identifier *datatype.Identifier) *CompositionBuilder {
	return b.push("identifier", identifier.ToArray())
}

// SetStatus sets status (preliminary | final | amended | entered-in-error).
func (b *CompositionBuilder) SetStatus(status string) *CompositionBuilder {
	return b.set("status", status)
}

// SetType sets type.
func (b *CompositionBuilder) SetType(typ *datatype.CodeableConcept) *CompositionBuilder {
	return b.set("type", typ.ToArray())
}

// AddCategory appends a category.
func (b *CompositionBuilder) AddCategory(category *datatype.CodeableConcept) *CompositionBuilder {
	return b.push("category", category.ToArray())
}

// SetTitle sets title.
func (b *CompositionBuilder) SetTitle(title string) *CompositionBuilder { return b.set("title", title) }

// SetDate sets date.
func (b *CompositionBuilder) SetDate(date string) *CompositionBuilder { return b.set("date", date) }

// AddAuthor appends an author.
func (b *CompositionBuilder) AddAuthor(author *datatype.Reference) *CompositionBuilder {
	return b.push("author", author.ToArray())
}

// SetSubject sets subject.
func (b *CompositionBuilder) SetSubject(subject *datatype.Reference) *CompositionBuilder {
	return b.set("subject", subject.ToArray())
}

// SetEncounter sets encounter.
func (b *CompositionBuilder) SetEncounter(encounter *datatype.Reference) *CompositionBuilder {
	return b.set("encounter", encounter.ToArray())
}

// SetCustodian sets custodian.
func (b *CompositionBuilder) SetCustodian(custodian *datatype.Reference) *CompositionBuilder {
	return b.set("custodian", custodian.ToArray())
}

// AddAttester appends an attester (mode + party).
func (b *CompositionBuilder) AddAttester(mode string, party *datatype.Reference) *CompositionBuilder {
	return b.push("attester", map[string]interface{}{
		"mode":  mode,
		"party": party.ToArray(),
	})
}

// AddSection appends a section exactly as given.
func (b *CompositionBuilder) AddSection(section map[string]interface{}) *CompositionBuilder {
	return b.push("section", section)
}

// Build returns the payload.
func (b *CompositionBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *CompositionBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
