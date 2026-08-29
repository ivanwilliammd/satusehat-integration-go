package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type ClinicalImpressionBuilder struct {
	Data map[string]interface{}
}

func NewClinicalImpressionBuilder() *ClinicalImpressionBuilder {
	return &ClinicalImpressionBuilder{Data: make(map[string]interface{})}
}

func (b *ClinicalImpressionBuilder) SetId(id string) *ClinicalImpressionBuilder {
	b.Data["id"] = id
	return b
}

func (b *ClinicalImpressionBuilder) AddIdentifier(ident *datatype.Identifier) *ClinicalImpressionBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), ident.ToArray())
	return b
}

func (b *ClinicalImpressionBuilder) SetStatus(status string) *ClinicalImpressionBuilder {
	b.Data["status"] = status
	return b
}

func (b *ClinicalImpressionBuilder) SetCode(cc *datatype.CodeableConcept) *ClinicalImpressionBuilder {
	b.Data["code"] = cc.ToArray()
	return b
}

func (b *ClinicalImpressionBuilder) SetSubject(ref *datatype.Reference) *ClinicalImpressionBuilder {
	b.Data["subject"] = ref.ToArray()
	return b
}

func (b *ClinicalImpressionBuilder) SetEncounter(ref *datatype.Reference) *ClinicalImpressionBuilder {
	b.Data["encounter"] = ref.ToArray()
	return b
}

func (b *ClinicalImpressionBuilder) SetEffectiveDateTime(dt string) *ClinicalImpressionBuilder {
	b.Data["effectiveDateTime"] = dt
	return b
}

func (b *ClinicalImpressionBuilder) SetEffectivePeriod(period *datatype.Period) *ClinicalImpressionBuilder {
	b.Data["effectivePeriod"] = period.ToArray()
	return b
}

func (b *ClinicalImpressionBuilder) SetDate(date string) *ClinicalImpressionBuilder {
	b.Data["date"] = date
	return b
}

func (b *ClinicalImpressionBuilder) SetAssessor(ref *datatype.Reference) *ClinicalImpressionBuilder {
	b.Data["assessor"] = ref.ToArray()
	return b
}

func (b *ClinicalImpressionBuilder) Build() map[string]interface{} {
	dt := make(map[string]interface{})
	for k, v := range b.Data {
		if v == nil || v == "" {
			continue
		}
		dt[k] = v
	}
	dt["resourceType"] = "ClinicalImpression"
	return dt
}
