package builder

import (
	"encoding/json"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

// SpecimenBuilder builds a FHIR R4 Specimen resource.
//
// Method-for-method counterpart of PayloadBuilderSpecimen in the PHP SDK. The PHP builder addresses
// the nested collection element through path strings ("collection/quantity"), so this port uses the
// shared setPath helper and produces the same nested object.
//
// @link https://www.hl7.org/fhir/specimen.html
type SpecimenBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewSpecimenBuilder creates an empty Specimen with its resourceType set.
func NewSpecimenBuilder() *SpecimenBuilder {
	b := &SpecimenBuilder{ResourceType: "Specimen", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

func (b *SpecimenBuilder) set(k string, v interface{}) *SpecimenBuilder {
	setPath(b.Data, k, v)
	return b
}

func (b *SpecimenBuilder) push(k string, v interface{}) *SpecimenBuilder {
	pushPath(b.Data, k, v)
	return b
}

// SetId sets the logical id.
func (b *SpecimenBuilder) SetId(id string) *SpecimenBuilder { return b.set("id", id) }

// SetStatus sets status (available | unavailable | unsatisfactory | entered-in-error).
func (b *SpecimenBuilder) SetStatus(status string) *SpecimenBuilder { return b.set("status", status) }

// SetType sets type.
func (b *SpecimenBuilder) SetType(typ *datatype.CodeableConcept) *SpecimenBuilder {
	return b.set("type", typ.ToArray())
}

// SetSubject sets subject.
func (b *SpecimenBuilder) SetSubject(subject *datatype.Reference) *SpecimenBuilder {
	return b.set("subject", subject.ToArray())
}

// SetReceivedTime sets receivedTime.
func (b *SpecimenBuilder) SetReceivedTime(dateTime string) *SpecimenBuilder {
	return b.set("receivedTime", dateTime)
}

// SetCollectedDateTime sets collection.collectedDateTime.
func (b *SpecimenBuilder) SetCollectedDateTime(dateTime string) *SpecimenBuilder {
	return b.set("collection/collectedDateTime", dateTime)
}

// SetCollector sets collection.collector.
func (b *SpecimenBuilder) SetCollector(collector *datatype.Reference) *SpecimenBuilder {
	return b.set("collection/collector", collector.ToArray())
}

// SetFastingStatusCodeableConcept sets collection.fastingStatusCodeableConcept.
func (b *SpecimenBuilder) SetFastingStatusCodeableConcept(status *datatype.CodeableConcept) *SpecimenBuilder {
	return b.set("collection/fastingStatusCodeableConcept", status.ToArray())
}

// SetMethod sets collection.method.
func (b *SpecimenBuilder) SetMethod(method *datatype.CodeableConcept) *SpecimenBuilder {
	return b.set("collection/method", method.ToArray())
}

// SetQuantity sets collection.quantity.
func (b *SpecimenBuilder) SetQuantity(quantity *datatype.Quantity) *SpecimenBuilder {
	return b.set("collection/quantity", quantity.ToArray())
}

// SetBodySite sets collection.bodySite.
func (b *SpecimenBuilder) SetBodySite(bodySite *datatype.CodeableConcept) *SpecimenBuilder {
	return b.set("collection/bodySite", bodySite.ToArray())
}

// AddIdentifier appends an identifier.
func (b *SpecimenBuilder) AddIdentifier(identifier *datatype.Identifier) *SpecimenBuilder {
	return b.push("identifier", identifier.ToArray())
}

// AddRequest appends a request reference.
func (b *SpecimenBuilder) AddRequest(request *datatype.Reference) *SpecimenBuilder {
	return b.push("request", request.ToArray())
}

// AddCondition appends a condition as free text.
func (b *SpecimenBuilder) AddCondition(text string) *SpecimenBuilder {
	return b.push("condition", map[string]interface{}{"text": text})
}

// AddProcessing appends a processing entry with timeDateTime.
func (b *SpecimenBuilder) AddProcessing(dateTime string) *SpecimenBuilder {
	return b.push("processing", map[string]interface{}{"timeDateTime": dateTime})
}

// AddExtension appends an extension with a string value.
func (b *SpecimenBuilder) AddExtension(url, value string) *SpecimenBuilder {
	return b.push("extension", map[string]interface{}{"url": url, "valueString": value})
}

// AddTransportedTime appends the Kemkes TransportedTime extension.
func (b *SpecimenBuilder) AddTransportedTime(dateTime string) *SpecimenBuilder {
	return b.push("extension", map[string]interface{}{
		"url":           "https://fhir.kemkes.go.id/r4/StructureDefinition/TransportedTime",
		"valueDateTime": dateTime,
	})
}

// AddTransportedPerson appends the Kemkes TransportedPerson extension as a valueContactDetail.
func (b *SpecimenBuilder) AddTransportedPerson(name string, telecom []map[string]interface{}) *SpecimenBuilder {
	if telecom == nil {
		telecom = []map[string]interface{}{}
	}
	return b.push("extension", map[string]interface{}{
		"url": "https://fhir.kemkes.go.id/r4/StructureDefinition/TransportedPerson",
		"valueContactDetail": map[string]interface{}{
			"name":    name,
			"telecom": telecom,
		},
	})
}

// AddReceivedPerson appends the Kemkes ReceivedPerson extension as a valueReference.
func (b *SpecimenBuilder) AddReceivedPerson(person *datatype.Reference) *SpecimenBuilder {
	return b.push("extension", map[string]interface{}{
		"url":            "https://fhir.kemkes.go.id/r4/StructureDefinition/ReceivedPerson",
		"valueReference": person.ToArray(),
	})
}

// Build returns the payload.
func (b *SpecimenBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *SpecimenBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
