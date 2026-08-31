package terminology

import "time"

type Snomedct struct {
	ID          int64      `json:"id"`
	Code        string     `json:"code"`
	FSN         string     `json:"fsn"`
	Preferred   *string    `json:"preferred,omitempty"`
	Acceptable  *string    `json:"acceptable,omitempty"`
	Version     *string    `json:"version,omitempty"`
	Hierarchy   *string    `json:"hierarchy,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
