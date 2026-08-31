package terminology

import "time"

type KodeWilayahIndonesia struct {
	ID            int64      `json:"id"`
	KodeWilayah   string     `json:"kode_wilayah"`
	NamaWilayah   string     `json:"nama_wilayah"`
	Level         *int       `json:"level,omitempty"`
	Parent        *string    `json:"parent,omitempty"`
	State         *string    `json:"state,omitempty"`
	Active        bool       `json:"active"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}
