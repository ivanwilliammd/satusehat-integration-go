package builder

import (
	"encoding/json"
	"strings"
	"time"
)

// MedicationDispenseBuilder builds a FHIR R4 MedicationDispense resource.
//
// Method-for-method counterpart of PayloadBuilderMedicationDispense in the PHP SDK. That builder
// takes plain ids and builds the Reference itself, so the same applies here: passing "1000000001"
// to SetSubject produces Patient/1000000001, while passing a full reference keeps it unchanged.
//
// @link https://www.hl7.org/fhir/medicationdispense.html
type MedicationDispenseBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewMedicationDispenseBuilder creates an empty MedicationDispense with its resourceType set.
func NewMedicationDispenseBuilder() *MedicationDispenseBuilder {
	b := &MedicationDispenseBuilder{ResourceType: "MedicationDispense", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

func (b *MedicationDispenseBuilder) set(k string, v interface{}) *MedicationDispenseBuilder {
	b.Data[k] = v
	return b
}

func (b *MedicationDispenseBuilder) push(k string, v interface{}) *MedicationDispenseBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

// ref membuat "Tipe/id" bila yang diberikan hanya id, dan membiarkan referensi penuh apa adanya.
func ref(typ, idOrReference string) string {
	if strings.Contains(idOrReference, "/") {
		return idOrReference
	}
	return typ + "/" + idOrReference
}

// isoTime meniru date('c', strtotime(...)) pada builder PHP. Bila masukan tidak dapat diurai, nilai
// asli dipertahankan apa adanya alih-alih dikarang.
func isoTime(v string) string {
	for _, layout := range []string{time.RFC3339, "2006-01-02 15:04:05", "2006-01-02"} {
		if t, err := time.Parse(layout, v); err == nil {
			// Keep the offset carried by the input. PHP's date('c') instead formats in the PHP
			// default timezone, so a server whose timezone differs from the input offset rewrites
			// the timestamp; the parity test pins the case where both agree.
			if layout == "2006-01-02" {
				return t.Format("2006-01-02T00:00:00-07:00")
			}
			return t.Format("2006-01-02T15:04:05-07:00")
		}
	}
	return v
}

// SetId sets the logical id.
func (b *MedicationDispenseBuilder) SetId(id string) *MedicationDispenseBuilder {
	return b.set("id", id)
}

// AddIdentifier appends an identifier from its system and value.
func (b *MedicationDispenseBuilder) AddIdentifier(system, value string) *MedicationDispenseBuilder {
	return b.push("identifier", map[string]interface{}{"system": system, "value": value})
}

// SetStatus sets status (preparation | in-progress | cancelled | on-hold | completed | entered-in-error | stopped | declined | unknown).
func (b *MedicationDispenseBuilder) SetStatus(status string) *MedicationDispenseBuilder {
	return b.set("status", status)
}

// SetMedicationReference sets medicationReference. display is optional.
func (b *MedicationDispenseBuilder) SetMedicationReference(reference, display string) *MedicationDispenseBuilder {
	m := map[string]interface{}{"reference": reference}
	if display != "" {
		m["display"] = display
	}
	return b.set("medicationReference", m)
}

// SetSubject sets subject; a bare id becomes Patient/<id>.
func (b *MedicationDispenseBuilder) SetSubject(subjectID, display string) *MedicationDispenseBuilder {
	m := map[string]interface{}{"reference": ref("Patient", subjectID)}
	if display != "" {
		m["display"] = display
	}
	return b.set("subject", m)
}

// SetContext sets context.reference; a bare id becomes Encounter/<id>.
func (b *MedicationDispenseBuilder) SetContext(contextID string) *MedicationDispenseBuilder {
	return b.set("context", map[string]interface{}{"reference": ref("Encounter", contextID)})
}

// AddPerformer appends a performer. display is optional.
func (b *MedicationDispenseBuilder) AddPerformer(actorID, display string) *MedicationDispenseBuilder {
	actor := map[string]interface{}{"reference": ref("Practitioner", actorID)}
	if display != "" {
		actor["display"] = display
	}
	return b.push("performer", map[string]interface{}{"actor": actor})
}

// SetLocation sets location. display is optional.
func (b *MedicationDispenseBuilder) SetLocation(locationID, display string) *MedicationDispenseBuilder {
	m := map[string]interface{}{"reference": ref("Location", locationID)}
	if display != "" {
		m["display"] = display
	}
	return b.set("location", m)
}

// AddAuthorizingPrescription appends an authorizingPrescription.
func (b *MedicationDispenseBuilder) AddAuthorizingPrescription(requestID string) *MedicationDispenseBuilder {
	return b.push("authorizingPrescription", map[string]interface{}{
		"reference": ref("MedicationRequest", requestID),
	})
}

// SetCategory sets category.coding. system defaults to the HL7 dispense category system.
func (b *MedicationDispenseBuilder) SetCategory(code, display, system string) *MedicationDispenseBuilder {
	if system == "" {
		system = "http://terminology.hl7.org/CodeSystem/medicationdispense-category"
	}
	coding := map[string]interface{}{"system": system, "code": code}
	if display != "" {
		coding["display"] = display
	}
	return b.set("category", map[string]interface{}{"coding": []interface{}{coding}})
}

// SetQuantity sets quantity.
func (b *MedicationDispenseBuilder) SetQuantity(value float64, unit, system, code string) *MedicationDispenseBuilder {
	q := map[string]interface{}{"value": value, "unit": unit}
	if system != "" {
		q["system"] = system
	}
	if code != "" {
		q["code"] = code
	}
	return b.set("quantity", q)
}

// SetDaysSupply sets daysSupply.
func (b *MedicationDispenseBuilder) SetDaysSupply(value float64, unit, system, code string) *MedicationDispenseBuilder {
	ds := map[string]interface{}{"value": value, "unit": unit}
	if system != "" {
		ds["system"] = system
	}
	if code != "" {
		ds["code"] = code
	}
	return b.set("daysSupply", ds)
}

// SetWhenPrepared sets whenPrepared in ISO 8601, as the PHP builder does.
func (b *MedicationDispenseBuilder) SetWhenPrepared(dateTime string) *MedicationDispenseBuilder {
	return b.set("whenPrepared", isoTime(dateTime))
}

// SetWhenHandedOver sets whenHandedOver in ISO 8601, as the PHP builder does.
func (b *MedicationDispenseBuilder) SetWhenHandedOver(dateTime string) *MedicationDispenseBuilder {
	return b.set("whenHandedOver", isoTime(dateTime))
}

// AddDosageInstruction appends a dosageInstruction built from flat arguments, matching the PHP
// signature rather than taking a Dosage object.
func (b *MedicationDispenseBuilder) AddDosageInstruction(sequence int, text string, doseValue float64, doseUnit, routeCode, routeDisplay, timingCode, timingDisplay, additionalCode, additionalText string) *MedicationDispenseBuilder {
	return b.push("dosageInstruction", map[string]interface{}{
		"sequence": sequence,
		"text":     text,
		"additionalInstruction": []interface{}{map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"code": additionalCode, "display": additionalText}},
		}},
		"timing": map[string]interface{}{
			"code": map[string]interface{}{
				"coding": []interface{}{map[string]interface{}{"code": timingCode, "display": timingDisplay}},
			},
		},
		"route": map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"code": routeCode, "display": routeDisplay}},
		},
		"doseAndRate": []interface{}{map[string]interface{}{
			"doseQuantity": map[string]interface{}{"value": doseValue, "unit": doseUnit},
		}},
	})
}

// SetSubstitution sets substitution. typeCode is optional; typeDisplay is written only with typeCode.
func (b *MedicationDispenseBuilder) SetSubstitution(wasSubstituted bool, typeCode, typeDisplay string) *MedicationDispenseBuilder {
	substitution := map[string]interface{}{"wasSubstituted": wasSubstituted}
	if typeCode != "" {
		coding := map[string]interface{}{"code": typeCode}
		if typeDisplay != "" {
			coding["display"] = typeDisplay
		}
		substitution["type"] = map[string]interface{}{"coding": []interface{}{coding}}
	}
	return b.set("substitution", substitution)
}

// AddContained appends a contained resource.
func (b *MedicationDispenseBuilder) AddContained(resource map[string]interface{}) *MedicationDispenseBuilder {
	return b.push("contained", resource)
}

// Build returns the payload.
func (b *MedicationDispenseBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *MedicationDispenseBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
