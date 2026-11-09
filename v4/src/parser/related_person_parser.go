package parser

// RelatedPersonParser extracts fields from a FHIR RelatedPerson resource map.
type RelatedPersonParser struct {
	relatedPerson map[string]any
}

// NewRelatedPersonParser creates a new RelatedPersonParser instance.
func NewRelatedPersonParser(relatedPerson map[string]any) *RelatedPersonParser {
	if relatedPerson == nil {
		relatedPerson = map[string]any{}
	}
	return &RelatedPersonParser{relatedPerson: relatedPerson}
}

func (p *RelatedPersonParser) GetIdentifiers() any {
	return p.relatedPerson["identifier"]
}

func (p *RelatedPersonParser) GetActive() any {
	return p.relatedPerson["active"]
}

func (p *RelatedPersonParser) GetPatientReference() string {
	return nestedStr(p.relatedPerson, "patient", "reference")
}

func (p *RelatedPersonParser) GetRelationships() any {
	return p.relatedPerson["relationship"]
}

func (p *RelatedPersonParser) GetNames() any {
	return p.relatedPerson["name"]
}

func (p *RelatedPersonParser) GetTelecoms() any {
	return p.relatedPerson["telecom"]
}

func (p *RelatedPersonParser) GetGender() string {
	return strVal(p.relatedPerson["gender"])
}

func (p *RelatedPersonParser) GetBirthDate() string {
	return strVal(p.relatedPerson["birthDate"])
}

func (p *RelatedPersonParser) GetAddresses() any {
	return p.relatedPerson["address"]
}

func (p *RelatedPersonParser) GetCommunications() any {
	return p.relatedPerson["communication"]
}
