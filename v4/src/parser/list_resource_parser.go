package parser

// ListResourceParser extracts fields from a FHIR ListResource resource map.
//
// The element list is derived from PayloadBuilderListResource in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type ListResourceParser struct {
	data map[string]any
}

// NewListResourceParser creates a parser for the given ListResource payload.
func NewListResourceParser(data map[string]any) *ListResourceParser {
	if data == nil {
		data = map[string]any{}
	}
	return &ListResourceParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *ListResourceParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *ListResourceParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *ListResourceParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *ListResourceParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *ListResourceParser) GetStatus() any {
	return p.data["status"]
}

// GetMode returns the mode element.
func (p *ListResourceParser) GetMode() any {
	return p.data["mode"]
}

// GetCode returns the code element.
func (p *ListResourceParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *ListResourceParser) GetSubject() any {
	return p.data["subject"]
}

// GetDate returns the date element.
func (p *ListResourceParser) GetDate() any {
	return p.data["date"]
}

// GetTitle returns the title element.
func (p *ListResourceParser) GetTitle() any {
	return p.data["title"]
}
