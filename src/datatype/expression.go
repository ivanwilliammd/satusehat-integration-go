package datatype

type Expression struct {
	Description *string `json:"description,omitempty"`
	Language    *string `json:"language,omitempty"`
	Expression  *string `json:"expression,omitempty"`
	Reference   *string `json:"reference,omitempty"`
}

func (e *Expression) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if e.Description != nil { m["description"] = *e.Description }
	if e.Language != nil { m["language"] = *e.Language }
	if e.Expression != nil { m["expression"] = *e.Expression }
	if e.Reference != nil { m["reference"] = *e.Reference }
	return m
}
