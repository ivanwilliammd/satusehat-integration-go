package builder

import (
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

// ImmunizationBuilder builds FHIR Immunization payload
type ImmunizationBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewImmunizationBuilder() *ImmunizationBuilder {
	b := &ImmunizationBuilder{ResourceType: "Immunization", Data: make(map[string]interface{})}
	return b
}

func (b *ImmunizationBuilder) setId(id string) *ImmunizationBuilder {
	b.Data["id"] = id
	return b
}

func (b *ImmunizationBuilder) addIdentifier(identifier *datatype.Identifier) *ImmunizationBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), identifier.ToArray())
	return b
}

func (b *ImmunizationBuilder) setStatus(status string) *ImmunizationBuilder {
	b.Data["status"] = status
	return b
}

func (b *ImmunizationBuilder) setVaccineCode(vaccineCode *datatype.CodeableConcept) *ImmunizationBuilder {
	b.Data["vaccineCode"] = vaccineCode.ToArray()
	return b
}

func (b *ImmunizationBuilder) setPatient(patient *datatype.Reference) *ImmunizationBuilder {
	b.Data["patient"] = patient.ToArray()
	return b
}

func (b *ImmunizationBuilder) setEncounter(encounter *datatype.Reference) *ImmunizationBuilder {
	b.Data["encounter"] = encounter.ToArray()
	return b
}

func (b *ImmunizationBuilder) setOccurrenceDateTime(dateTime string) *ImmunizationBuilder {
	b.Data["occurrenceDateTime"] = dateTime
	return b
}

func (b *ImmunizationBuilder) setPrimarySource(primarySource bool) *ImmunizationBuilder {
	b.Data["primarySource"] = primarySource
	return b
}

func (b *ImmunizationBuilder) setReportOrigin(origin *datatype.CodeableConcept) *ImmunizationBuilder {
	b.Data["reportOrigin"] = origin.ToArray()
	return b
}

func (b *ImmunizationBuilder) setLocation(location *datatype.Reference) *ImmunizationBuilder {
	b.Data["location"] = location.ToArray()
	return b
}

func (b *ImmunizationBuilder) setManufacturer(manufacturer *datatype.Reference) *ImmunizationBuilder {
	b.Data["manufacturer"] = manufacturer.ToArray()
	return b
}

func (b *ImmunizationBuilder) setLotNumber(lotNumber string) *ImmunizationBuilder {
	b.Data["lotNumber"] = lotNumber
	return b
}

func (b *ImmunizationBuilder) setExpirationDate(expirationDate string) *ImmunizationBuilder {
	b.Data["expirationDate"] = expirationDate
	return b
}

func (b *ImmunizationBuilder) setSite(site *datatype.CodeableConcept) *ImmunizationBuilder {
	b.Data["site"] = site.ToArray()
	return b
}

func (b *ImmunizationBuilder) setRoute(route *datatype.CodeableConcept) *ImmunizationBuilder {
	b.Data["route"] = route.ToArray()
	return b
}

func (b *ImmunizationBuilder) setDoseQuantity(quantity *datatype.Quantity) *ImmunizationBuilder {
	b.Data["doseQuantity"] = quantity.ToArray()
	return b
}

func (b *ImmunizationBuilder) setPerformer(performer *datatype.Reference) *ImmunizationBuilder {
	b.Data["performer"] = performer.ToArray()
	return b
}

func (b *ImmunizationBuilder) addNote(note *datatype.Annotation) *ImmunizationBuilder {
	if _, ok := b.Data["note"]; !ok {
		b.Data["note"] = make([]interface{}, 0)
	}
	b.Data["note"] = append(b.Data["note"].([]interface{}), note.ToArray())
	return b
}

func (b *ImmunizationBuilder) setReasonCode(reasonCode *datatype.CodeableConcept) *ImmunizationBuilder {
	b.Data["reasonCode"] = reasonCode.ToArray()
	return b
}

func (b *ImmunizationBuilder) setReasonReference(reasonRef *datatype.Reference) *ImmunizationBuilder {
	b.Data["reasonReference"] = reasonRef.ToArray()
	return b
}

func (b *ImmunizationBuilder) setIsSubpotent(isSubpotent bool) *ImmunizationBuilder {
	b.Data["isSubpotent"] = isSubpotent
	return b
}

func (b *ImmunizationBuilder) addReaction(
	date string,
	detail *datatype.Reference,
	severity string,
) *ImmunizationBuilder {
	reaction := map[string]interface{}{}
	if date != "" {
		reaction["date"] = date
	}
	if detail != nil {
		reaction["detail"] = detail.ToArray()
	}
	if severity != "" {
		reaction["severity"] = severity
	}
	if _, ok := b.Data["reaction"]; !ok {
		b.Data["reaction"] = make([]interface{}, 0)
	}
	b.Data["reaction"] = append(b.Data["reaction"].([]interface{}), reaction)
	return b
}

func (b *ImmunizationBuilder) addProtocolDose(sequence int, description string, targetDisease []datatype.CodeableConcept) *ImmunizationBuilder {
	protocol := map[string]interface{}{}
	if sequence != 0 {
		protocol["sequence"] = sequence
	}
	if description != "" {
		protocol["description"] = description
	}
	if len(targetDisease) > 0 {
		diseases := make([]interface{}, len(targetDisease))
		for i, d := range targetDisease {
			diseases[i] = d.ToArray()
		}
		protocol["targetDisease"] = diseases
	}
	if _, ok := b.Data["protocolApplied"]; !ok {
		b.Data["protocolApplied"] = make([]interface{}, 0)
	}
	b.Data["protocolApplied"] = append(b.Data["protocolApplied"].([]interface{}), protocol)
	return b
}

func (b *ImmunizationBuilder) addExtension(url string, value interface{}, valueType string) *ImmunizationBuilder {
	ext := map[string]interface{}{"url": url}
	if valueType != "" {
		capitalized := strings.ToUpper(valueType[:1]) + valueType[1:]
		ext["value"+capitalized] = value
	} else {
		ext["valueString"] = value
	}
	if _, ok := b.Data["extension"]; !ok {
		b.Data["extension"] = make([]interface{}, 0)
	}
	b.Data["extension"] = append(b.Data["extension"].([]interface{}), ext)
	return b
}

func (b *ImmunizationBuilder) Build() map[string]interface{} {
	return b.Data
}
