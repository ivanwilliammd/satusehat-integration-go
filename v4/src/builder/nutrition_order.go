package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type NutritionOrderBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewNutritionOrderBuilder() *NutritionOrderBuilder {
	b := &NutritionOrderBuilder{ResourceType: "NutritionOrder", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "NutritionOrder"
	return b
}

func (b *NutritionOrderBuilder) SetId(id string) *NutritionOrderBuilder {
	b.Data["id"] = id
	return b
}

func (b *NutritionOrderBuilder) AddIdentifier(identifier *datatype.Identifier) *NutritionOrderBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *NutritionOrderBuilder) SetStatus(status string) *NutritionOrderBuilder {
	b.Data["status"] = status
	return b
}

func (b *NutritionOrderBuilder) SetSubject(reference string, display ...string) *NutritionOrderBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *NutritionOrderBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *NutritionOrderBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetEncounter sets encounter from a Reference, as the PHP builder does.
func (b *NutritionOrderBuilder) SetEncounter(encounter *datatype.Reference) *NutritionOrderBuilder {
	b.Data["encounter"] = encounter.ToArray()
	return b
}

// SetPatient sets patient from a Reference, as the PHP builder does.
func (b *NutritionOrderBuilder) SetPatient(patient *datatype.Reference) *NutritionOrderBuilder {
	b.Data["patient"] = patient.ToArray()
	return b
}

// SetIntent sets intent.
func (b *NutritionOrderBuilder) SetIntent(intent string) *NutritionOrderBuilder {
	b.Data["intent"] = intent
	return b
}

// SetDateTime sets dateTime.
func (b *NutritionOrderBuilder) SetDateTime(dateTime string) *NutritionOrderBuilder {
	b.Data["dateTime"] = dateTime
	return b
}

// SetOrderer sets orderer from a reference.
func (b *NutritionOrderBuilder) SetOrderer(orderer *datatype.Reference) *NutritionOrderBuilder {
	b.Data["orderer"] = orderer.ToArray()
	return b
}

// AddAllergyType appends an allergy type coding. The canonical builder writes it to "allergyType"; FHIR R4
// names the element allergyIntolerance, so this reproduces the canonical key rather than a FHIR-correct one.
func (b *NutritionOrderBuilder) AddAllergyType(allergyType *datatype.CodeableConcept) *NutritionOrderBuilder {
	pushPath(b.Data, "allergyType", allergyType.ToArray())
	return b
}

// AddFoodPreferenceModifier appends a food preference modifier.
func (b *NutritionOrderBuilder) AddFoodPreferenceModifier(modifier *datatype.CodeableConcept) *NutritionOrderBuilder {
	pushPath(b.Data, "foodPreferenceModifier", modifier.ToArray())
	return b
}

// AddExcludeFoodModifier appends an excluded food modifier.
func (b *NutritionOrderBuilder) AddExcludeFoodModifier(modifier *datatype.CodeableConcept) *NutritionOrderBuilder {
	pushPath(b.Data, "excludeFoodModifier", modifier.ToArray())
	return b
}

// SetOralDietType appends a type to the oralDiet block, creating the block when absent.
func (b *NutritionOrderBuilder) SetOralDietType(dietType *datatype.CodeableConcept) *NutritionOrderBuilder {
	pushPath(b.Data, "oralDiet/type", dietType.ToArray())
	return b
}

// AddOralDietSchedule appends a timing to the oralDiet block, creating the block when absent.
func (b *NutritionOrderBuilder) AddOralDietSchedule(schedule *datatype.Timing) *NutritionOrderBuilder {
	pushPath(b.Data, "oralDiet/schedule", schedule.ToArray())
	return b
}

// AddSupplement appends a supplement the caller has already shaped.
func (b *NutritionOrderBuilder) AddSupplement(supplement map[string]interface{}) *NutritionOrderBuilder {
	pushPath(b.Data, "supplement", supplement)
	return b
}

// SetEnteralFormula sets the enteral formula the caller has already shaped.
func (b *NutritionOrderBuilder) SetEnteralFormula(enteralFormula map[string]interface{}) *NutritionOrderBuilder {
	b.Data["enteralFormula"] = enteralFormula
	return b
}
