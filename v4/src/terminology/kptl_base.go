package terminology

import "time"

// KptlBase represents a procedure/tindakan base code
type KptlBase struct {
	ID          int64      `json:"id"`
	Status      string     `json:"status"`
	BaseCode    string     `json:"base_code"`
	BaseDisplay string     `json:"base_display"`
	Modifier1   *string    `json:"modifier_1,omitempty"`
	Modifier2   *string    `json:"modifier_2,omitempty"`
	Modifier3   *string    `json:"modifier_3,omitempty"`
	Modifier4   *string    `json:"modifier_4,omitempty"`
	Modifier5   *string    `json:"modifier_5,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
