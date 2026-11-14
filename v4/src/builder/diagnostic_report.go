package builder

import (
	"encoding/json"
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

// DiagnosticReportBuilder builds a FHIR R4 DiagnosticReport resource.
//
// Method-for-method counterpart of PayloadBuilderDiagnosticReport in the PHP SDK. Note that
// effective[x] is limited to effectiveDateTime there, and that performer and result are lists.
//
// @link https://www.hl7.org/fhir/diagnosticreport.html
type DiagnosticReportBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewDiagnosticReportBuilder creates an empty DiagnosticReport with its resourceType set.
func NewDiagnosticReportBuilder() *DiagnosticReportBuilder {
	b := &DiagnosticReportBuilder{ResourceType: "DiagnosticReport", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

func (b *DiagnosticReportBuilder) set(k string, v interface{}) *DiagnosticReportBuilder {
	b.Data[k] = v
	return b
}

func (b *DiagnosticReportBuilder) push(k string, v interface{}) *DiagnosticReportBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

// SetId sets the logical id.
func (b *DiagnosticReportBuilder) SetId(id string) *DiagnosticReportBuilder { return b.set("id", id) }

// AddIdentifier appends an identifier.
func (b *DiagnosticReportBuilder) AddIdentifier(identifier *datatype.Identifier) *DiagnosticReportBuilder {
	return b.push("identifier", identifier.ToArray())
}

// SetStatus sets status (registered | partial | preliminary | final | amended | corrected | appended | cancelled | entered-in-error | unknown).
func (b *DiagnosticReportBuilder) SetStatus(status string) *DiagnosticReportBuilder {
	return b.set("status", status)
}

// AddCategory appends a category.
func (b *DiagnosticReportBuilder) AddCategory(category *datatype.CodeableConcept) *DiagnosticReportBuilder {
	return b.push("category", category.ToArray())
}

// SetCode sets code.
func (b *DiagnosticReportBuilder) SetCode(code *datatype.CodeableConcept) *DiagnosticReportBuilder {
	return b.set("code", code.ToArray())
}

// SetSubject sets subject.
func (b *DiagnosticReportBuilder) SetSubject(subject *datatype.Reference) *DiagnosticReportBuilder {
	return b.set("subject", subject.ToArray())
}

// SetEncounter sets encounter.
func (b *DiagnosticReportBuilder) SetEncounter(encounter *datatype.Reference) *DiagnosticReportBuilder {
	return b.set("encounter", encounter.ToArray())
}

// SetEffectiveDateTime sets effectiveDateTime.
func (b *DiagnosticReportBuilder) SetEffectiveDateTime(dateTime string) *DiagnosticReportBuilder {
	return b.set("effectiveDateTime", dateTime)
}

// SetIssued sets issued.
func (b *DiagnosticReportBuilder) SetIssued(instant string) *DiagnosticReportBuilder {
	return b.set("issued", instant)
}

// AddPerformer appends a performer.
func (b *DiagnosticReportBuilder) AddPerformer(performer *datatype.Reference) *DiagnosticReportBuilder {
	return b.push("performer", performer.ToArray())
}

// AddResult appends a result reference.
func (b *DiagnosticReportBuilder) AddResult(result *datatype.Reference) *DiagnosticReportBuilder {
	return b.push("result", result.ToArray())
}

// AddSpecimen appends a specimen reference.
func (b *DiagnosticReportBuilder) AddSpecimen(specimen *datatype.Reference) *DiagnosticReportBuilder {
	return b.push("specimen", specimen.ToArray())
}

// AddConclusionCode appends a conclusionCode.
func (b *DiagnosticReportBuilder) AddConclusionCode(conclusionCode *datatype.CodeableConcept) *DiagnosticReportBuilder {
	return b.push("conclusionCode", conclusionCode.ToArray())
}

// AddBasedOn appends a basedOn reference.
func (b *DiagnosticReportBuilder) AddBasedOn(basedOn *datatype.Reference) *DiagnosticReportBuilder {
	return b.push("basedOn", basedOn.ToArray())
}

// SetConclusion sets conclusion.
func (b *DiagnosticReportBuilder) SetConclusion(conclusion string) *DiagnosticReportBuilder {
	return b.set("conclusion", conclusion)
}

// AddExtension appends an extension. When valueType is empty the value is written as valueString.
func (b *DiagnosticReportBuilder) AddExtension(url string, value interface{}, valueType string) *DiagnosticReportBuilder {
	extension := map[string]interface{}{"url": url}
	if valueType == "" {
		extension["valueString"] = value
	} else {
		extension["value"+strings.ToUpper(valueType[:1])+valueType[1:]] = value
	}
	return b.push("extension", extension)
}

// Build returns the payload.
func (b *DiagnosticReportBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *DiagnosticReportBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
