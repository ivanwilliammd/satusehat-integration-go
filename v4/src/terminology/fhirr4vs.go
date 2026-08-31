package terminology

import "time"

type FhirR4Vs struct {
	ID                              int64      `json:"id"`
	FhirID                          string     `json:"fhir_id"`
	URL                             string     `json:"url"`
	Version                         string     `json:"version"`
	Name                            string     `json:"name"`
	Title                           string     `json:"title"`
	Status                          string     `json:"status"`
	Experimental                    bool       `json:"experimental"`
	Description                     string     `json:"description"`
	Date                            *string    `json:"date,omitempty"`
	Publisher                       string     `json:"publisher"`
	Immutable                       *bool      `json:"immutable,omitempty"`
	ComposeIncludeSystem            *string    `json:"compose_include_system,omitempty"`
	ComposeIncludeCode              *string    `json:"compose_include_code,omitempty"`
	ComposeIncludeDisplay           *string    `json:"compose_include_display,omitempty"`
	ComposeIncludeFilterProperty    *string    `json:"compose_include_filter_property,omitempty"`
	ComposeIncludeFilterOp          *string    `json:"compose_include_filter_op,omitempty"`
	ComposeIncludeFilterValue       *string    `json:"compose_include_filter_value,omitempty"`
	CreatedAt                       *time.Time `json:"created_at,omitempty"`
	UpdatedAt                       *time.Time `json:"updated_at,omitempty"`
}
