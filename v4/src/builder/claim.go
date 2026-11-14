package builder

import (
	"encoding/json"
	"fmt"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ClaimBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewClaimBuilder() *ClaimBuilder {
	b := &ClaimBuilder{ResourceType: "Claim", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Claim"
	return b
}

func (b *ClaimBuilder) SetId(id string) *ClaimBuilder {
	b.Data["id"] = id
	return b
}

func (b *ClaimBuilder) AddIdentifier(identifier *datatype.Identifier) *ClaimBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ClaimBuilder) SetStatus(status string) *ClaimBuilder {
	b.Data["status"] = status
	return b
}

func (b *ClaimBuilder) SetSubject(reference string, display ...string) *ClaimBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *ClaimBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ClaimBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetPatient sets patient by reference, with an optional display as the PHP builder allows.
func (b *ClaimBuilder) SetPatient(reference string, display ...string) *ClaimBuilder {
	ref := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		ref.Display = display[0]
	}
	b.Data["patient"] = ref.ToArray()
	return b
}

// SetInsurer sets insurer by reference, as the PHP builder does.
func (b *ClaimBuilder) SetInsurer(reference string) *ClaimBuilder {
	b.Data["insurer"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

// SetCreated sets created, as the PHP builder does.
func (b *ClaimBuilder) SetCreated(dateTime string) *ClaimBuilder {
	b.Data["created"] = dateTime
	return b
}

// SetType sets type with a single coding, as the PHP builder does.
func (b *ClaimBuilder) SetType(system, code, display string) *ClaimBuilder {
	b.Data["type"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// AddCoverage adds a coverage entry, as the PHP builder does.
func (b *ClaimBuilder) AddCoverage(reference string, focal *bool, prioritySystem, priorityCode *string) *ClaimBuilder {
	entri := map[string]interface{}{"reference": reference}
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

// SetEnterer sets enterer by reference, as the PHP builder does.
func (b *ClaimBuilder) SetEnterer(reference string) *ClaimBuilder {
	b.Data["enterer"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

// SetProvider sets provider by reference, as the PHP builder does.
func (b *ClaimBuilder) SetProvider(reference string) *ClaimBuilder {
	b.Data["provider"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

// AddItem adds an item line with its coding and the optional sequence links, as the PHP builder does.
func (b *ClaimBuilder) AddItem(sequence int, system, code, display string, careTeamSequence, diagnosisSequence, procedureSequence, informationSequence *int) *ClaimBuilder {
	item := map[string]interface{}{
		"sequence": sequence,
		"productOrService": map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
		},
	}
	for kunci, nilai := range map[string]*int{
		"careTeamSequence":    careTeamSequence,
		"diagnosisSequence":   diagnosisSequence,
		"procedureSequence":   procedureSequence,
		"informationSequence": informationSequence,
	} {
		if nilai != nil {
			item[kunci] = []interface{}{*nilai}
		}
	}
	daftar, _ := b.Data["item"].([]interface{})
	b.Data["item"] = append(daftar, item)
	return b
}

// AddItemQuantity sets the quantity of an existing item line; the line must exist first, as in PHP.
func (b *ClaimBuilder) AddItemQuantity(itemIdx int, value float64, code string) *ClaimBuilder {
	setPath(b.Data, fmt.Sprintf("item/%d/quantity/value", itemIdx), value)
	setPath(b.Data, fmt.Sprintf("item/%d/quantity/currency", itemIdx), code)
	return b
}

// AddItemUnitPrice sets the unit price of an existing item line.
func (b *ClaimBuilder) AddItemUnitPrice(itemIdx int, value float64, code string) *ClaimBuilder {
	setPath(b.Data, fmt.Sprintf("item/%d/unitPrice/value", itemIdx), value)
	setPath(b.Data, fmt.Sprintf("item/%d/unitPrice/currency", itemIdx), code)
	return b
}

// AddItemServicedPeriod sets the service period of an existing item line.
func (b *ClaimBuilder) AddItemServicedPeriod(itemIdx int, start, end string) *ClaimBuilder {
	setPath(b.Data, fmt.Sprintf("item/%d/servicedPeriod/start", itemIdx), start)
	if end != "" {
		setPath(b.Data, fmt.Sprintf("item/%d/servicedPeriod/end", itemIdx), end)
	}
	return b
}

// AddItemLocation sets the coding of an existing item line's location.
func (b *ClaimBuilder) AddItemLocation(itemIdx int, system, code, display string) *ClaimBuilder {
	pushPath(b.Data, fmt.Sprintf("item/%d/location", itemIdx), map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	})
	return b
}

// AddCareTeam adds a care team member; the role coding is written only when both parts are given.
func (b *ClaimBuilder) AddCareTeam(sequence int, providerReference string, roleSystem, roleCode, roleDisplay *string) *ClaimBuilder {
	anggota := map[string]interface{}{"sequence": sequence, "provider": map[string]interface{}{"reference": providerReference}}
	if roleSystem != nil && roleCode != nil {
		display := ""
		if roleDisplay != nil {
			display = *roleDisplay
		}
		anggota["role"] = map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": *roleSystem, "code": *roleCode, "display": display}},
		}
	}
	pushPath(b.Data, "careTeam", anggota)
	return b
}

// AddDiagnosis adds a diagnosis line with an optional type coding.
func (b *ClaimBuilder) AddDiagnosis(sequence int, system, code string, typeSystem, typeCode *string) *ClaimBuilder {
	diagnosis := map[string]interface{}{
		"sequence": sequence,
		"diagnosisCodeableConcept": map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": system, "code": code}},
		},
	}
	if typeSystem != nil && typeCode != nil {
		diagnosis["type"] = map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": *typeSystem, "code": *typeCode}},
		}
	}
	pushPath(b.Data, "diagnosis", diagnosis)
	return b
}

// AddProcedure adds a procedure line with an optional date and type coding.
func (b *ClaimBuilder) AddProcedure(sequence int, system, code string, date, typeSystem, typeCode *string) *ClaimBuilder {
	procedure := map[string]interface{}{
		"sequence": sequence,
		"procedureCodeableConcept": map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": system, "code": code}},
		},
	}
	if date != nil && *date != "" {
		procedure["date"] = *date
	}
	if typeSystem != nil && typeCode != nil {
		procedure["type"] = map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": *typeSystem, "code": *typeCode}},
		}
	}
	pushPath(b.Data, "procedure", procedure)
	return b
}

// AddInsuranceAccount adds an account reference to the insurance block, writing the nested path the
// canonical builder uses ("insurance/account").
func (b *ClaimBuilder) AddInsuranceAccount(reference string) *ClaimBuilder {
	pushPath(b.Data, "insurance/account", map[string]interface{}{"reference": reference})
	return b
}

// AddRelated mirrors the canonical builder, which writes "reference" as a coding object rather than a
// Reference. FHIR R4 declares related.reference as a Reference, so this ports a deviation and not a
// correct shape; it is kept so a payload matches the PHP SDK byte for byte.
func (b *ClaimBuilder) AddRelated(referenceSystem, referenceCode, referenceDisplay, relationshipSystem, relationshipCode *string) *ClaimBuilder {
	related := map[string]interface{}{}
	if referenceSystem != nil && referenceCode != nil {
		display := ""
		if referenceDisplay != nil {
			display = *referenceDisplay
		}
		related["reference"] = map[string]interface{}{"system": *referenceSystem, "code": *referenceCode, "display": display}
	}
	if relationshipSystem != nil && relationshipCode != nil {
		related["relationship"] = map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": *relationshipSystem, "code": *relationshipCode}},
		}
	}
	pushPath(b.Data, "related", related)
	return b
}

// SetBillablePeriod writes the billable period through nested paths, as the canonical builder does.
func (b *ClaimBuilder) SetBillablePeriod(start, end string) *ClaimBuilder {
	setPath(b.Data, "billablePeriod/start", start)
	if end != "" {
		setPath(b.Data, "billablePeriod/end", end)
	}
	return b
}

// SetFacility sets facility with an optional display.
func (b *ClaimBuilder) SetFacility(reference, display string) *ClaimBuilder {
	facility := map[string]interface{}{"reference": reference}
	if display != "" {
		facility["display"] = display
	}
	b.Data["facility"] = facility
	return b
}

// SetPrescriber sets prescriber by reference.
func (b *ClaimBuilder) SetPrescriber(reference string) *ClaimBuilder {
	b.Data["prescriber"] = map[string]interface{}{"reference": reference}
	return b
}

// SetPriority sets priority with a single coding; an empty display is written, as the canonical builder does.
func (b *ClaimBuilder) SetPriority(system, code, display string) *ClaimBuilder {
	b.Data["priority"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetPriorityPayor adds a payor reference. Named "set" in the canonical builder but it pushes, so repeated
// calls accumulate.
func (b *ClaimBuilder) SetPriorityPayor(reference string) *ClaimBuilder {
	pushPath(b.Data, "priorityPayor", map[string]interface{}{"reference": reference})
	return b
}

// SetSubType sets subType with a single coding.
func (b *ClaimBuilder) SetSubType(system, code, display string) *ClaimBuilder {
	b.Data["subType"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetTotal sets the claimed total with its currency.
func (b *ClaimBuilder) SetTotal(value float64, code string) *ClaimBuilder {
	b.Data["total"] = map[string]interface{}{"value": value, "currency": code}
	return b
}

// SetUse sets use.
func (b *ClaimBuilder) SetUse(use string) *ClaimBuilder {
	b.Data["use"] = use
	return b
}
