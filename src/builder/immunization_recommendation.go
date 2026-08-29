package builder

import (
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

// ImmunizationRecommendationBuilder builds FHIR ImmunizationRecommendation payload
type ImmunizationRecommendationBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewImmunizationRecommendationBuilder() *ImmunizationRecommendationBuilder {
	b := &ImmunizationRecommendationBuilder{ResourceType: "ImmunizationRecommendation", Data: make(map[string]interface{})}
	return b
}

func (b *ImmunizationRecommendationBuilder) setId(id string) *ImmunizationRecommendationBuilder {
	b.Data["id"] = id
	return b
}

func (b *ImmunizationRecommendationBuilder) addIdentifier(identifier *datatype.Identifier) *ImmunizationRecommendationBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), identifier.ToArray())
	return b
}

func (b *ImmunizationRecommendationBuilder) setPatient(patient *datatype.Reference) *ImmunizationRecommendationBuilder {
	b.Data["patient"] = patient.ToArray()
	return b
}

func (b *ImmunizationRecommendationBuilder) addRecommendation(
	vaccineCode *datatype.CodeableConcept,
	targetDisease *datatype.CodeableConcept,
	doseNumber int,
	seriesDoses int,
	forecastStatus *datatype.CodeableConcept,
	dateCriterion []map[string]interface{},
	protocol *map[string]interface{},
	supportingImmunization []datatype.Reference,
	supportingPatientInfo []datatype.Reference,
) *ImmunizationRecommendationBuilder {
	rec := map[string]interface{}{}
	if vaccineCode != nil {
		rec["vaccineCode"] = vaccineCode.ToArray()
	}
	if targetDisease != nil {
		rec["targetDisease"] = targetDisease.ToArray()
	}
	if doseNumber != 0 {
		rec["doseNumber"] = doseNumber
	}
	if seriesDoses != 0 {
		rec["seriesDoses"] = seriesDoses
	}
	if forecastStatus != nil {
		rec["forecastStatus"] = forecastStatus.ToArray()
	}
	if len(dateCriterion) > 0 {
		rec["dateCriterion"] = dateCriterion
	}
	if protocol != nil {
		rec["protocol"] = *protocol
	}
	if len(supportingImmunization) > 0 {
		supp := make([]interface{}, len(supportingImmunization))
		for i, s := range supportingImmunization {
			supp[i] = s.ToArray()
		}
		rec["supportingImmunization"] = supp
	}
	if len(supportingPatientInfo) > 0 {
		info := make([]interface{}, len(supportingPatientInfo))
		for i, p := range supportingPatientInfo {
			info[i] = p.ToArray()
		}
		rec["supportingPatientInformation"] = info
	}
	if _, ok := b.Data["recommendation"]; !ok {
		b.Data["recommendation"] = make([]interface{}, 0)
	}
	b.Data["recommendation"] = append(b.Data["recommendation"].([]interface{}), rec)
	return b
}

func (b *ImmunizationRecommendationBuilder) addExtension(url string, value interface{}, valueType string) *ImmunizationRecommendationBuilder {
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

func (b *ImmunizationRecommendationBuilder) Build() map[string]interface{} {
	return b.Data
}
