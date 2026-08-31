package terminology

import "time"

type Icd10 struct {
	ID       int64     `json:"id"`
	Icd10Code string   `json:"icd10_code"`
	Icd10En   string   `json:"icd10_en"`
	Icd10ID   *string  `json:"icd10_id,omitempty"`
	Active    bool     `json:"active"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
