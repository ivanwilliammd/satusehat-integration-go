package datatype

type RelatedArtifact struct {
	Type     *string     `json:"type,omitempty"`
	Label    *string     `json:"label,omitempty"`
	Display  *string     `json:"display,omitempty"`
	Citation *string     `json:"citation,omitempty"`
	URL      *string     `json:"url,omitempty"`
	Document *Attachment `json:"document,omitempty"`
	Resource *string     `json:"resource,omitempty"`
}

func (r *RelatedArtifact) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if r.Type != nil { m["type"] = *r.Type }
	if r.Label != nil { m["label"] = *r.Label }
	if r.Display != nil { m["display"] = *r.Display }
	if r.Citation != nil { m["citation"] = *r.Citation }
	if r.URL != nil { m["url"] = *r.URL }
	if r.Document != nil { m["document"] = r.Document.ToArray() }
	if r.Resource != nil { m["resource"] = *r.Resource }
	return m
}

func (r *RelatedArtifact) SetDocument(doc *Attachment) *RelatedArtifact {
	r.Document = doc
	return r
}
