package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ChargeItemBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewChargeItemBuilder() *ChargeItemBuilder {
	b := &ChargeItemBuilder{ResourceType: "ChargeItem", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "ChargeItem"
	return b
}

func (b *ChargeItemBuilder) SetId(id string) *ChargeItemBuilder {
	b.Data["id"] = id
	return b
}

func (b *ChargeItemBuilder) AddIdentifier(identifier *datatype.Identifier) *ChargeItemBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ChargeItemBuilder) SetStatus(status string) *ChargeItemBuilder {
	b.Data["status"] = status
	return b
}

func (b *ChargeItemBuilder) SetSubject(reference string, display ...string) *ChargeItemBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *ChargeItemBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ChargeItemBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *ChargeItemBuilder) SetCode(system, code, display string) *ChargeItemBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *ChargeItemBuilder) SetEncounter(reference string) *ChargeItemBuilder {
	// ChargeItem menyebut encounter sebagai "context"; menulis "encounter" menghasilkan medan yang tidak
	// dikenal server, jadi bentuk kanon yang diikuti.
	b.Data["context"] = map[string]interface{}{"reference": reference}
	return b
}

// SetEnterer sets enterer by reference, as the PHP builder does.
func (b *ChargeItemBuilder) SetEnterer(reference string) *ChargeItemBuilder {
	b.Data["enterer"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

// SetPartOf adds a partOf reference; the element is 0..* in FHIR and the PHP builder pushes.
func (b *ChargeItemBuilder) SetPartOf(reference string) *ChargeItemBuilder {
	pushPath(b.Data, "partOf", datatype.Reference{Reference: reference}.ToArray())
	return b
}

// SetFocus sets focus by reference.
func (b *ChargeItemBuilder) SetFocus(reference string) *ChargeItemBuilder {
	b.Data["focus"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

// SetPerformedDate sets the performed date time.
func (b *ChargeItemBuilder) SetPerformedDate(dateTime string) *ChargeItemBuilder {
	b.Data["performedDateTime"] = dateTime
	return b
}

// SetPerformerActor adds a performer with its actor and optional function coding.
func (b *ChargeItemBuilder) SetPerformerActor(reference string, functionSystem, functionCode *string) *ChargeItemBuilder {
	performer := map[string]interface{}{"actor": map[string]interface{}{"reference": reference}}
	if functionSystem != nil && functionCode != nil {
		performer["function"] = map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": *functionSystem, "code": *functionCode}},
		}
	}
	pushPath(b.Data, "performer", performer)
	return b
}

// SetPerformerCostCenter adds a performer entry without a function.
func (b *ChargeItemBuilder) SetPerformerCostCenter(reference string) *ChargeItemBuilder {
	pushPath(b.Data, "performer", map[string]interface{}{"actor": map[string]interface{}{"reference": reference}})
	return b
}

// SetOrganization sets the performing organization.
func (b *ChargeItemBuilder) SetOrganization(reference string) *ChargeItemBuilder {
	b.Data["performingOrganization"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

// SetRequestingOrganization sets the requesting organization.
func (b *ChargeItemBuilder) SetRequestingOrganization(reference string) *ChargeItemBuilder {
	b.Data["requestingOrganization"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

// SetQuantity sets the quantity as a plain number, as the PHP builder does.
func (b *ChargeItemBuilder) SetQuantity(quantity int) *ChargeItemBuilder {
	b.Data["quantity"] = quantity
	return b
}

// SetBodysite adds a body site coding.
func (b *ChargeItemBuilder) SetBodysite(system, code, display string) *ChargeItemBuilder {
	pushPath(b.Data, "bodysite", map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	})
	return b
}

// SetProductCode sets the product as a CodeableConcept.
func (b *ChargeItemBuilder) SetProductCode(system, code, display string) *ChargeItemBuilder {
	b.Data["productCodeableConcept"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetPriceOverride sets priceOverride with its currency.
func (b *ChargeItemBuilder) SetPriceOverride(value float64, code string) *ChargeItemBuilder {
	b.Data["priceOverride"] = map[string]interface{}{"value": value, "currency": code}
	return b
}

// SetOverrideReason sets overrideReason.
func (b *ChargeItemBuilder) SetOverrideReason(overrideReason string) *ChargeItemBuilder {
	b.Data["overrideReason"] = overrideReason
	return b
}

// SetEnteredDate sets enteredDate.
func (b *ChargeItemBuilder) SetEnteredDate(dateTime string) *ChargeItemBuilder {
	b.Data["enteredDate"] = dateTime
	return b
}

// SetReason adds a reason coding.
func (b *ChargeItemBuilder) SetReason(system, code, display string) *ChargeItemBuilder {
	pushPath(b.Data, "reason", map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	})
	return b
}

// SetService adds a service reference.
func (b *ChargeItemBuilder) SetService(reference string) *ChargeItemBuilder {
	pushPath(b.Data, "service", datatype.Reference{Reference: reference}.ToArray())
	return b
}

// SetAccount adds an account reference.
func (b *ChargeItemBuilder) SetAccount(reference string) *ChargeItemBuilder {
	pushPath(b.Data, "account", datatype.Reference{Reference: reference}.ToArray())
	return b
}
