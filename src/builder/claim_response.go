package builder

import (
	"encoding/json"
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type ClaimResponseBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewClaimResponseBuilder() *ClaimResponseBuilder {
	b := &ClaimResponseBuilder{ResourceType: "ClaimResponse", Data: make(map[string]interface{})}
	return b
}

func (b *ClaimResponseBuilder) setMetaProfile(profile string) *ClaimResponseBuilder {
	if _, ok := b.Data["meta"]; !ok {
		b.Data["meta"] = make(map[string]interface{})
	}
	b.Data["meta"].(map[string]interface{})["profile"] = profile
	return b
}

func (b *ClaimResponseBuilder) setId(id string) *ClaimResponseBuilder {
	b.Data["id"] = id
	return b
}

func (b *ClaimResponseBuilder) addIdentifier(identifier *datatype.Identifier) *ClaimResponseBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), identifier.ToArray())
	return b
}

func (b *ClaimResponseBuilder) setStatus(status string) *ClaimResponseBuilder {
	b.Data["status"] = status
	return b
}

func (b *ClaimResponseBuilder) setType(system, code, display string) *ClaimResponseBuilder {
	b.Data["type"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

func (b *ClaimResponseBuilder) setSubType(system, code, display string) *ClaimResponseBuilder {
	b.Data["subType"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

func (b *ClaimResponseBuilder) setUse(use string) *ClaimResponseBuilder {
	b.Data["use"] = use
	return b
}

func (b *ClaimResponseBuilder) setPatient(reference string) *ClaimResponseBuilder {
	b.Data["patient"] = map[string]interface{}{"reference": reference}
	return b
}

func (b *ClaimResponseBuilder) setCreated(dateTime string) *ClaimResponseBuilder {
	b.Data["created"] = dateTime
	return b
}

func (b *ClaimResponseBuilder) setInsurer(reference string) *ClaimResponseBuilder {
	b.Data["insurer"] = map[string]interface{}{"reference": reference}
	return b
}

func (b *ClaimResponseBuilder) setRequestor(reference string) *ClaimResponseBuilder {
	b.Data["requestor"] = map[string]interface{}{"reference": reference}
	return b
}

func (b *ClaimResponseBuilder) setRequest(reference string) *ClaimResponseBuilder {
	b.Data["request"] = map[string]interface{}{"reference": reference}
	return b
}

func (b *ClaimResponseBuilder) setOutcome(outcome string) *ClaimResponseBuilder {
	b.Data["outcome"] = outcome
	return b
}

func (b *ClaimResponseBuilder) setDisposition(disposition string) *ClaimResponseBuilder {
	b.Data["disposition"] = disposition
	return b
}

func (b *ClaimResponseBuilder) setPreAuthRef(ref string) *ClaimResponseBuilder {
	b.Data["preAuthRef"] = ref
	return b
}

func (b *ClaimResponseBuilder) setPreAuthPeriod(start, end string) *ClaimResponseBuilder {
	b.Data["preAuthPeriod"] = map[string]interface{}{"start": start}
	if end != "" {
		b.Data["preAuthPeriod"].(map[string]interface{})["end"] = end
	}
	return b
}

func (b *ClaimResponseBuilder) setPayeeType(system, code, display string) *ClaimResponseBuilder {
	b.Data["payeeType"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

func (b *ClaimResponseBuilder) addItem(itemSequence, detailSequence, subDetailSequence int) *ClaimResponseBuilder {
	item := map[string]interface{}{"itemSequence": itemSequence}
	if detailSequence > 0 {
		item["detailSequence"] = detailSequence
	}
	if subDetailSequence > 0 {
		item["subDetailSequence"] = subDetailSequence
	}
	b.Data["item"] = append(b.Data["item"].([]interface{}), item)
	return b
}

func (b *ClaimResponseBuilder) addItemAdjudication(itemSequence, categorySystem, categoryCode int, value float64) *ClaimResponseBuilder {
	adjudication := map[string]interface{}{
		"itemSequence": itemSequence,
		"category":     map[string]interface{}{"coding": []interface{}{map[string]interface{}{"system": categorySystem, "code": categoryCode}}},
		"value":        value,
	}
	b.Data["item"] = append(b.Data["item"].([]interface{}), adjudication)
	return b
}

func (b *ClaimResponseBuilder) addError(itemSequence, codeSystem, code, display string) *ClaimResponseBuilder {
	b.Data["error"] = append(b.Data["error"].([]interface{}), map[string]interface{}{
		"itemSequence": itemSequence,
		"code":         map[string]interface{}{"coding": []interface{}{map[string]interface{}{"system": codeSystem, "code": code, "display": display}}},
	})
	return b
}

func (b *ClaimResponseBuilder) setPayment(amount float64, paymentType, date string) *ClaimResponseBuilder {
	b.Data["payment"] = map[string]interface{}{
		"amount": map[string]interface{}{"value": amount},
		"type":   map[string]interface{}{"coding": []interface{}{map[string]interface{}{"code": paymentType}}},
		"date":   date,
	}
	return b
}

func (b *ClaimResponseBuilder) setNote(text string) *ClaimResponseBuilder {
	b.Data["formNote"] = append(b.Data["formNote"].([]interface{}), map[string]interface{}{"text": text})
	return b
}

func (b *ClaimResponseBuilder) AddExtension(url string, value interface{}, valueType string) *ClaimResponseBuilder {
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

func (b *ClaimResponseBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ClaimResponseBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}
