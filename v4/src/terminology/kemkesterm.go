package terminology

import "time"

type KemkesTerm struct {
	ID            int64      `json:"id"`
	ResourceType  string     `json:"resource_type"`
	AttributePath string     `json:"attribute_path"`
	Code          string     `json:"code"`
	ParentCode    *string    `json:"parent_code,omitempty"`
	Display       string     `json:"display"`
	DisplayEn     string     `json:"display_en"`
	CodeSystem    string     `json:"code_system"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}
