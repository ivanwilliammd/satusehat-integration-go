package parser

// CatalogEntryParser extracts fields from a FHIR CatalogEntry resource map.
//
// The element list is derived from PayloadBuilderCatalogEntry in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type CatalogEntryParser struct {
	data map[string]any
}

// NewCatalogEntryParser creates a parser for the given CatalogEntry payload.
func NewCatalogEntryParser(data map[string]any) *CatalogEntryParser {
	if data == nil {
		data = map[string]any{}
	}
	return &CatalogEntryParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *CatalogEntryParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *CatalogEntryParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *CatalogEntryParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *CatalogEntryParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *CatalogEntryParser) GetStatus() any {
	return p.data["status"]
}

// GetType returns the type element.
func (p *CatalogEntryParser) GetType() any {
	return p.data["type"]
}

// GetReference returns the reference element.
func (p *CatalogEntryParser) GetReference() any {
	return p.data["reference"]
}

// GetValidityPeriod returns the validityPeriod element.
func (p *CatalogEntryParser) GetValidityPeriod() any {
	return p.data["validityPeriod"]
}
