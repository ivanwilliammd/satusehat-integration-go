package terminology

import "time"

// KptlModifier represents procedure modifier categories
type KptlModifier struct {
	ID               int64      `json:"id"`
	KategoriKelompok string     `json:"kategori_kelompok"`
	Item             string     `json:"item"`
	ModifierCode     string     `json:"modifier_code"`
	CreatedAt        *time.Time `json:"created_at,omitempty"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
}
