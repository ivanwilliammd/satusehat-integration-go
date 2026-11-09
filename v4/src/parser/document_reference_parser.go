package parser

// DocumentReferenceParser extracts fields from a FHIR DocumentReference resource map.
type DocumentReferenceParser struct {
	documentReference map[string]any
}

// NewDocumentReferenceParser creates a new DocumentReferenceParser instance.
func NewDocumentReferenceParser(documentReference map[string]any) *DocumentReferenceParser {
	if documentReference == nil {
		documentReference = map[string]any{}
	}
	return &DocumentReferenceParser{documentReference: documentReference}
}

func (p *DocumentReferenceParser) GetIdentifiers() any {
	return p.documentReference["identifier"]
}

func (p *DocumentReferenceParser) GetMasterIdentifier() any {
	return p.documentReference["masterIdentifier"]
}

func (p *DocumentReferenceParser) GetStatus() string {
	return strVal(p.documentReference["status"])
}

func (p *DocumentReferenceParser) GetDocStatus() string {
	return strVal(p.documentReference["docStatus"])
}

func (p *DocumentReferenceParser) GetTypes() any {
	return p.documentReference["type"]
}

func (p *DocumentReferenceParser) GetCategories() any {
	return p.documentReference["category"]
}

func (p *DocumentReferenceParser) GetSubjectReference() string {
	return nestedStr(p.documentReference, "subject", "reference")
}

func (p *DocumentReferenceParser) GetDate() string {
	return strVal(p.documentReference["date"])
}

func (p *DocumentReferenceParser) GetAuthors() any {
	return p.documentReference["author"]
}

func (p *DocumentReferenceParser) GetCustodianReference() string {
	return nestedStr(p.documentReference, "custodian", "reference")
}

func (p *DocumentReferenceParser) GetRelatesTo() any {
	return p.documentReference["relatesTo"]
}

func (p *DocumentReferenceParser) GetDescription() string {
	return strVal(p.documentReference["description"])
}

func (p *DocumentReferenceParser) GetSecurityLabels() any {
	return p.documentReference["securityLabel"]
}

func (p *DocumentReferenceParser) GetContents() any {
	return p.documentReference["content"]
}

func (p *DocumentReferenceParser) GetContext() any {
	return p.documentReference["context"]
}
