package terminology

import "time"

type Kfa struct {
	ID                int64      `json:"id"`
	KfaCode           string     `json:"kfa_code"`
	ProductTemplate   string     `json:"product_template"`
	DisplayName       string     `json:"display_name"`
	Brand             string     `json:"brand"`
	UomDrugForm       string     `json:"uom_drug_form"`
	DrugFormHL7       string     `json:"drug_form_hl7"`
	MedicationForm    string     `json:"medication_form"`
	MedicationFormCode string    `json:"medication_form_code"`
	LogisticDose      int        `json:"logistic_dose"`
	DrugClass         string     `json:"drug_class"`
	AtcClass          string     `json:"atc_class"`
	Fornas            bool       `json:"fornas"`
	LkppPrice         *float64   `json:"lkpp_price,omitempty"`
	IzinEdar          *string    `json:"izin_edar,omitempty"`
	Het               *float64   `json:"het,omitempty"`
	Manufacturer      string     `json:"manufacturer"`
	LkppShow          bool       `json:"lkpp_show"`
	Tag               *string    `json:"tag,omitempty"`
	Status            string     `json:"status"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
}
