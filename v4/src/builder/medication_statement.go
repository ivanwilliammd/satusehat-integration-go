package builder

import (
	"encoding/json"
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

// MedicationStatementBuilder builds a FHIR R4 MedicationStatement resource.
//
// Method-for-method counterpart of PayloadBuilderMedicationStatement in the PHP SDK. References are
// given as id or full reference: a bare id is prefixed with the resource type, while absolute
// references (urn:, http://, https://) and values already containing "/" are kept as they are.
//
// @link https://www.hl7.org/fhir/medicationstatement.html
type MedicationStatementBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewMedicationStatementBuilder creates an empty MedicationStatement with its resourceType set.
func NewMedicationStatementBuilder() *MedicationStatementBuilder {
	b := &MedicationStatementBuilder{ResourceType: "MedicationStatement", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

func (b *MedicationStatementBuilder) set(k string, v interface{}) *MedicationStatementBuilder {
	b.Data[k] = v
	return b
}

func (b *MedicationStatementBuilder) push(k string, v interface{}) *MedicationStatementBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

// absolut meniru preg_match('/^(urn:|https?:\/\//') pada PHP: true bila referensi absolut.
func absolut(v string) bool {
	return strings.HasPrefix(v, "urn:") || strings.HasPrefix(v, "http://") || strings.HasPrefix(v, "https://")
}

// refID membentuk "Tipe/id" bila yang diberikan hanya id; referensi absolut dibiarkan.
func refID(typ, v string) string {
	if absolut(v) || strings.Contains(v, "/") {
		return v
	}
	return typ + "/" + v
}

func refWithDisplay(typ, v, display string) map[string]interface{} {
	m := map[string]interface{}{"reference": refID(typ, v)}
	if display != "" {
		m["display"] = display
	}
	return m
}

// SetId sets the logical id.
func (b *MedicationStatementBuilder) SetId(id string) *MedicationStatementBuilder {
	return b.set("id", id)
}

// AddIdentifier appends an identifier.
func (b *MedicationStatementBuilder) AddIdentifier(identifier *datatype.Identifier) *MedicationStatementBuilder {
	return b.push("identifier", identifier.ToArray())
}

// SetStatus sets status (active | completed | entered-in-error | intended | stopped | on-hold | unknown | not-taken).
func (b *MedicationStatementBuilder) SetStatus(status string) *MedicationStatementBuilder {
	return b.set("status", status)
}

// AddStatusReason appends a statusReason. display defaults to code, as in PHP.
func (b *MedicationStatementBuilder) AddStatusReason(code, display, system string) *MedicationStatementBuilder {
	if display == "" {
		display = code
	}
	coding := map[string]interface{}{"code": code, "display": display}
	if system != "" {
		coding["system"] = system
	}
	return b.push("statusReason", map[string]interface{}{
		"coding": []interface{}{coding},
	})
}

// SetCategory sets category.
func (b *MedicationStatementBuilder) SetCategory(category *datatype.CodeableConcept) *MedicationStatementBuilder {
	return b.set("category", category.ToArray())
}

// SetMedicationCodeableConcept sets medicationCodeableConcept.
func (b *MedicationStatementBuilder) SetMedicationCodeableConcept(medication *datatype.CodeableConcept) *MedicationStatementBuilder {
	return b.set("medicationCodeableConcept", medication.ToArray())
}

// SetMedicationReference sets medicationReference; a bare id becomes Medication/<id>.
func (b *MedicationStatementBuilder) SetMedicationReference(reference, display string) *MedicationStatementBuilder {
	return b.set("medicationReference", refWithDisplay("Medication", reference, display))
}

// SetSubject sets subject; a bare id becomes Patient/<id>.
func (b *MedicationStatementBuilder) SetSubject(reference, display string) *MedicationStatementBuilder {
	return b.set("subject", refWithDisplay("Patient", reference, display))
}

// SetContext sets context; a bare id becomes Encounter/<id>.
func (b *MedicationStatementBuilder) SetContext(reference, display string) *MedicationStatementBuilder {
	return b.set("context", refWithDisplay("Encounter", reference, display))
}

// SetDateAsserted sets dateAsserted.
func (b *MedicationStatementBuilder) SetDateAsserted(dateAsserted string) *MedicationStatementBuilder {
	return b.set("dateAsserted", dateAsserted)
}

// SetEffectiveDateTime sets effectiveDateTime.
func (b *MedicationStatementBuilder) SetEffectiveDateTime(effectiveDateTime string) *MedicationStatementBuilder {
	return b.set("effectiveDateTime", effectiveDateTime)
}

// SetEffectivePeriod sets effectivePeriod.
func (b *MedicationStatementBuilder) SetEffectivePeriod(effectivePeriod *datatype.Period) *MedicationStatementBuilder {
	return b.set("effectivePeriod", effectivePeriod.ToArray())
}

// SetInformationSource sets informationSource; a bare id becomes Patient/<id>.
func (b *MedicationStatementBuilder) SetInformationSource(reference, display string) *MedicationStatementBuilder {
	return b.set("informationSource", refWithDisplay("Patient", reference, display))
}

// SetDerivedFrom appends a derivedFrom reference. The value is written as given, like PHP.
func (b *MedicationStatementBuilder) SetDerivedFrom(derivedFrom string) *MedicationStatementBuilder {
	return b.push("derivedFrom", map[string]interface{}{"reference": derivedFrom})
}

// SetReasonCode appends a reasonCode.
func (b *MedicationStatementBuilder) SetReasonCode(reason *datatype.CodeableConcept) *MedicationStatementBuilder {
	return b.push("reasonCode", reason.ToArray())
}

// SetReasonReference appends a reasonReference.
func (b *MedicationStatementBuilder) SetReasonReference(reason *datatype.Reference) *MedicationStatementBuilder {
	return b.push("reasonReference", reason.ToArray())
}

// AddNote appends a note.
func (b *MedicationStatementBuilder) AddNote(text string) *MedicationStatementBuilder {
	return b.push("note", map[string]interface{}{"text": text})
}

// AddContained appends a contained resource.
func (b *MedicationStatementBuilder) AddContained(resource map[string]interface{}) *MedicationStatementBuilder {
	return b.push("contained", resource)
}

// AddDosageInstruction appends a dosage entry built from flat arguments, matching the PHP signature.
// Note the element name: dosage, not dosageInstruction.
func (b *MedicationStatementBuilder) AddDosageInstruction(text string, frequency, period int, periodUnit string) *MedicationStatementBuilder {
	dosage := map[string]interface{}{}
	if text != "" {
		dosage["text"] = text
	}
	repeat := map[string]interface{}{}
	if frequency != 0 {
		dosage["sequence"] = frequency
		repeat["frequency"] = frequency
	}
	if period != 0 {
		repeat["period"] = period
	}
	if periodUnit != "" {
		repeat["periodUnit"] = periodUnit
	}
	if len(repeat) > 0 {
		dosage["timing"] = map[string]interface{}{"repeat": repeat}
	}
	return b.push("dosage", dosage)
}

// AddExtension appends an extension. When valueType is empty the value is written as valueString.
func (b *MedicationStatementBuilder) AddExtension(url string, value interface{}, valueType string) *MedicationStatementBuilder {
	extension := map[string]interface{}{"url": url}
	if valueType == "" {
		extension["valueString"] = value
	} else {
		extension["value"+strings.ToUpper(valueType[:1])+valueType[1:]] = value
	}
	return b.push("extension", extension)
}

// Build returns the payload.
func (b *MedicationStatementBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *MedicationStatementBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
