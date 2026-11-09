package parser

// ServiceRequestParser extracts fields from a FHIR ServiceRequest resource map.
type ServiceRequestParser struct {
	serviceRequest map[string]any
}

// NewServiceRequestParser creates a new ServiceRequestParser instance.
func NewServiceRequestParser(serviceRequest map[string]any) *ServiceRequestParser {
	if serviceRequest == nil {
		serviceRequest = map[string]any{}
	}
	return &ServiceRequestParser{serviceRequest: serviceRequest}
}

func (p *ServiceRequestParser) GetIdentifiers() any {
	return p.serviceRequest["identifier"]
}

func (p *ServiceRequestParser) GetRequisition() any {
	return p.serviceRequest["requisition"]
}

func (p *ServiceRequestParser) GetStatus() string {
	return strVal(p.serviceRequest["status"])
}

func (p *ServiceRequestParser) GetIntent() string {
	return strVal(p.serviceRequest["intent"])
}

func (p *ServiceRequestParser) GetCategories() any {
	return p.serviceRequest["category"]
}

func (p *ServiceRequestParser) GetPriority() string {
	return strVal(p.serviceRequest["priority"])
}

func (p *ServiceRequestParser) GetDoNotPerform() any {
	return p.serviceRequest["doNotPerform"]
}

func (p *ServiceRequestParser) GetCode() any {
	return p.serviceRequest["code"]
}

func (p *ServiceRequestParser) GetQuantityQuantity() any {
	return p.serviceRequest["quantityQuantity"]
}

func (p *ServiceRequestParser) GetSubjectReference() string {
	return nestedStr(p.serviceRequest, "subject", "reference")
}

func (p *ServiceRequestParser) GetSubjectDisplay() string {
	return nestedStr(p.serviceRequest, "subject", "display")
}

func (p *ServiceRequestParser) GetEncounterReference() string {
	return removePrefix(nestedStr(p.serviceRequest, "encounter", "reference"), "Encounter/")
}

func (p *ServiceRequestParser) GetEncounterDisplay() string {
	return nestedStr(p.serviceRequest, "encounter", "display")
}

func (p *ServiceRequestParser) GetOccurrenceDateTime() string {
	return strVal(p.serviceRequest["occurrenceDateTime"])
}

func (p *ServiceRequestParser) GetAuthoredOn() string {
	return strVal(p.serviceRequest["authoredOn"])
}

func (p *ServiceRequestParser) GetRequesterReference() string {
	return nestedStr(p.serviceRequest, "requester", "reference")
}

func (p *ServiceRequestParser) GetRequesterDisplay() string {
	return nestedStr(p.serviceRequest, "requester", "display")
}

func (p *ServiceRequestParser) GetPerformers() any {
	return p.serviceRequest["performer"]
}

func (p *ServiceRequestParser) GetReasonCodes() any {
	return p.serviceRequest["reasonCode"]
}

func (p *ServiceRequestParser) GetSupportingInfo() any {
	return p.serviceRequest["supportingInfo"]
}

func (p *ServiceRequestParser) GetSpecimens() any {
	return p.serviceRequest["specimen"]
}

func (p *ServiceRequestParser) GetNotes() any {
	return p.serviceRequest["note"]
}

func (p *ServiceRequestParser) GetPatientInstruction() string {
	return strVal(p.serviceRequest["patientInstruction"])
}

func (p *ServiceRequestParser) GetRelevantHistory() any {
	return p.serviceRequest["relevantHistory"]
}
