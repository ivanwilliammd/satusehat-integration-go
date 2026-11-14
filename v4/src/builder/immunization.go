package builder

import (
	"encoding/json"
	"fmt"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

// ImmunizationBuilder builds a FHIR R4 Immunization resource.
//
// Method-for-method counterpart of PayloadBuilderImmunization in the PHP SDK. Two behaviours are
// copied deliberately, including their quirks:
//
//   - status is validated: only completed, entered-in-error and not-done are accepted;
//   - setVaccineCodeFromCode only accepts the single code the PHP builder allowlists (93001282).
//
// Deviation from PHP: an invalid value does not panic. The first error is stored and returned by
// Err(), the element is not written, and the rest of the builder keeps working. PHP throws
// immediately; Go callers should check Err() after building.
//
// @link https://www.hl7.org/fhir/immunization.html
type ImmunizationBuilder struct {
	ResourceType string
	Data         map[string]interface{}
	err          error
}

// NewImmunizationBuilder creates an empty Immunization with its resourceType set.
func NewImmunizationBuilder() *ImmunizationBuilder {
	b := &ImmunizationBuilder{ResourceType: "Immunization", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

// Err returns the first validation error recorded by a setter, if any.
func (b *ImmunizationBuilder) Err() error { return b.err }

func (b *ImmunizationBuilder) set(k string, v interface{}) *ImmunizationBuilder {
	b.Data[k] = v
	return b
}

func (b *ImmunizationBuilder) push(k string, v interface{}) *ImmunizationBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

func (b *ImmunizationBuilder) fail(format string, args ...interface{}) *ImmunizationBuilder {
	if b.err == nil {
		b.err = fmt.Errorf(format, args...)
	}
	return b
}

// SetMetaProfile appends meta.profile.
func (b *ImmunizationBuilder) SetMetaProfile(profile string) *ImmunizationBuilder {
	// PHP's push('meta/profile', …) writes meta.profile; a literal "meta/profile" key would be a
	// valid-looking payload that no FHIR server reads.
	meta, _ := b.Data["meta"].(map[string]interface{})
	if meta == nil {
		meta = map[string]interface{}{}
		b.Data["meta"] = meta
	}
	list, _ := meta["profile"].([]interface{})
	meta["profile"] = append(list, profile)
	return b
}

// SetId sets the logical id.
func (b *ImmunizationBuilder) SetId(id string) *ImmunizationBuilder { return b.set("id", id) }

// AddIdentifier appends an identifier.
func (b *ImmunizationBuilder) AddIdentifier(identifier *datatype.Identifier) *ImmunizationBuilder {
	return b.push("identifier", identifier.ToArray())
}

// SetStatus sets status. Only completed, entered-in-error and not-done are accepted.
func (b *ImmunizationBuilder) SetStatus(status string) *ImmunizationBuilder {
	switch status {
	case "completed", "entered-in-error", "not-done":
		return b.set("status", status)
	}
	return b.fail("Invalid Immunization.status: %s", status)
}

// SetVaccineCode sets vaccineCode.
func (b *ImmunizationBuilder) SetVaccineCode(vaccine *datatype.CodeableConcept) *ImmunizationBuilder {
	return b.set("vaccineCode", vaccine.ToArray())
}

// SetVaccineCodeFromCode sets vaccineCode from a code, validating it against the PHP allowlist.
func (b *ImmunizationBuilder) SetVaccineCodeFromCode(code, display, system string) *ImmunizationBuilder {
	if code != "93001282" {
		return b.fail("Invalid vaccine code: %s", code)
	}
	if system == "" {
		system = "http://snomed.info/sct"
	}
	return b.SetVaccineCode(&datatype.CodeableConcept{Coding: []datatype.Coding{{
		System: system, Code: code, Display: display,
	}}})
}

// SetPatient sets patient, optionally with a display value.
func (b *ImmunizationBuilder) SetPatient(patient *datatype.Reference, display string) *ImmunizationBuilder {
	m := patient.ToArray()
	if display != "" {
		m["display"] = display
	}
	return b.set("patient", m)
}

// SetOccurrenceDateTime sets occurrenceDateTime.
func (b *ImmunizationBuilder) SetOccurrenceDateTime(dateTime string) *ImmunizationBuilder {
	return b.set("occurrenceDateTime", dateTime)
}

// AddPerformer appends a performer. function is optional.
func (b *ImmunizationBuilder) AddPerformer(actor *datatype.Reference, function *datatype.CodeableConcept) *ImmunizationBuilder {
	performer := map[string]interface{}{"actor": actor.ToArray()}
	if function != nil {
		performer["function"] = function.ToArray()
	}
	return b.push("performer", performer)
}

// SetDoseQuantity sets doseQuantity.
func (b *ImmunizationBuilder) SetDoseQuantity(value float64, unit, system, code string) *ImmunizationBuilder {
	q := &datatype.Quantity{Value: &value}
	if unit != "" {
		q.Unit = &unit
	}
	if system != "" {
		q.System = &system
	}
	if code != "" {
		q.Code = &code
	}
	return b.set("doseQuantity", q.ToArray())
}

// SetLocation sets location, optionally with a display value.
func (b *ImmunizationBuilder) SetLocation(location *datatype.Reference, display string) *ImmunizationBuilder {
	m := location.ToArray()
	if display != "" {
		m["display"] = display
	}
	return b.set("location", m)
}

// SetLotNumber sets lotNumber.
func (b *ImmunizationBuilder) SetLotNumber(lotNumber string) *ImmunizationBuilder {
	return b.set("lotNumber", lotNumber)
}

// SetRecorded sets recorded.
func (b *ImmunizationBuilder) SetRecorded(dateTime string) *ImmunizationBuilder {
	return b.set("recorded", dateTime)
}

// SetPrimarySource sets primarySource.
func (b *ImmunizationBuilder) SetPrimarySource(primarySource bool) *ImmunizationBuilder {
	return b.set("primarySource", primarySource)
}

// AddProtocolApplied appends a protocolApplied entry. FHIR R4 puts the series as a CodeableConcept
// and keeps seriesDoses[x] separate, so a series is written as series and seriesDosesPositiveInt is
// only written when the caller passes it.
func (b *ImmunizationBuilder) AddProtocolApplied(doseNumberPositiveInt int, series *datatype.CodeableConcept, seriesDosesPositiveInt *int) *ImmunizationBuilder {
	row := map[string]interface{}{"doseNumberPositiveInt": doseNumberPositiveInt}
	if series != nil {
		row["series"] = series.ToArray()
	}
	if seriesDosesPositiveInt != nil {
		row["seriesDosesPositiveInt"] = *seriesDosesPositiveInt
	}
	return b.push("protocolApplied", row)
}

// AddReasonCode appends a reasonCode.
func (b *ImmunizationBuilder) AddReasonCode(reason *datatype.CodeableConcept) *ImmunizationBuilder {
	return b.push("reasonCode", reason.ToArray())
}

// SetRoute sets route.
func (b *ImmunizationBuilder) SetRoute(route *datatype.CodeableConcept) *ImmunizationBuilder {
	return b.set("route", route.ToArray())
}

// SetRouteFromCode sets route from a code, defaulting to the HL7 route system.
func (b *ImmunizationBuilder) SetRouteFromCode(code, display, system string) *ImmunizationBuilder {
	if system == "" {
		system = "http://terminology.hl7.org/CodeSystem/v3-RouteOfAdministration"
	}
	return b.SetRoute(&datatype.CodeableConcept{Coding: []datatype.Coding{{
		System: system, Code: code, Display: display,
	}}})
}

// AddExtension appends an extension with a string value.
func (b *ImmunizationBuilder) AddExtension(url, value string) *ImmunizationBuilder {
	return b.push("extension", map[string]interface{}{"url": url, "valueString": value})
}

// Build returns the payload.
func (b *ImmunizationBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *ImmunizationBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
