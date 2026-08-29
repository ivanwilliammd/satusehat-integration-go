package builder

import (
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

// DiagnosticReportBuilder builds FHIR DiagnosticReport payload
type DiagnosticReportBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewDiagnosticReportBuilder() *DiagnosticReportBuilder {
	b := &DiagnosticReportBuilder{ResourceType: "DiagnosticReport", Data: make(map[string]interface{})}
	return b
}

func (b *DiagnosticReportBuilder) setId(id string) *DiagnosticReportBuilder {
	b.Data["id"] = id
	return b
}

func (b *DiagnosticReportBuilder) addIdentifier(identifier *datatype.Identifier) *DiagnosticReportBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), identifier.ToArray())
	return b
}

func (b *DiagnosticReportBuilder) setStatus(status string) *DiagnosticReportBuilder {
	b.Data["status"] = status
	return b
}

func (b *DiagnosticReportBuilder) addCategory(category *datatype.CodeableConcept) *DiagnosticReportBuilder {
	if _, ok := b.Data["category"]; !ok {
		b.Data["category"] = make([]interface{}, 0)
	}
	b.Data["category"] = append(b.Data["category"].([]interface{}), category.ToArray())
	return b
}

func (b *DiagnosticReportBuilder) setCode(code *datatype.CodeableConcept) *DiagnosticReportBuilder {
	b.Data["code"] = code.ToArray()
	return b
}

func (b *DiagnosticReportBuilder) setSubject(subject *datatype.Reference) *DiagnosticReportBuilder {
	b.Data["subject"] = subject.ToArray()
	return b
}

func (b *DiagnosticReportBuilder) setEncounter(encounter *datatype.Reference) *DiagnosticReportBuilder {
	b.Data["encounter"] = encounter.ToArray()
	return b
}

func (b *DiagnosticReportBuilder) setEffectiveDateTime(dateTime string) *DiagnosticReportBuilder {
	b.Data["effectiveDateTime"] = dateTime
	return b
}

func (b *DiagnosticReportBuilder) setIssued(issued string) *DiagnosticReportBuilder {
	b.Data["issued"] = issued
	return b
}

func (b *DiagnosticReportBuilder) addPerformer(performer *datatype.Reference) *DiagnosticReportBuilder {
	if _, ok := b.Data["performer"]; !ok {
		b.Data["performer"] = make([]interface{}, 0)
	}
	b.Data["performer"] = append(b.Data["performer"].([]interface{}), performer.ToArray())
	return b
}

func (b *DiagnosticReportBuilder) addResult(result *datatype.Reference) *DiagnosticReportBuilder {
	if _, ok := b.Data["result"]; !ok {
		b.Data["result"] = make([]interface{}, 0)
	}
	b.Data["result"] = append(b.Data["result"].([]interface{}), result.ToArray())
	return b
}

func (b *DiagnosticReportBuilder) addSpecimen(specimen *datatype.Reference) *DiagnosticReportBuilder {
	if _, ok := b.Data["specimen"]; !ok {
		b.Data["specimen"] = make([]interface{}, 0)
	}
	b.Data["specimen"] = append(b.Data["specimen"].([]interface{}), specimen.ToArray())
	return b
}

func (b *DiagnosticReportBuilder) addConclusionCode(conclusionCode *datatype.CodeableConcept) *DiagnosticReportBuilder {
	if _, ok := b.Data["conclusionCode"]; !ok {
		b.Data["conclusionCode"] = make([]interface{}, 0)
	}
	b.Data["conclusionCode"] = append(b.Data["conclusionCode"].([]interface{}), conclusionCode.ToArray())
	return b
}

func (b *DiagnosticReportBuilder) addBasedOn(basedOn *datatype.Reference) *DiagnosticReportBuilder {
	if _, ok := b.Data["basedOn"]; !ok {
		b.Data["basedOn"] = make([]interface{}, 0)
	}
	b.Data["basedOn"] = append(b.Data["basedOn"].([]interface{}), basedOn.ToArray())
	return b
}

func (b *DiagnosticReportBuilder) setConclusion(conclusion string) *DiagnosticReportBuilder {
	b.Data["conclusion"] = conclusion
	return b
}

func (b *DiagnosticReportBuilder) addExtension(url string, value interface{}, valueType string) *DiagnosticReportBuilder {
	extension := map[string]interface{}{"url": url}
	if valueType != "" {
		capitalized := strings.ToUpper(valueType[:1]) + valueType[1:]
		extension["value"+capitalized] = value
	} else {
		extension["valueString"] = value
	}
	if _, ok := b.Data["extension"]; !ok {
		b.Data["extension"] = make([]interface{}, 0)
	}
	b.Data["extension"] = append(b.Data["extension"].([]interface{}), extension)
	return b
}

func (b *DiagnosticReportBuilder) Build() map[string]interface{} {
	return b.Data
}
