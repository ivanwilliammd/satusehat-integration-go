package datatype

type Narrative struct {
	Status *string `json:"status,omitempty"`
	Div    *string `json:"div,omitempty"`
}

func (n *Narrative) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if n.Status != nil { m["status"] = *n.Status }
	if n.Div != nil { m["div"] = *n.Div }
	return m
}
