package builder

import (
	"encoding/json"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

// AllergyIntoleranceBuilder builds a FHIR R4 AllergyIntolerance resource.
//
// Method-for-method counterpart of PayloadBuilderAllergyIntolerance in the PHP SDK. clinicalStatus
// and verificationStatus are written as coded concepts with the HL7 systems baked in, exactly as
// the PHP builder does, so callers pass a bare code such as "active".
//
// @link https://www.hl7.org/fhir/allergyintolerance.html
type AllergyIntoleranceBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewAllergyIntoleranceBuilder creates an empty AllergyIntolerance with its resourceType set.
func NewAllergyIntoleranceBuilder() *AllergyIntoleranceBuilder {
	b := &AllergyIntoleranceBuilder{ResourceType: "AllergyIntolerance", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

func (b *AllergyIntoleranceBuilder) set(k string, v interface{}) *AllergyIntoleranceBuilder {
	b.Data[k] = v
	return b
}

func (b *AllergyIntoleranceBuilder) push(k string, v interface{}) *AllergyIntoleranceBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

// SetId sets the logical id.
func (b *AllergyIntoleranceBuilder) SetId(id string) *AllergyIntoleranceBuilder {
	return b.set("id", id)
}

// AddIdentifier appends an identifier.
func (b *AllergyIntoleranceBuilder) AddIdentifier(identifier *datatype.Identifier) *AllergyIntoleranceBuilder {
	return b.push("identifier", identifier.ToArray())
}

// SetClinicalStatus sets clinicalStatus with the allergyintolerance-clinical system.
func (b *AllergyIntoleranceBuilder) SetClinicalStatus(status string) *AllergyIntoleranceBuilder {
	return b.set("clinicalStatus", codedConcept("http://terminology.hl7.org/CodeSystem/allergyintolerance-clinical", status))
}

// SetVerificationStatus sets verificationStatus with the allergyintolerance-verification system.
func (b *AllergyIntoleranceBuilder) SetVerificationStatus(status string) *AllergyIntoleranceBuilder {
	return b.set("verificationStatus", codedConcept("http://terminology.hl7.org/CodeSystem/allergyintolerance-verification", status))
}

// SetType sets type (allergy | intolerance).
func (b *AllergyIntoleranceBuilder) SetType(typ string) *AllergyIntoleranceBuilder {
	return b.set("type", typ)
}

// AddCategory appends a category code (food | medication | environment | biologic).
func (b *AllergyIntoleranceBuilder) AddCategory(category string) *AllergyIntoleranceBuilder {
	return b.push("category", category)
}

// SetCriticality sets criticality (low | high | unable-to-assess).
func (b *AllergyIntoleranceBuilder) SetCriticality(criticality string) *AllergyIntoleranceBuilder {
	return b.set("criticality", criticality)
}

// SetCode sets code.
func (b *AllergyIntoleranceBuilder) SetCode(code *datatype.CodeableConcept) *AllergyIntoleranceBuilder {
	return b.set("code", code.ToArray())
}

// SetPatient sets patient.
func (b *AllergyIntoleranceBuilder) SetPatient(patient *datatype.Reference) *AllergyIntoleranceBuilder {
	return b.set("patient", patient.ToArray())
}

// SetEncounter sets encounter.
func (b *AllergyIntoleranceBuilder) SetEncounter(encounter *datatype.Reference) *AllergyIntoleranceBuilder {
	return b.set("encounter", encounter.ToArray())
}

// SetOnsetDateTime sets onsetDateTime.
func (b *AllergyIntoleranceBuilder) SetOnsetDateTime(dateTime string) *AllergyIntoleranceBuilder {
	return b.set("onsetDateTime", dateTime)
}

// SetRecordedDate sets recordedDate.
func (b *AllergyIntoleranceBuilder) SetRecordedDate(dateTime string) *AllergyIntoleranceBuilder {
	return b.set("recordedDate", dateTime)
}

// SetRecorder sets recorder.
func (b *AllergyIntoleranceBuilder) SetRecorder(recorder *datatype.Reference) *AllergyIntoleranceBuilder {
	return b.set("recorder", recorder.ToArray())
}

// SetAsserter sets asserter.
func (b *AllergyIntoleranceBuilder) SetAsserter(asserter *datatype.Reference) *AllergyIntoleranceBuilder {
	return b.set("asserter", asserter.ToArray())
}

// SetLastOccurrence sets lastOccurrence.
func (b *AllergyIntoleranceBuilder) SetLastOccurrence(dateTime string) *AllergyIntoleranceBuilder {
	return b.set("lastOccurrence", dateTime)
}

// AddNote appends a note.
func (b *AllergyIntoleranceBuilder) AddNote(text string) *AllergyIntoleranceBuilder {
	return b.push("note", map[string]interface{}{"text": text})
}

// AddReaction appends a reaction. description, onset, severity, exposureRoute and note are optional:
// pass "" or nil to omit them.
func (b *AllergyIntoleranceBuilder) AddReaction(substance, manifestation *datatype.CodeableConcept, description, onset, severity string, exposureRoute *datatype.CodeableConcept, note string) *AllergyIntoleranceBuilder {
	reaction := map[string]interface{}{
		"substance":     substance.ToArray(),
		"manifestation": []interface{}{manifestation.ToArray()},
	}
	if description != "" {
		reaction["description"] = description
	}
	if onset != "" {
		reaction["onset"] = onset
	}
	if severity != "" {
		reaction["severity"] = severity
	}
	if exposureRoute != nil {
		reaction["exposureRoute"] = exposureRoute.ToArray()
	}
	if note != "" {
		reaction["note"] = []interface{}{map[string]interface{}{"text": note}}
	}
	return b.push("reaction", reaction)
}

// Build returns the payload.
func (b *AllergyIntoleranceBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *AllergyIntoleranceBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
