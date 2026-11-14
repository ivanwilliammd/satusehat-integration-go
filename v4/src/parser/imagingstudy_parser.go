package parser

// ImagingStudyParser extracts fields from a FHIR ImagingStudy resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/ImagingStudy.json, copied unchanged to
// src/parser/testdata/ImagingStudy.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type ImagingStudyParser struct {
	payload map[string]any
}

// NewImagingStudyParser creates a new ImagingStudyParser instance.
func NewImagingStudyParser(payload map[string]any) *ImagingStudyParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &ImagingStudyParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *ImagingStudyParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *ImagingStudyParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetId returns the id element unchanged.
func (p *ImagingStudyParser) GetId() any {
	return p.payload["id"]
}

// GetIdentifier returns the identifier element unchanged.
func (p *ImagingStudyParser) GetIdentifier() any {
	return p.payload["identifier"]
}

// GetStatus returns the status element unchanged.
func (p *ImagingStudyParser) GetStatus() any {
	return p.payload["status"]
}

// GetModality returns the modality element unchanged.
func (p *ImagingStudyParser) GetModality() any {
	return p.payload["modality"]
}

// GetSubject returns the subject element unchanged.
func (p *ImagingStudyParser) GetSubject() any {
	return p.payload["subject"]
}

// GetStarted returns the started element unchanged.
func (p *ImagingStudyParser) GetStarted() any {
	return p.payload["started"]
}

// GetBasedOn returns the basedOn element unchanged.
func (p *ImagingStudyParser) GetBasedOn() any {
	return p.payload["basedOn"]
}

// GetNumberOfSeries returns the numberOfSeries element unchanged.
func (p *ImagingStudyParser) GetNumberOfSeries() any {
	return p.payload["numberOfSeries"]
}

// GetNumberOfInstances returns the numberOfInstances element unchanged.
func (p *ImagingStudyParser) GetNumberOfInstances() any {
	return p.payload["numberOfInstances"]
}

// GetDescription returns the description element unchanged.
func (p *ImagingStudyParser) GetDescription() any {
	return p.payload["description"]
}

// GetSeries returns the series element unchanged.
func (p *ImagingStudyParser) GetSeries() any {
	return p.payload["series"]
}
