package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type ExplanationOfBenefitBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewExplanationOfBenefitBuilder() *ExplanationOfBenefitBuilder {
	b := &ExplanationOfBenefitBuilder{ResourceType: "ExplanationOfBenefit", Data: make(map[string]interface{})}
	return b
}

func (b *ExplanationOfBenefitBuilder) SetId(id string) *ExplanationOfBenefitBuilder { b.Data["id"] = id; return b }
func (b *ExplanationOfBenefitBuilder) AddIdentifier(id *datatype.Identifier) *ExplanationOfBenefitBuilder {
	if _, ok := b.Data["identifier"]; !ok { b.Data["identifier"] = make([]interface{}, 0) }
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), id.ToArray())
	return b
}
func (b *ExplanationOfBenefitBuilder) SetStatus(status string) *ExplanationOfBenefitBuilder { b.Data["status"] = status; return b }
func (b *ExplanationOfBenefitBuilder) SetType(code *datatype.CodeableConcept) *ExplanationOfBenefitBuilder { b.Data["type"] = code.ToArray(); return b }
func (b *ExplanationOfBenefitBuilder) SetUse(use string) *ExplanationOfBenefitBuilder { b.Data["use"] = use; return b }
func (b *ExplanationOfBenefitBuilder) SetPatient(patientRef string) *ExplanationOfBenefitBuilder { b.Data["patient"] = map[string]interface{}{"reference": patientRef}; return b }
func (b *ExplanationOfBenefitBuilder) SetBillablePeriod(start string, end string) *ExplanationOfBenefitBuilder {
	p := map[string]interface{}{"start": start}
	if end != "" { p["end"] = end }
	b.Data["billablePeriod"] = p
	return b
}
func (b *ExplanationOfBenefitBuilder) SetCreated(created string) *ExplanationOfBenefitBuilder { b.Data["created"] = created; return b }
func (b *ExplanationOfBenefitBuilder) SetInsurer(insurerRef string) *ExplanationOfBenefitBuilder { b.Data["insurer"] = map[string]interface{}{"reference": insurerRef}; return b }
func (b *ExplanationOfBenefitBuilder) SetProvider(providerRef string) *ExplanationOfBenefitBuilder { b.Data["provider"] = map[string]interface{}{"reference": providerRef}; return b }
func (b *ExplanationOfBenefitBuilder) SetOutcome(outcome string) *ExplanationOfBenefitBuilder { b.Data["outcome"] = outcome; return b }
func (b *ExplanationOfBenefitBuilder) AddCareTeam(memberRef string, qualificationCode string) *ExplanationOfBenefitBuilder {
	if _, ok := b.Data["careTeam"]; !ok { b.Data["careTeam"] = make([]interface{}, 0) }
	b.Data["careTeam"] = append(b.Data["careTeam"].([]interface{}), map[string]interface{}{
		"sequence":    len(b.Data["careTeam"].([]interface{})) + 1,
		"provider":   map[string]interface{}{"reference": memberRef},
		"qualification": map[string]interface{}{"code": map[string]interface{}{"coding": []interface{}{map[string]interface{}{"code": qualificationCode}}}},
	})
	return b
}
func (b *ExplanationOfBenefitBuilder) AddDiagnosis(code string, type_ string) *ExplanationOfBenefitBuilder {
	if _, ok := b.Data["diagnosis"]; !ok { b.Data["diagnosis"] = make([]interface{}, 0) }
	b.Data["diagnosis"] = append(b.Data["diagnosis"].([]interface{}), map[string]interface{}{
		"sequence":  len(b.Data["diagnosis"].([]interface{})) + 1,
		"diagnosisCodeableConcept": map[string]interface{}{"coding": []interface{}{map[string]interface{}{"code": code}}},
		"type": []interface{}{map[string]interface{}{"coding": []interface{}{map[string]interface{}{"code": type_}}}},
	})
	return b
}
func (b *ExplanationOfBenefitBuilder) AddProcedure(code string, date string) *ExplanationOfBenefitBuilder {
	if _, ok := b.Data["procedure"]; !ok { b.Data["procedure"] = make([]interface{}, 0) }
	b.Data["procedure"] = append(b.Data["procedure"].([]interface{}), map[string]interface{}{
		"sequence":      len(b.Data["procedure"].([]interface{})) + 1,
		"procedureCodeableConcept": map[string]interface{}{"coding": []interface{}{map[string]interface{}{"code": code}}},
		"date": date,
	})
	return b
}
func (b *ExplanationOfBenefitBuilder) AddInsurance(coverageRef string) *ExplanationOfBenefitBuilder {
	if _, ok := b.Data["insurance"]; !ok { b.Data["insurance"] = make([]interface{}, 0) }
	b.Data["insurance"] = append(b.Data["insurance"].([]interface{}), map[string]interface{}{
		"focal":     true,
		"coverage": map[string]interface{}{"reference": coverageRef},
	})
	return b
}
func (b *ExplanationOfBenefitBuilder) AddItem(sequence int, serviceCode string, unitPrice float64, currency string) *ExplanationOfBenefitBuilder {
	if _, ok := b.Data["item"]; !ok { b.Data["item"] = make([]interface{}, 0) }
	item := map[string]interface{}{
		"sequence": sequence,
		"service": map[string]interface{}{"coding": []interface{}{map[string]interface{}{"code": serviceCode}}},
		"unitPrice": []interface{}{map[string]interface{}{"value": unitPrice, "currency": currency}},
	}
	b.Data["item"] = append(b.Data["item"].([]interface{}), item)
	return b
}
func (b *ExplanationOfBenefitBuilder) AddItemDiagnosisLink(itemSeq int, diagnosisSeq int) *ExplanationOfBenefitBuilder {
	if items, ok := b.Data["item"].([]interface{}); ok {
		for _, it := range items {
			if m, ok := it.(map[string]interface{}); ok && int(m["sequence"].(int)) == itemSeq {
				if _, ok := m["diagnosisLinkType"]; !ok {
					m["diagnosisLinkType"] = []interface{}{map[string]interface{}{"coding": []interface{}{map[string]interface{}{"code": "resolved-from-claimitem"}}}}
				}
				m["diagnosisSequence"] = []interface{}{diagnosisSeq}
				break
			}
		}
	}
	return b
}
func (b *ExplanationOfBenefitBuilder) SetTotal(amount float64, currency string) *ExplanationOfBenefitBuilder {
	b.Data["total"] = []interface{}{
		map[string]interface{}{
			"category": map[string]interface{}{"coding": []interface{}{map[string]interface{}{"code": "submitted"}}},
			"amount": map[string]interface{}{"value": amount, "currency": currency},
		},
	}
	return b
}
func (b *ExplanationOfBenefitBuilder) Build() map[string]interface{} { return b.Data }
