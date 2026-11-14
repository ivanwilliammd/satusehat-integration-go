package parser

// AuditEventParser extracts fields from a FHIR AuditEvent resource map.
//
// The element list is derived from PayloadBuilderAuditEvent in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type AuditEventParser struct {
	data map[string]any
}

// NewAuditEventParser creates a parser for the given AuditEvent payload.
func NewAuditEventParser(data map[string]any) *AuditEventParser {
	if data == nil {
		data = map[string]any{}
	}
	return &AuditEventParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *AuditEventParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *AuditEventParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *AuditEventParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *AuditEventParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *AuditEventParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *AuditEventParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *AuditEventParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *AuditEventParser) GetEncounter() any {
	return p.data["encounter"]
}
