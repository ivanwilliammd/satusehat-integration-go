package parser

// ChargeItemDefinitionParser extracts fields from a FHIR ChargeItemDefinition resource map.
//
// The element list is derived from PayloadBuilderChargeItemDefinition in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type ChargeItemDefinitionParser struct {
	data map[string]any
}

// NewChargeItemDefinitionParser creates a parser for the given ChargeItemDefinition payload.
func NewChargeItemDefinitionParser(data map[string]any) *ChargeItemDefinitionParser {
	if data == nil {
		data = map[string]any{}
	}
	return &ChargeItemDefinitionParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *ChargeItemDefinitionParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *ChargeItemDefinitionParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *ChargeItemDefinitionParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *ChargeItemDefinitionParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetUrl returns the url element.
func (p *ChargeItemDefinitionParser) GetUrl() any {
	return p.data["url"]
}

// GetVersion returns the version element.
func (p *ChargeItemDefinitionParser) GetVersion() any {
	return p.data["version"]
}

// GetName returns the name element.
func (p *ChargeItemDefinitionParser) GetName() any {
	return p.data["name"]
}

// GetTitle returns the title element.
func (p *ChargeItemDefinitionParser) GetTitle() any {
	return p.data["title"]
}

// GetStatus returns the status element.
func (p *ChargeItemDefinitionParser) GetStatus() any {
	return p.data["status"]
}

// GetDate returns the date element.
func (p *ChargeItemDefinitionParser) GetDate() any {
	return p.data["date"]
}

// GetPublisher returns the publisher element.
func (p *ChargeItemDefinitionParser) GetPublisher() any {
	return p.data["publisher"]
}

// GetDescription returns the description element.
func (p *ChargeItemDefinitionParser) GetDescription() any {
	return p.data["description"]
}

// GetCode returns the code element.
func (p *ChargeItemDefinitionParser) GetCode() any {
	return p.data["code"]
}

// GetInstance returns the instance element.
func (p *ChargeItemDefinitionParser) GetInstance() any {
	return p.data["instance"]
}

// GetProperty returns the property element.
func (p *ChargeItemDefinitionParser) GetProperty() any {
	return p.data["property"]
}

// GetApplies returns the applies element.
func (p *ChargeItemDefinitionParser) GetApplies() any {
	return p.data["applies"]
}
