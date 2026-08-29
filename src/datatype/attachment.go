package datatype

type Attachment struct {
	ContentType *string `json:"contentType,omitempty"`
	Language    *string `json:"language,omitempty"`
	Data        *string `json:"data,omitempty"`
	URL         *string `json:"url,omitempty"`
	Size        *int64  `json:"size,omitempty"`
	Hash        *string `json:"hash,omitempty"`
	Title       *string `json:"title,omitempty"`
	Creation    *string `json:"creation,omitempty"`
}

func (a *Attachment) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if a.ContentType != nil { m["contentType"] = *a.ContentType }
	if a.Language != nil { m["language"] = *a.Language }
	if a.Data != nil { m["data"] = *a.Data }
	if a.URL != nil { m["url"] = *a.URL }
	if a.Size != nil { m["size"] = *a.Size }
	if a.Hash != nil { m["hash"] = *a.Hash }
	if a.Title != nil { m["title"] = *a.Title }
	if a.Creation != nil { m["creation"] = *a.Creation }
	return m
}
