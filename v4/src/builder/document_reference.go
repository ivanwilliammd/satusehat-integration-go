package builder

import (
	"encoding/json"
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

// DocumentReferenceBuilder builds a FHIR R4 DocumentReference resource.
//
// Method-for-method counterpart of PayloadBuilderDocumentReference in the PHP SDK. Note that addType
// pushes into a list, while FHIR declares type as 0..1 — kept as-is so payloads stay identical.
//
// @link https://www.hl7.org/fhir/documentreference.html
type DocumentReferenceBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewDocumentReferenceBuilder creates an empty DocumentReference with its resourceType set.
func NewDocumentReferenceBuilder() *DocumentReferenceBuilder {
	b := &DocumentReferenceBuilder{ResourceType: "DocumentReference", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

func (b *DocumentReferenceBuilder) set(k string, v interface{}) *DocumentReferenceBuilder {
	b.Data[k] = v
	return b
}

func (b *DocumentReferenceBuilder) push(k string, v interface{}) *DocumentReferenceBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

// SetId sets the logical id.
func (b *DocumentReferenceBuilder) SetId(id string) *DocumentReferenceBuilder { return b.set("id", id) }

// AddIdentifier appends an identifier.
func (b *DocumentReferenceBuilder) AddIdentifier(identifier *datatype.Identifier) *DocumentReferenceBuilder {
	return b.push("identifier", identifier.ToArray())
}

// SetMasterIdentifier sets masterIdentifier.
func (b *DocumentReferenceBuilder) SetMasterIdentifier(identifier *datatype.Identifier) *DocumentReferenceBuilder {
	return b.set("masterIdentifier", identifier.ToArray())
}

// SetStatus sets status (current | superseded | entered-in-error).
func (b *DocumentReferenceBuilder) SetStatus(status string) *DocumentReferenceBuilder {
	return b.set("status", status)
}

// SetDocStatus sets docStatus (preliminary | final | amended | entered-in-error).
func (b *DocumentReferenceBuilder) SetDocStatus(docStatus string) *DocumentReferenceBuilder {
	return b.set("docStatus", docStatus)
}

// AddType appends a type. PHP pushes into a list although FHIR declares 0..1.
func (b *DocumentReferenceBuilder) AddType(typ *datatype.CodeableConcept) *DocumentReferenceBuilder {
	return b.push("type", typ.ToArray())
}

// AddCategory appends a category.
func (b *DocumentReferenceBuilder) AddCategory(category *datatype.CodeableConcept) *DocumentReferenceBuilder {
	return b.push("category", category.ToArray())
}

// SetSubject sets subject.
func (b *DocumentReferenceBuilder) SetSubject(subject *datatype.Reference) *DocumentReferenceBuilder {
	return b.set("subject", subject.ToArray())
}

// SetDate sets date.
func (b *DocumentReferenceBuilder) SetDate(date string) *DocumentReferenceBuilder {
	return b.set("date", date)
}

// AddAuthor appends an author.
func (b *DocumentReferenceBuilder) AddAuthor(author *datatype.Reference) *DocumentReferenceBuilder {
	return b.push("author", author.ToArray())
}

// SetCustodian sets custodian.
func (b *DocumentReferenceBuilder) SetCustodian(custodian *datatype.Reference) *DocumentReferenceBuilder {
	return b.set("custodian", custodian.ToArray())
}

// AddRelatesTo appends a relatesTo entry (code + target).
func (b *DocumentReferenceBuilder) AddRelatesTo(code string, target *datatype.Reference) *DocumentReferenceBuilder {
	return b.push("relatesTo", map[string]interface{}{
		"code":   code,
		"target": target.ToArray(),
	})
}

// SetDescription sets description.
func (b *DocumentReferenceBuilder) SetDescription(description string) *DocumentReferenceBuilder {
	return b.set("description", description)
}

// AddSecurityLabel appends a securityLabel.
func (b *DocumentReferenceBuilder) AddSecurityLabel(label *datatype.CodeableConcept) *DocumentReferenceBuilder {
	return b.push("securityLabel", label.ToArray())
}

// AddContent appends a content entry. format is optional.
func (b *DocumentReferenceBuilder) AddContent(attachment *datatype.Attachment, format *datatype.CodeableConcept) *DocumentReferenceBuilder {
	content := map[string]interface{}{"attachment": attachment.ToArray()}
	if format != nil {
		content["format"] = format.ToArray()
	}
	return b.push("content", content)
}

// SetContext sets context exactly as given.
func (b *DocumentReferenceBuilder) SetContext(context map[string]interface{}) *DocumentReferenceBuilder {
	return b.set("context", context)
}

// AddExtension appends an extension. When valueType is empty the value is written as valueString.
func (b *DocumentReferenceBuilder) AddExtension(url string, value interface{}, valueType string) *DocumentReferenceBuilder {
	extension := map[string]interface{}{"url": url}
	if valueType == "" {
		extension["valueString"] = value
	} else {
		extension["value"+strings.ToUpper(valueType[:1])+valueType[1:]] = value
	}
	return b.push("extension", extension)
}

// Build returns the payload.
func (b *DocumentReferenceBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *DocumentReferenceBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
