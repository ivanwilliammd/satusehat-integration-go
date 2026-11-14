package parser

// SubstanceReferenceInformationParser extracts fields from a FHIR SubstanceReferenceInformation resource map.
//
// The element list is derived from PayloadBuilderSubstanceReferenceInformation in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type SubstanceReferenceInformationParser struct {
	data map[string]any
}

// NewSubstanceReferenceInformationParser creates a parser for the given SubstanceReferenceInformation payload.
func NewSubstanceReferenceInformationParser(data map[string]any) *SubstanceReferenceInformationParser {
	if data == nil {
		data = map[string]any{}
	}
	return &SubstanceReferenceInformationParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *SubstanceReferenceInformationParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *SubstanceReferenceInformationParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *SubstanceReferenceInformationParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *SubstanceReferenceInformationParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *SubstanceReferenceInformationParser) GetStatus() any {
	return p.data["status"]
}

// GetComment returns the comment element.
func (p *SubstanceReferenceInformationParser) GetComment() any {
	return p.data["comment"]
}

// GetGene returns the gene element.
func (p *SubstanceReferenceInformationParser) GetGene() any {
	return p.data["gene"]
}

// GetGeneElement returns the geneElement element.
func (p *SubstanceReferenceInformationParser) GetGeneElement() any {
	return p.data["geneElement"]
}
