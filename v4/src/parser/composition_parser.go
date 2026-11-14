package parser

// CompositionParser extracts fields from a FHIR Composition resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/Composition.json, copied unchanged to
// src/parser/testdata/Composition.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type CompositionParser struct {
	payload map[string]any
}

// NewCompositionParser creates a new CompositionParser instance.
func NewCompositionParser(payload map[string]any) *CompositionParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &CompositionParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *CompositionParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *CompositionParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetStatus returns the status element unchanged.
func (p *CompositionParser) GetStatus() any {
	return p.payload["status"]
}

// GetAttester returns the attester element unchanged.
func (p *CompositionParser) GetAttester() any {
	return p.payload["attester"]
}

// GetType returns the type element unchanged.
func (p *CompositionParser) GetType() any {
	return p.payload["type"]
}

// GetSubject returns the subject element unchanged.
func (p *CompositionParser) GetSubject() any {
	return p.payload["subject"]
}

// GetEncounter returns the encounter element unchanged.
func (p *CompositionParser) GetEncounter() any {
	return p.payload["encounter"]
}

// GetDate returns the date element unchanged.
func (p *CompositionParser) GetDate() any {
	return p.payload["date"]
}

// GetAuthor returns the author element unchanged.
func (p *CompositionParser) GetAuthor() any {
	return p.payload["author"]
}

// GetTitle returns the title element unchanged.
func (p *CompositionParser) GetTitle() any {
	return p.payload["title"]
}

// GetSection returns the section element unchanged.
func (p *CompositionParser) GetSection() any {
	return p.payload["section"]
}
