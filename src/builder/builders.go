package builder

type BaseBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewBaseBuilder(rt string) *BaseBuilder {
	return &BaseBuilder{ResourceType: rt, Data: make(map[string]interface{})}
}

func (b *BaseBuilder) GetResourceType() string       { return b.ResourceType }

func (b *BaseBuilder) ToJSON() map[string]interface{} { return b.Data }

func (b *BaseBuilder) SetID(id string) *BaseBuilder   { b.Data["id"] = id; return b }

func (b *BaseBuilder) AddMeta(meta map[string]interface{}) *BaseBuilder { b.Data["meta"] = meta; return b }

func (b *BaseBuilder) AddExtension(ext string, val interface{}) *BaseBuilder { b.Data[ext] = val; return b }


// Account Builder
type AccountBuilder struct{ BaseBuilder }

func NewAccountBuilder() *AccountBuilder {
	b := &AccountBuilder{}

	b.ResourceType = "Account"

	b.Data["resourceType"] = "Account"

	return b
}

// AllergyIntolerance Builder
type AllergyIntoleranceBuilder struct{ BaseBuilder }

func NewAllergyIntoleranceBuilder() *AllergyIntoleranceBuilder {
	b := &AllergyIntoleranceBuilder{}

	b.ResourceType = "AllergyIntolerance"

	b.Data["resourceType"] = "AllergyIntolerance"

	return b
}

// Bundle Builder
type BundleBuilder struct{ BaseBuilder }

func NewBundleBuilder() *BundleBuilder {
	b := &BundleBuilder{}

	b.ResourceType = "Bundle"

	b.Data["resourceType"] = "Bundle"

	return b
}

// CarePlan Builder
type CarePlanBuilder struct{ BaseBuilder }

func NewCarePlanBuilder() *CarePlanBuilder {
	b := &CarePlanBuilder{}

	b.ResourceType = "CarePlan"

	b.Data["resourceType"] = "CarePlan"

	return b
}

// ChargeItem Builder
type ChargeItemBuilder struct{ BaseBuilder }

func NewChargeItemBuilder() *ChargeItemBuilder {
	b := &ChargeItemBuilder{}

	b.ResourceType = "ChargeItem"

	b.Data["resourceType"] = "ChargeItem"

	return b
}

// ChargeItemDefinition Builder
type ChargeItemDefinitionBuilder struct{ BaseBuilder }

func NewChargeItemDefinitionBuilder() *ChargeItemDefinitionBuilder {
	b := &ChargeItemDefinitionBuilder{}

	b.ResourceType = "ChargeItemDefinition"

	b.Data["resourceType"] = "ChargeItemDefinition"

	return b
}

// ChargeItemResponse Builder
type ChargeItemResponseBuilder struct{ BaseBuilder }

func NewChargeItemResponseBuilder() *ChargeItemResponseBuilder {
	b := &ChargeItemResponseBuilder{}

	b.ResourceType = "ChargeItemResponse"

	b.Data["resourceType"] = "ChargeItemResponse"

	return b
}

// Claim Builder
type ClaimBuilder struct{ BaseBuilder }

func NewClaimBuilder() *ClaimBuilder {
	b := &ClaimBuilder{}

	b.ResourceType = "Claim"

	b.Data["resourceType"] = "Claim"

	return b
}

// ClaimResponse Builder
type ClaimResponseBuilder struct{ BaseBuilder }

func NewClaimResponseBuilder() *ClaimResponseBuilder {
	b := &ClaimResponseBuilder{}

	b.ResourceType = "ClaimResponse"

	b.Data["resourceType"] = "ClaimResponse"

	return b
}

// ClinicalImpression Builder
type ClinicalImpressionBuilder struct{ BaseBuilder }

func NewClinicalImpressionBuilder() *ClinicalImpressionBuilder {
	b := &ClinicalImpressionBuilder{}

	b.ResourceType = "ClinicalImpression"

	b.Data["resourceType"] = "ClinicalImpression"

	return b
}

// Composition Builder
type CompositionBuilder struct{ BaseBuilder }

func NewCompositionBuilder() *CompositionBuilder {
	b := &CompositionBuilder{}

	b.ResourceType = "Composition"

	b.Data["resourceType"] = "Composition"

	return b
}

// Condition Builder
type ConditionBuilder struct{ BaseBuilder }

func NewConditionBuilder() *ConditionBuilder {
	b := &ConditionBuilder{}

	b.ResourceType = "Condition"

	b.Data["resourceType"] = "Condition"

	return b
}

// Coverage Builder
type CoverageBuilder struct{ BaseBuilder }

func NewCoverageBuilder() *CoverageBuilder {
	b := &CoverageBuilder{}

	b.ResourceType = "Coverage"

	b.Data["resourceType"] = "Coverage"

	return b
}

// CoverageEligibilityRequest Builder
type CoverageEligibilityRequestBuilder struct{ BaseBuilder }

func NewCoverageEligibilityRequestBuilder() *CoverageEligibilityRequestBuilder {
	b := &CoverageEligibilityRequestBuilder{}

	b.ResourceType = "CoverageEligibilityRequest"

	b.Data["resourceType"] = "CoverageEligibilityRequest"

	return b
}

// CoverageEligibilityResponse Builder
type CoverageEligibilityResponseBuilder struct{ BaseBuilder }

func NewCoverageEligibilityResponseBuilder() *CoverageEligibilityResponseBuilder {
	b := &CoverageEligibilityResponseBuilder{}

	b.ResourceType = "CoverageEligibilityResponse"

	b.Data["resourceType"] = "CoverageEligibilityResponse"

	return b
}

// Device Builder
type DeviceBuilder struct{ BaseBuilder }

func NewDeviceBuilder() *DeviceBuilder {
	b := &DeviceBuilder{}

	b.ResourceType = "Device"

	b.Data["resourceType"] = "Device"

	return b
}

// DiagnosticReport Builder
type DiagnosticReportBuilder struct{ BaseBuilder }

func NewDiagnosticReportBuilder() *DiagnosticReportBuilder {
	b := &DiagnosticReportBuilder{}

	b.ResourceType = "DiagnosticReport"

	b.Data["resourceType"] = "DiagnosticReport"

	return b
}

// DocumentReference Builder
type DocumentReferenceBuilder struct{ BaseBuilder }

func NewDocumentReferenceBuilder() *DocumentReferenceBuilder {
	b := &DocumentReferenceBuilder{}

	b.ResourceType = "DocumentReference"

	b.Data["resourceType"] = "DocumentReference"

	return b
}

// Encounter Builder
type EncounterBuilder struct{ BaseBuilder }

func NewEncounterBuilder() *EncounterBuilder {
	b := &EncounterBuilder{}

	b.ResourceType = "Encounter"

	b.Data["resourceType"] = "Encounter"

	return b
}

// EpisodeOfCare Builder
type EpisodeOfCareBuilder struct{ BaseBuilder }

func NewEpisodeOfCareBuilder() *EpisodeOfCareBuilder {
	b := &EpisodeOfCareBuilder{}

	b.ResourceType = "EpisodeOfCare"

	b.Data["resourceType"] = "EpisodeOfCare"

	return b
}

// FamilyMemberHistory Builder
type FamilyMemberHistoryBuilder struct{ BaseBuilder }

func NewFamilyMemberHistoryBuilder() *FamilyMemberHistoryBuilder {
	b := &FamilyMemberHistoryBuilder{}

	b.ResourceType = "FamilyMemberHistory"

	b.Data["resourceType"] = "FamilyMemberHistory"

	return b
}

// GenomicStudy Builder
type GenomicStudyBuilder struct{ BaseBuilder }

func NewGenomicStudyBuilder() *GenomicStudyBuilder {
	b := &GenomicStudyBuilder{}

	b.ResourceType = "GenomicStudy"

	b.Data["resourceType"] = "GenomicStudy"

	return b
}

// Goal Builder
type GoalBuilder struct{ BaseBuilder }

func NewGoalBuilder() *GoalBuilder {
	b := &GoalBuilder{}

	b.ResourceType = "Goal"

	b.Data["resourceType"] = "Goal"

	return b
}

// Group Builder
type GroupBuilder struct{ BaseBuilder }

func NewGroupBuilder() *GroupBuilder {
	b := &GroupBuilder{}

	b.ResourceType = "Group"

	b.Data["resourceType"] = "Group"

	return b
}

// ImagingStudy Builder
type ImagingStudyBuilder struct{ BaseBuilder }

func NewImagingStudyBuilder() *ImagingStudyBuilder {
	b := &ImagingStudyBuilder{}

	b.ResourceType = "ImagingStudy"

	b.Data["resourceType"] = "ImagingStudy"

	return b
}

// Immunization Builder
type ImmunizationBuilder struct{ BaseBuilder }

func NewImmunizationBuilder() *ImmunizationBuilder {
	b := &ImmunizationBuilder{}

	b.ResourceType = "Immunization"

	b.Data["resourceType"] = "Immunization"

	return b
}

// Invoice Builder
type InvoiceBuilder struct{ BaseBuilder }

func NewInvoiceBuilder() *InvoiceBuilder {
	b := &InvoiceBuilder{}

	b.ResourceType = "Invoice"

	b.Data["resourceType"] = "Invoice"

	return b
}

// Location Builder
type LocationBuilder struct{ BaseBuilder }

func NewLocationBuilder() *LocationBuilder {
	b := &LocationBuilder{}

	b.ResourceType = "Location"

	b.Data["resourceType"] = "Location"

	return b
}

// Medication Builder
type MedicationBuilder struct{ BaseBuilder }

func NewMedicationBuilder() *MedicationBuilder {
	b := &MedicationBuilder{}

	b.ResourceType = "Medication"

	b.Data["resourceType"] = "Medication"

	return b
}

// MedicationAdministration Builder
type MedicationAdministrationBuilder struct{ BaseBuilder }

func NewMedicationAdministrationBuilder() *MedicationAdministrationBuilder {
	b := &MedicationAdministrationBuilder{}

	b.ResourceType = "MedicationAdministration"

	b.Data["resourceType"] = "MedicationAdministration"

	return b
}

// MedicationDispense Builder
type MedicationDispenseBuilder struct{ BaseBuilder }

func NewMedicationDispenseBuilder() *MedicationDispenseBuilder {
	b := &MedicationDispenseBuilder{}

	b.ResourceType = "MedicationDispense"

	b.Data["resourceType"] = "MedicationDispense"

	return b
}

// MedicationRequest Builder
type MedicationRequestBuilder struct{ BaseBuilder }

func NewMedicationRequestBuilder() *MedicationRequestBuilder {
	b := &MedicationRequestBuilder{}

	b.ResourceType = "MedicationRequest"

	b.Data["resourceType"] = "MedicationRequest"

	return b
}

// MedicationStatement Builder
type MedicationStatementBuilder struct{ BaseBuilder }

func NewMedicationStatementBuilder() *MedicationStatementBuilder {
	b := &MedicationStatementBuilder{}

	b.ResourceType = "MedicationStatement"

	b.Data["resourceType"] = "MedicationStatement"

	return b
}

// MolecularSequence Builder
type MolecularSequenceBuilder struct{ BaseBuilder }

func NewMolecularSequenceBuilder() *MolecularSequenceBuilder {
	b := &MolecularSequenceBuilder{}

	b.ResourceType = "MolecularSequence"

	b.Data["resourceType"] = "MolecularSequence"

	return b
}

// NutritionOrder Builder
type NutritionOrderBuilder struct{ BaseBuilder }

func NewNutritionOrderBuilder() *NutritionOrderBuilder {
	b := &NutritionOrderBuilder{}

	b.ResourceType = "NutritionOrder"

	b.Data["resourceType"] = "NutritionOrder"

	return b
}

// Observation Builder
type ObservationBuilder struct{ BaseBuilder }

func NewObservationBuilder() *ObservationBuilder {
	b := &ObservationBuilder{}

	b.ResourceType = "Observation"

	b.Data["resourceType"] = "Observation"

	return b
}

// Organization Builder
type OrganizationBuilder struct{ BaseBuilder }

func NewOrganizationBuilder() *OrganizationBuilder {
	b := &OrganizationBuilder{}

	b.ResourceType = "Organization"

	b.Data["resourceType"] = "Organization"

	return b
}

// Patient Builder
type PatientBuilder struct{ BaseBuilder }

func NewPatientBuilder() *PatientBuilder {
	b := &PatientBuilder{}

	b.ResourceType = "Patient"

	b.Data["resourceType"] = "Patient"

	return b
}

// PaymentNotice Builder
type PaymentNoticeBuilder struct{ BaseBuilder }

func NewPaymentNoticeBuilder() *PaymentNoticeBuilder {
	b := &PaymentNoticeBuilder{}

	b.ResourceType = "PaymentNotice"

	b.Data["resourceType"] = "PaymentNotice"

	return b
}

// PaymentReconciliation Builder
type PaymentReconciliationBuilder struct{ BaseBuilder }

func NewPaymentReconciliationBuilder() *PaymentReconciliationBuilder {
	b := &PaymentReconciliationBuilder{}

	b.ResourceType = "PaymentReconciliation"

	b.Data["resourceType"] = "PaymentReconciliation"

	return b
}

// Practitioner Builder
type PractitionerBuilder struct{ BaseBuilder }

func NewPractitionerBuilder() *PractitionerBuilder {
	b := &PractitionerBuilder{}

	b.ResourceType = "Practitioner"

	b.Data["resourceType"] = "Practitioner"

	return b
}

// PractitionerRole Builder
type PractitionerRoleBuilder struct{ BaseBuilder }

func NewPractitionerRoleBuilder() *PractitionerRoleBuilder {
	b := &PractitionerRoleBuilder{}

	b.ResourceType = "PractitionerRole"

	b.Data["resourceType"] = "PractitionerRole"

	return b
}

// Procedure Builder
type ProcedureBuilder struct{ BaseBuilder }

func NewProcedureBuilder() *ProcedureBuilder {
	b := &ProcedureBuilder{}

	b.ResourceType = "Procedure"

	b.Data["resourceType"] = "Procedure"

	return b
}

// QuestionnaireResponse Builder
type QuestionnaireResponseBuilder struct{ BaseBuilder }

func NewQuestionnaireResponseBuilder() *QuestionnaireResponseBuilder {
	b := &QuestionnaireResponseBuilder{}

	b.ResourceType = "QuestionnaireResponse"

	b.Data["resourceType"] = "QuestionnaireResponse"

	return b
}

// RelatedPerson Builder
type RelatedPersonBuilder struct{ BaseBuilder }

func NewRelatedPersonBuilder() *RelatedPersonBuilder {
	b := &RelatedPersonBuilder{}

	b.ResourceType = "RelatedPerson"

	b.Data["resourceType"] = "RelatedPerson"

	return b
}

// RiskAssessment Builder
type RiskAssessmentBuilder struct{ BaseBuilder }

func NewRiskAssessmentBuilder() *RiskAssessmentBuilder {
	b := &RiskAssessmentBuilder{}

	b.ResourceType = "RiskAssessment"

	b.Data["resourceType"] = "RiskAssessment"

	return b
}

// ServiceRequest Builder
type ServiceRequestBuilder struct{ BaseBuilder }

func NewServiceRequestBuilder() *ServiceRequestBuilder {
	b := &ServiceRequestBuilder{}

	b.ResourceType = "ServiceRequest"

	b.Data["resourceType"] = "ServiceRequest"

	return b
}

// Specimen Builder
type SpecimenBuilder struct{ BaseBuilder }

func NewSpecimenBuilder() *SpecimenBuilder {
	b := &SpecimenBuilder{}

	b.ResourceType = "Specimen"

	b.Data["resourceType"] = "Specimen"

	return b
}

// Substance Builder
type SubstanceBuilder struct{ BaseBuilder }

func NewSubstanceBuilder() *SubstanceBuilder {
	b := &SubstanceBuilder{}

	b.ResourceType = "Substance"

	b.Data["resourceType"] = "Substance"

	return b
}

// Task Builder
type TaskBuilder struct{ BaseBuilder }

func NewTaskBuilder() *TaskBuilder {
	b := &TaskBuilder{}

	b.ResourceType = "Task"

	b.Data["resourceType"] = "Task"

	return b
}
