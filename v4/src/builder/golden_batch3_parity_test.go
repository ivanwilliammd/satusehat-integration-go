package builder_test

// Goldens for the ten medicinal-product builders, none of which had one. The group is uniform: identifier,
// status, one reference and two codings — the least likely place for a defect, which is why proving it costs
// little and settles ten resources at once.

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
)

func TestMedicinalProductMatchesPHPPayload(t *testing.T) {
	b := builder.NewMedicinalProductBuilder()
	b.SetId("mp-1")
	b.AddIdentifier("http://sys.id/ni", "NI-20")
	b.SetStatus("active")
	b.SetName("Amoksisilin")
	b.SetManufacturer("Organization/1000000001", "Kimia Farma")
	b.SetProductClass("http://terminology.hl7.org/CodeSystem/medicinal-product-class", "antibiotic", "Antibiotic")
	samakan(t, bacaGolden(t, "medicinal_product"), b)
}

func TestMedicinalProductAuthorizationMatchesPHPPayload(t *testing.T) {
	b := builder.NewMedicinalProductAuthorizationBuilder()
	b.SetId("mpa-1")
	b.AddIdentifier("http://sys.id/ni", "NI-21")
	b.SetStatus("active")
	b.SetCountry("urn:iso:std:iso:3166", "ID", "Indonesia")
	b.SetRegulator("Organization/1000000002", "BPOM")
	b.SetValidityPeriod("2026-01-05/2031-01-05")
	samakan(t, bacaGolden(t, "medicinal_product_authorization"), b)
}

func TestMedicinalProductContraindicationMatchesPHPPayload(t *testing.T) {
	b := builder.NewMedicinalProductContraindicationBuilder()
	b.SetId("mpc-1")
	b.AddIdentifier("http://sys.id/ni", "NI-22")
	b.SetStatus("active")
	b.SetSubject("MedicinalProduct/mp-1", "Amoksisilin")
	b.SetDisease("http://snomed.info/sct", "91936005", "Alergi penisilin")
	b.SetComorbidity("Condition/cond-9", "Gagal ginjal")
	samakan(t, bacaGolden(t, "medicinal_product_contraindication"), b)
}

func TestMedicinalProductIndicationMatchesPHPPayload(t *testing.T) {
	b := builder.NewMedicinalProductIndicationBuilder()
	b.SetId("mpi-1")
	b.AddIdentifier("http://sys.id/ni", "NI-23")
	b.SetStatus("active")
	b.SetSubject("MedicinalProduct/mp-1", "Amoksisilin")
	b.SetDisease("http://snomed.info/sct", "233604007", "Pneumonia")
	b.SetComorbidity("Condition/cond-10", "Diabetes")
	samakan(t, bacaGolden(t, "medicinal_product_indication"), b)
}

func TestMedicinalProductIngredientMatchesPHPPayload(t *testing.T) {
	b := builder.NewMedicinalProductIngredientBuilder()
	b.SetId("mpg-1")
	b.AddIdentifier("http://sys.id/ni", "NI-24")
	b.SetStatus("active")
	b.SetRole("http://terminology.hl7.org/CodeSystem/medicinal-product-ingredient-role", "active", "Active")
	b.SetSubstance("Amoksisilin trihidrat")
	b.SetQuantity("500 mg")
	samakan(t, bacaGolden(t, "medicinal_product_ingredient"), b)
}

func TestMedicinalProductInteractionMatchesPHPPayload(t *testing.T) {
	b := builder.NewMedicinalProductInteractionBuilder()
	b.SetId("mpn-1")
	b.AddIdentifier("http://sys.id/ni", "NI-25")
	b.SetStatus("active")
	b.SetSubject("MedicinalProduct/mp-1", "Amoksisilin")
	b.SetDescription("Interaksi dengan warfarin")
	b.SetInteractant("http://snomed.info/sct", "372756006", "Warfarin")
	samakan(t, bacaGolden(t, "medicinal_product_interaction"), b)
}

func TestMedicinalProductManufacturedMatchesPHPPayload(t *testing.T) {
	b := builder.NewMedicinalProductManufacturedBuilder()
	b.SetId("mpm-1")
	b.AddIdentifier("http://sys.id/ni", "NI-26")
	b.SetStatus("active")
	b.SetQuantity("500 mg")
	b.SetManufacturer("Organization/1000000001", "Kimia Farma")
	b.SetIngredient("http://snomed.info/sct", "387517004", "Amoksisilin")
	samakan(t, bacaGolden(t, "medicinal_product_manufactured"), b)
}

func TestMedicinalProductPackagedMatchesPHPPayload(t *testing.T) {
	b := builder.NewMedicinalProductPackagedBuilder()
	b.SetId("mpp-1")
	b.AddIdentifier("http://sys.id/ni", "NI-27")
	b.SetStatus("active")
	b.SetPackageItem("http://terminology.hl7.org/CodeSystem/package-type", "box", "Box")
	b.SetDescription("Dus 10 tablet")
	b.SetManufacturer("Organization/1000000001", "Kimia Farma")
	samakan(t, bacaGolden(t, "medicinal_product_packaged"), b)
}

func TestMedicinalProductPharmaceuticalMatchesPHPPayload(t *testing.T) {
	b := builder.NewMedicinalProductPharmaceuticalBuilder()
	b.SetId("mph-1")
	b.AddIdentifier("http://sys.id/ni", "NI-28")
	b.SetStatus("active")
	b.SetRoute("http://snomed.info/sct", "26643006", "Oral")
	b.SetDoseForm("http://snomed.info/sct", "385055001", "Tablet")
	b.SetIngredient("http://snomed.info/sct", "387517004", "Amoksisilin")
	samakan(t, bacaGolden(t, "medicinal_product_pharmaceutical"), b)
}

func TestMedicinalProductUndesirableEffectMatchesPHPPayload(t *testing.T) {
	b := builder.NewMedicinalProductUndesirableEffectBuilder()
	b.SetId("mpu-1")
	b.AddIdentifier("http://sys.id/ni", "NI-29")
	b.SetStatus("active")
	b.SetSubject("MedicinalProduct/mp-1", "Amoksisilin")
	b.SetSymptomConditionEffect("http://snomed.info/sct", "422587007", "Nausea")
	b.SetClassification("http://terminology.hl7.org/CodeSystem/undesirable-effect-frequency", "common", "Common")
	samakan(t, bacaGolden(t, "medicinal_product_undesirable_effect"), b)
}
