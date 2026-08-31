package terminology

import "time"

type FhirR4Term struct {
	ID                     int64      `json:"id"`
	FhirID                 string     `json:"fhir_id"`
	URL                    string     `json:"url"`
	Version                string     `json:"version"`
	Name                   string     `json:"name"`
	Title                  string     `json:"title"`
	Status                 string     `json:"status"`
	Experimental           bool       `json:"experimental"`
	Description            string     `json:"description"`
	Date                   string     `json:"date"`
	Publisher              string     `json:"publisher"`
	Content                string     `json:"content"`
	ConceptCodeL1          *string    `json:"concept_code_l1,omitempty"`
	ConceptDisplayL1       *string    `json:"concept_display_l1,omitempty"`
	ConceptDefinitionL1     *string    `json:"concept_definition_l1,omitempty"`
	ConceptCodeL2          *string    `json:"concept_code_l2,omitempty"`
	ConceptDisplayL2       *string    `json:"concept_display_l2,omitempty"`
	ConceptDefinitionL2    *string    `json:"concept_definition_l2,omitempty"`
	CreatedAt              *time.Time `json:"created_at,omitempty"`
	UpdatedAt              *time.Time `json:"updated_at,omitempty"`
}
