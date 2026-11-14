package terminology

import "time"

// KptlKamar represents kamar procedure/service codes
type KptlKamar struct {
	ID                     int64      `json:"id"`
	NamaTindakanDanLayanan string     `json:"nama_tindakan_dan_layanan"`
	BaseCode               string     `json:"base_code"`
	AllowedModifiers       string     `json:"allowed_modifiers"`
	KodeKptl               string     `json:"kode_kptl"`
	Display                string     `json:"display"`
	CodeSystem             string     `json:"code_system"`
	Version                string     `json:"version"`
	CreatedAt              *time.Time `json:"created_at,omitempty"`
	UpdatedAt              *time.Time `json:"updated_at,omitempty"`
}
