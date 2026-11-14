package builder

import (
	"encoding/json"
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
