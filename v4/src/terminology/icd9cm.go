package terminology

import "time"

type Icd9cm struct {
	ID         int64     `json:"id"`
	Icd9cmCode string    `json:"icd9cm_code"`
	Icd9cmEn   string    `json:"icd9cm_en"`
	Icd9cmID   *string   `json:"icd9cm_id,omitempty"`
	Active     bool      `json:"active"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}
