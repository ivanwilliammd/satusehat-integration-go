package builder

import (
	"encoding/json"
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type ClaimBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewClaimBuilder() *ClaimBuilder {
	b := &ClaimBuilder{ResourceType: "Claim", Data: make(map[string]interface{})}
	return b
}

func (b *ClaimBuilder) setMetaProfile(profile string) *ClaimBuilder {
	if _, ok := b.Data["meta"]; !ok {
		b.Data["meta"] = make(map[string]interface{})
	}
	b.Data["meta"].(map[string]interface{})["profile"] = profile
	return b
}

func (b *ClaimBuilder) setId(id string) *ClaimBuilder {
	b.Data["id"] = id
	return b
}

func (b *ClaimBuilder) addIdentifier(identifier *datatype.Identifier) *ClaimBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), identifier.ToArray())
	return b
}

func (b *ClaimBuilder) setStatus(status string) *ClaimBuilder {
	b.Data["status"] = status
	return b
}

func (b *ClaimBuilder) setType(system, code, display string) *ClaimBuilder {
	b.Data["type"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

func (b *ClaimBuilder) setSubType(system, code, display string) *ClaimBuilder {
	b.Data["subType"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

func (b *ClaimBuilder) setUse(use string) *ClaimBuilder {
	b.Data["use"] = use
	return b
}

func (b *ClaimBuilder) setPatient(reference, display string) *ClaimBuilder {
	patient := map[string]interface{}{"reference": reference}
	if display != "" {
		patient["display"] = display
	}
	b.Data["patient"] = patient
	return b
}

func (b *ClaimBuilder) setBillablePeriod(start, end string) *ClaimBuilder {
	b.Data["billablePeriod"] = map[string]interface{}{"start": start}
	if end != "" {
		b.Data["billablePeriod"].(map[string]interface{})["end"] = end
	}
	return b
}

func (b *ClaimBuilder) setCreated(dateTime string) *ClaimBuilder {
	b.Data["created"] = dateTime
	return b
}

func (b *ClaimBuilder) setProvider(reference string) *ClaimBuilder {
	b.Data["provider"] = map[string]interface{}{"reference": reference}
	return b
}

func (b *ClaimBuilder) setInsurer(reference string) *ClaimBuilder {
	b.Data["insurer"] = map[string]interface{}{"reference": reference}
	return b
}

func (b *ClaimBuilder) setFacility(reference, display string) *ClaimBuilder {
	facility := map[string]interface{}{"reference": reference}
	if display != "" {
		facility["display"] = display
	}
	b.Data["facility"] = facility
	return b
}

func (b *ClaimBuilder) addRelated(referenceSystem, referenceCode, relationshipSystem, relationshipCode string) *ClaimBuilder {
	related := map[string]interface{}{}
	if referenceSystem != "" && referenceCode != "" {
		related["reference"] = map[string]interface{}{"system": referenceSystem, "code": referenceCode}
	}
	if relationshipSystem != "" && relationshipCode != "" {
		related["relationship"] = map[string]interface{}{"coding": []interface{}{map[string]interface{}{"system": relationshipSystem, "code": relationshipCode}}}
	}
	b.Data["related"] = append(b.Data["related"].([]interface{}), related)
	return b
}

func (b *ClaimBuilder) addItem(sequence int, productSystem, productCode, productDisplay string) *ClaimBuilder {
	item := map[string]interface{}{
		"sequence":            sequence,
		"productOrService":    map[string]interface{}{"coding": []interface{}{map[string]interface{}{"system": productSystem, "code": productCode, "display": productDisplay}}},
	}
	b.Data["item"] = append(b.Data["item"].([]interface{}), item)
	return b
}

func (b *ClaimBuilder) addItemServicedPeriod(itemIdx int, start, end string) *ClaimBuilder {
	serviced := map[string]interface{}{"start": start}
	if end != "" {
		serviced["end"] = end
	}
	items := b.Data["item"].([]interface{})
	if itemIdx < len(items) {
		items[itemIdx].(map[string]interface{})["servicedPeriod"] = serviced
	}
	return b
}

func (b *ClaimBuilder) addItemQuantity(itemIdx int, value float64, code string) *ClaimBuilder {
	items := b.Data["item"].([]interface{})
	if itemIdx < len(items) {
		items[itemIdx].(map[string]interface{})["net"] = map[string]interface{}{"value": value, "currency": code}
	}
	return b
}

func (b *ClaimBuilder) addItemUdi(itemIdx int, reference string) *ClaimBuilder {
	items := b.Data["item"].([]interface{})
	if itemIdx < len(items) {
		items[itemIdx].(map[string]interface{})["udi"] = []interface{}{map[string]interface{}{"issuer": "url", "value": reference}}
	}
	return b
}

func (b *ClaimBuilder) addCareTeam(sequence int, provider, role string) *ClaimBuilder {
	careTeam := map[string]interface{}{
		"sequence": sequence,
		"provider": map[string]interface{}{"reference": provider},
	}
	if role != "" {
		careTeam["role"] = map[string]interface{}{"coding": []interface{}{map[string]interface{}{"code": role}}}
	}
	b.Data["careTeam"] = append(b.Data["careTeam"].([]interface{}), careTeam)
	return b
}

func (b *ClaimBuilder) addDiagnosis(sequence int, system, code string) *ClaimBuilder {
	diagnosis := map[string]interface{}{
		"sequence":                     sequence,
		"diagnosisCodeableConcept":     map[string]interface{}{"coding": []interface{}{map[string]interface{}{"system": system, "code": code}}},
	}
	b.Data["diagnosis"] = append(b.Data["diagnosis"].([]interface{}), diagnosis)
	return b
}

func (b *ClaimBuilder) addProcedure(sequence int, system, code string) *ClaimBuilder {
	procedure := map[string]interface{}{
		"sequence":                  sequence,
		"procedureCodeableConcept":  map[string]interface{}{"coding": []interface{}{map[string]interface{}{"system": system, "code": code}}},
	}
	b.Data["procedure"] = append(b.Data["procedure"].([]interface{}), procedure)
	return b
}

func (b *ClaimBuilder) addInsurance(accountRef string) *ClaimBuilder {
	b.Data["insurance"] = append(b.Data["insurance"].([]interface{}), map[string]interface{}{"account": map[string]interface{}{"reference": accountRef}})
	return b
}

func (b *ClaimBuilder) setTotal(value float64, currency string) *ClaimBuilder {
	b.Data["total"] = map[string]interface{}{"value": value, "currency": currency}
	return b
}

func (b *ClaimBuilder) AddExtension(url string, value interface{}, valueType string) *ClaimBuilder {
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

func (b *ClaimBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ClaimBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}
