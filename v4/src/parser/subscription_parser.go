package parser

// SubscriptionParser extracts fields from a FHIR Subscription resource map.
//
// The element list is derived from PayloadBuilderSubscription in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type SubscriptionParser struct {
	data map[string]any
}

// NewSubscriptionParser creates a parser for the given Subscription payload.
func NewSubscriptionParser(data map[string]any) *SubscriptionParser {
	if data == nil {
		data = map[string]any{}
	}
	return &SubscriptionParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *SubscriptionParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *SubscriptionParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *SubscriptionParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *SubscriptionParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *SubscriptionParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *SubscriptionParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *SubscriptionParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *SubscriptionParser) GetEncounter() any {
	return p.data["encounter"]
}
