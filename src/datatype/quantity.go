package datatype

type Quantity struct {
    Value      *float64 `json:"value,omitempty"`
    Comparator *string  `json:"comparator,omitempty"`
    Unit       *string  `json:"unit,omitempty"`
    System     *string  `json:"system,omitempty"`
    Code       *string  `json:"code,omitempty"`
}
