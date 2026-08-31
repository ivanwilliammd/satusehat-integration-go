package terminology

import "time"

type Cvx struct {
	ID                   int64      `json:"id"`
	CvxCode              string     `json:"cvx_code"`
	CvxShortDescription  string     `json:"cvx_short_description"`
	FullVaccineName      string     `json:"full_vaccine_name"`
	Note                 *string    `json:"note,omitempty"`
	VaccineStatus        string     `json:"vaccine_status"`
	InternalID           int        `json:"internal_id"`
	NonVaccine           bool       `json:"nonvaccine"`
	UpdateDate           *time.Time `json:"update_date,omitempty"`
	CreatedAt            *time.Time `json:"created_at,omitempty"`
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
}
