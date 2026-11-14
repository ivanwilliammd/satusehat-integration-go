package parser

// BundleParser extracts fields from a FHIR Bundle resource map.
//
// The element list is derived from PayloadBuilderBundle in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type BundleParser struct {
	data map[string]any
}

// NewBundleParser creates a parser for the given Bundle payload.
func NewBundleParser(data map[string]any) *BundleParser {
	if data == nil {
		data = map[string]any{}
	}
	return &BundleParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *BundleParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *BundleParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *BundleParser) GetId() any {
	return p.data["id"]
}

// GetType returns the type element.
func (p *BundleParser) GetType() any {
	return p.data["type"]
}

// GetTimestamp returns the timestamp element.
func (p *BundleParser) GetTimestamp() any {
	return p.data["timestamp"]
}

// GetTotal returns the total element.
func (p *BundleParser) GetTotal() any {
	return p.data["total"]
}

// GetMeta returns the meta element.
func (p *BundleParser) GetMeta() any {
	return p.data["meta"]
}

// GetLink returns the link element.
func (p *BundleParser) GetLink() any {
	return p.data["link"]
}

// GetEntry returns the entry element.
func (p *BundleParser) GetEntry() any {
	return p.data["entry"]
}

// GetSignature returns the signature element.
func (p *BundleParser) GetSignature() any {
	return p.data["signature"]
}
