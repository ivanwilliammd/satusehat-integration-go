package builder

import (
	"encoding/json"
	"fmt"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ClaimResponseBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewClaimResponseBuilder() *ClaimResponseBuilder {
	b := &ClaimResponseBuilder{ResourceType: "ClaimResponse", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "ClaimResponse"
	return b
}

func (b *ClaimResponseBuilder) SetId(id string) *ClaimResponseBuilder {
	b.Data["id"] = id
	return b
}

func (b *ClaimResponseBuilder) AddIdentifier(identifier *datatype.Identifier) *ClaimResponseBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ClaimResponseBuilder) SetStatus(status string) *ClaimResponseBuilder {
	b.Data["status"] = status
	return b
}

func (b *ClaimResponseBuilder) SetSubject(reference string, display ...string) *ClaimResponseBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *ClaimResponseBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ClaimResponseBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetPatient sets patient by reference, as the PHP builder does.
func (b *ClaimResponseBuilder) SetPatient(reference string) *ClaimResponseBuilder {
	b.Data["patient"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

// SetInsurer sets insurer by reference, as the PHP builder does.
func (b *ClaimResponseBuilder) SetInsurer(reference string) *ClaimResponseBuilder {
	b.Data["insurer"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

// SetCreated sets created, as the PHP builder does.
func (b *ClaimResponseBuilder) SetCreated(dateTime string) *ClaimResponseBuilder {
	b.Data["created"] = dateTime
	return b
}

// SetType sets type with a single coding, as the PHP builder does.
func (b *ClaimResponseBuilder) SetType(system, code, display string) *ClaimResponseBuilder {
	b.Data["type"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// AddCoverage adds a coverage entry with its sequence, as the PHP builder does.
func (b *ClaimResponseBuilder) AddCoverage(sequence int, reference string, focal *bool, prioritySystem, priorityCode *string) *ClaimResponseBuilder {
	entri := map[string]interface{}{"sequence": sequence, "reference": reference}
	if focal != nil {
		entri["focal"] = *focal
	}
	if prioritySystem != nil && priorityCode != nil {
		entri["priority"] = map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": *prioritySystem, "code": *priorityCode}},
		}
	}
	daftar, _ := b.Data["coverage"].([]interface{})
	b.Data["coverage"] = append(daftar, entri)
	return b
}

// SetOutcome sets the outcome code, as the PHP builder does.
func (b *ClaimResponseBuilder) SetOutcome(outcome string) *ClaimResponseBuilder {
	b.Data["outcome"] = outcome
	return b
}

// SetRequest sets request by reference, as the PHP builder does.
func (b *ClaimResponseBuilder) SetRequest(reference string) *ClaimResponseBuilder {
	b.Data["request"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

// SetDisposition sets disposition, as the PHP builder does.
func (b *ClaimResponseBuilder) SetDisposition(disposition string) *ClaimResponseBuilder {
	b.Data["disposition"] = disposition
	return b
}

// AddItem adds an adjudication item with its sequence links, as the PHP builder does.
func (b *ClaimResponseBuilder) AddItem(itemSequence int, detailSequence, subDetailSequence *int) *ClaimResponseBuilder {
	item := map[string]interface{}{"itemSequence": itemSequence}
	if detailSequence != nil {
		item["detailSequence"] = *detailSequence
	}
	if subDetailSequence != nil {
		item["subDetailSequence"] = *subDetailSequence
	}
	daftar, _ := b.Data["item"].([]interface{})
	b.Data["item"] = append(daftar, item)
	return b
}

// AddAddItemAdjudication adds an adjudication to the most recent addItem, as the PHP builder does.
func (b *ClaimResponseBuilder) AddAddItemAdjudication(sequenceLinkId int, categorySystem, categoryCode string, value *float64) *ClaimResponseBuilder {
	daftar, _ := b.Data["addItem"].([]interface{})
	if len(daftar) == 0 {
		return b
	}
	itemIdx := len(daftar) - 1
	adj := map[string]interface{}{
		"itemSequence": sequenceLinkId,
		"category": map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": categorySystem, "code": categoryCode}},
		},
	}
	if value != nil {
		adj["value"] = *value
	}
	pushPath(b.Data, fmt.Sprintf("addItem/%d/adjudication", itemIdx), adj)
	return b
}

// AddAddItem adds a line the payer is adding; itemSequence is a list, as the canonical builder writes it.
func (b *ClaimResponseBuilder) AddAddItem(itemSequence int, productOrServiceSystem, productOrServiceCode, productOrServiceDisplay string) *ClaimResponseBuilder {
	pushPath(b.Data, "addItem", map[string]interface{}{
		"itemSequence": []interface{}{itemSequence},
		"productOrService": map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{
				"system": productOrServiceSystem, "code": productOrServiceCode, "display": productOrServiceDisplay,
			}},
		},
	})
	return b
}

// AddError adds a processing error for a line.
func (b *ClaimResponseBuilder) AddError(itemSequence int, codeSystem, code, display string) *ClaimResponseBuilder {
	pushPath(b.Data, "error", map[string]interface{}{
		"itemSequence": itemSequence,
		"code": map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": codeSystem, "code": code, "display": display}},
		},
	})
	return b
}

// AddItemAdjudication adds an adjudication line. The canonical builder appends it to "item", not to a
// key named adjudication, so the payload matches what the PHP SDK sends.
func (b *ClaimResponseBuilder) AddItemAdjudication(itemSequence int, categorySystem, categoryCode string, value *float64, reasonSystem, reasonCode *string) *ClaimResponseBuilder {
	adjudication := map[string]interface{}{
		"itemSequence": itemSequence,
		"category": map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": categorySystem, "code": categoryCode}},
		},
	}
	if value != nil {
		adjudication["value"] = *value
	}
	if reasonSystem != nil && reasonCode != nil {
		adjudication["reason"] = map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": *reasonSystem, "code": *reasonCode}},
		}
	}
	pushPath(b.Data, "item", adjudication)
	return b
}

// SetPayeeType sets payeeType with a single coding.
func (b *ClaimResponseBuilder) SetPayeeType(system, code, display string) *ClaimResponseBuilder {
	b.Data["payeeType"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetPayment writes the payment block through nested paths, as the canonical builder does.
func (b *ClaimResponseBuilder) SetPayment(value float64, typeSystem, typeCode, date string, reference *string) *ClaimResponseBuilder {
	setPath(b.Data, "payment/type", map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": typeSystem, "code": typeCode}},
	})
	setPath(b.Data, "payment/date", date)
	setPath(b.Data, "payment/amount/value", value)
	if reference != nil && *reference != "" {
		setPath(b.Data, "payment/reference", *reference)
	}
	return b
}

// SetPreAuthPeriod writes the pre-authorisation period through nested paths.
func (b *ClaimResponseBuilder) SetPreAuthPeriod(start, end string) *ClaimResponseBuilder {
	setPath(b.Data, "preAuthPeriod/start", start)
	if end != "" {
		setPath(b.Data, "preAuthPeriod/end", end)
	}
	return b
}

// SetPreAuthRef sets preAuthRef.
func (b *ClaimResponseBuilder) SetPreAuthRef(preAuthRef string) *ClaimResponseBuilder {
	b.Data["preAuthRef"] = preAuthRef
	return b
}

// SetRequestor sets requestor by reference.
func (b *ClaimResponseBuilder) SetRequestor(reference string) *ClaimResponseBuilder {
	b.Data["requestor"] = map[string]interface{}{"reference": reference}
	return b
}

// SetSubType sets subType with a single coding.
func (b *ClaimResponseBuilder) SetSubType(system, code, display string) *ClaimResponseBuilder {
	b.Data["subType"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetTotal adds a total with its category. Named "set" in the canonical builder but it pushes, so repeated
// calls accumulate one total per category.
func (b *ClaimResponseBuilder) SetTotal(value float64, categorySystem, categoryCode string) *ClaimResponseBuilder {
	pushPath(b.Data, "total", map[string]interface{}{
		"category": map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": categorySystem, "code": categoryCode}},
		},
		"value": value,
	})
	return b
}

// SetUse sets use.
func (b *ClaimResponseBuilder) SetUse(use string) *ClaimResponseBuilder {
	b.Data["use"] = use
	return b
}
