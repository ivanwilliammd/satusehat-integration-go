package terminology

import "time"

// KptlBaseModifierMapping maps base codes to modifiers
type KptlBaseModifierMapping struct {
	ID            int64      `json:"id"`
	Display       string     `json:"display"`
	Modifier1     *string    `json:"modifier_1,omitempty"`
	Modifier2     *string    `json:"modifier_2,omitempty"`
	Modifier3     *string    `json:"modifier_3,omitempty"`
	Modifier4     *string    `json:"modifier_4,omitempty"`
	Modifier5     *string    `json:"modifier_5,omitempty"`
	BaseCode      string     `json:"base_code"`
	ModifierCode1 *string    `json:"modifier_code_1,omitempty"`
	ModifierCode2 *string    `json:"modifier_code_2,omitempty"`
	ModifierCode3 *string    `json:"modifier_code_3,omitempty"`
	ModifierCode4 *string    `json:"modifier_code_4,omitempty"`
	ModifierCode5 *string    `json:"modifier_code_5,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}
