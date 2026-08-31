package terminology

import "time"

type Ucum struct {
	ID              int64      `json:"id"`
	Code            string     `json:"code"`
	DescriptiveName string     `json:"descriptive_name"`
	CodeSystem      string     `json:"code_system"`
	Definition      *string    `json:"definition,omitempty"`
	DateCreated     *time.Time `json:"date_created,omitempty"`
	Synonym         *string    `json:"synonym,omitempty"`
	Status          string     `json:"status"`
	KindOfQuantity  string     `json:"kind_of_quantity"`
	DateRevised     *time.Time `json:"date_revised,omitempty"`
	ConceptID       string     `json:"concept_id"`
	Dimension       string     `json:"dimension"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
}
