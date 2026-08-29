package datatype

type Annotation struct {
	AuthorReference *Reference `json:"authorReference,omitempty"`
	AuthorString    *string   `json:"authorString,omitempty"`
	Time            *string   `json:"time,omitempty"`
	Text            string    `json:"text,omitempty"`
}

func (a *Annotation) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if a.AuthorReference != nil { m["authorReference"] = a.AuthorReference.ToArray() }
	if a.AuthorString != nil { m["authorString"] = *a.AuthorString }
	if a.Time != nil { m["time"] = *a.Time }
	if a.Text != "" { m["text"] = a.Text }
	return m
}
