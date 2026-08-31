package terminology

import "time"

type Loinc struct {
	ID                        int64      `json:"id"`
	LOINCNUM                  string    `json:"LOINC_NUM"`
	COMPONENT                 *string   `json:"COMPONENT,omitempty"`
	PROPERTY                  *string   `json:"PROPERTY,omitempty"`
	TIMEASPCT                 *string   `json:"TIME_ASPCT,omitempty"`
	SYSTEM                    *string   `json:"SYSTEM,omitempty"`
	SCALETYP                  *string   `json:"SCALE_TYP,omitempty"`
	METHODTYP                 *string   `json:"METHOD_TYP,omitempty"`
	CLASS                     *string   `json:"CLASS,omitempty"`
	CLASSTYPE                 *string   `json:"CLASSTYPE,omitempty"`
	LONGCOMMONNAME            *string   `json:"LONG_COMMON_NAME,omitempty"`
	SHORTNAME                 *string   `json:"SHORTNAME,omitempty"`
	EXTERNALCOPYRIGHTNOTICE   *string   `json:"EXTERNAL_COPYRIGHT_NOTICE,omitempty"`
	STATUS                    *string   `json:"STATUS,omitempty"`
	VersionFirstReleased       *string   `json:"VersionFirstReleased,omitempty"`
	VersionLastChanged        *string   `json:"VersionLastChanged,omitempty"`
	CreatedAt                 *time.Time `json:"created_at,omitempty"`
	UpdatedAt                 *time.Time `json:"updated_at,omitempty"`
	DeletedAt                 *time.Time `json:"deleted_at,omitempty"`
}
