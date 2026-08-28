package datatype

type Period struct {
    Start *string `json:"start,omitempty"`
    End   *string `json:"end,omitempty"`
}

type Address struct {
    Use        *string  `json:"use,omitempty"`
    Type       *string  `json:"type,omitempty"`
    Text       *string  `json:"text,omitempty"`
    Line       []string `json:"line,omitempty"`
    City       *string  `json:"city,omitempty"`
    District   *string  `json:"district,omitempty"`
    State      *string  `json:"state,omitempty"`
    PostalCode *string  `json:"postalCode,omitempty"`
    Country    *string  `json:"country,omitempty"`
    Period     *Period  `json:"period,omitempty"`
}
