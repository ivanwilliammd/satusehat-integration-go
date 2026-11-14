package parser_test

// These parsers were derived from the PHP builders because their resources have no official
// SATUSEHAT example. Each test feeds a payload built from the same element list and reads it back,
// so a getter that points at the wrong key fails immediately.

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestActivityDefinitionParser(t *testing.T) {
	p := parser.NewActivityDefinitionParser(map[string]any{
		"resourceType": "ActivityDefinition",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"description":  "nilai-description",
		"kind":         "nilai-kind",
		"code":         "nilai-code",
		"author":       "nilai-author",
		"timing":       "nilai-timing",
	})

	if got := p.GetResourceType(); got != "ActivityDefinition" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetDescription(); got != "nilai-description" {
		t.Errorf("description: dapat %v", got)
	}
	if got := p.GetKind(); got != "nilai-kind" {
		t.Errorf("kind: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetAuthor(); got != "nilai-author" {
		t.Errorf("author: dapat %v", got)
	}
	if got := p.GetTiming(); got != "nilai-timing" {
		t.Errorf("timing: dapat %v", got)
	}
}
func TestAdverseEventParser(t *testing.T) {
	p := parser.NewAdverseEventParser(map[string]any{
		"resourceType": "AdverseEvent",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "AdverseEvent" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestAppointmentParser(t *testing.T) {
	p := parser.NewAppointmentParser(map[string]any{
		"resourceType": "Appointment",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "Appointment" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestAppointmentResponseParser(t *testing.T) {
	p := parser.NewAppointmentResponseParser(map[string]any{
		"resourceType": "AppointmentResponse",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "AppointmentResponse" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestAuditEventParser(t *testing.T) {
	p := parser.NewAuditEventParser(map[string]any{
		"resourceType": "AuditEvent",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "AuditEvent" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestBasicParser(t *testing.T) {
	p := parser.NewBasicParser(map[string]any{
		"resourceType": "Basic",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "Basic" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestBinaryParser(t *testing.T) {
	p := parser.NewBinaryParser(map[string]any{
		"resourceType": "Binary",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "Binary" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestBiologicallyDerivedProductParser(t *testing.T) {
	p := parser.NewBiologicallyDerivedProductParser(map[string]any{
		"resourceType": "BiologicallyDerivedProduct",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "BiologicallyDerivedProduct" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestBodyStructureParser(t *testing.T) {
	p := parser.NewBodyStructureParser(map[string]any{
		"resourceType": "BodyStructure",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "BodyStructure" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestBundleParser(t *testing.T) {
	p := parser.NewBundleParser(map[string]any{
		"resourceType": "Bundle",
		"id":           "nilai-id",
		"type":         "nilai-type",
		"timestamp":    "nilai-timestamp",
		"total":        "nilai-total",
		"meta":         "nilai-meta",
		"link":         "nilai-link",
		"entry":        "nilai-entry",
		"signature":    "nilai-signature",
	})

	if got := p.GetResourceType(); got != "Bundle" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetType(); got != "nilai-type" {
		t.Errorf("type: dapat %v", got)
	}
	if got := p.GetTimestamp(); got != "nilai-timestamp" {
		t.Errorf("timestamp: dapat %v", got)
	}
	if got := p.GetTotal(); got != "nilai-total" {
		t.Errorf("total: dapat %v", got)
	}
	if got := p.GetMeta(); got != "nilai-meta" {
		t.Errorf("meta: dapat %v", got)
	}
	if got := p.GetLink(); got != "nilai-link" {
		t.Errorf("link: dapat %v", got)
	}
	if got := p.GetEntry(); got != "nilai-entry" {
		t.Errorf("entry: dapat %v", got)
	}
	if got := p.GetSignature(); got != "nilai-signature" {
		t.Errorf("signature: dapat %v", got)
	}
}
func TestCapabilityStatementParser(t *testing.T) {
	p := parser.NewCapabilityStatementParser(map[string]any{
		"resourceType": "CapabilityStatement",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"date":         "nilai-date",
		"kind":         "nilai-kind",
		"fhirVersion":  "nilai-fhirVersion",
		"format":       "nilai-format",
		"rest":         "nilai-rest",
	})

	if got := p.GetResourceType(); got != "CapabilityStatement" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetDate(); got != "nilai-date" {
		t.Errorf("date: dapat %v", got)
	}
	if got := p.GetKind(); got != "nilai-kind" {
		t.Errorf("kind: dapat %v", got)
	}
	if got := p.GetFhirVersion(); got != "nilai-fhirVersion" {
		t.Errorf("fhirVersion: dapat %v", got)
	}
	if got := p.GetFormat(); got != "nilai-format" {
		t.Errorf("format: dapat %v", got)
	}
	if got := p.GetRest(); got != "nilai-rest" {
		t.Errorf("rest: dapat %v", got)
	}
}
func TestCareTeamParser(t *testing.T) {
	p := parser.NewCareTeamParser(map[string]any{
		"resourceType": "CareTeam",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "CareTeam" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestCatalogEntryParser(t *testing.T) {
	p := parser.NewCatalogEntryParser(map[string]any{
		"resourceType":   "CatalogEntry",
		"id":             "nilai-id",
		"identifier":     "nilai-identifier",
		"status":         "nilai-status",
		"type":           "nilai-type",
		"reference":      "nilai-reference",
		"validityPeriod": "nilai-validityPeriod",
	})

	if got := p.GetResourceType(); got != "CatalogEntry" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetType(); got != "nilai-type" {
		t.Errorf("type: dapat %v", got)
	}
	if got := p.GetReference(); got != "nilai-reference" {
		t.Errorf("reference: dapat %v", got)
	}
	if got := p.GetValidityPeriod(); got != "nilai-validityPeriod" {
		t.Errorf("validityPeriod: dapat %v", got)
	}
}
func TestChargeItemDefinitionParser(t *testing.T) {
	p := parser.NewChargeItemDefinitionParser(map[string]any{
		"resourceType": "ChargeItemDefinition",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"url":          "nilai-url",
		"version":      "nilai-version",
		"name":         "nilai-name",
		"title":        "nilai-title",
		"status":       "nilai-status",
		"date":         "nilai-date",
		"publisher":    "nilai-publisher",
		"description":  "nilai-description",
		"code":         "nilai-code",
		"instance":     "nilai-instance",
		"property":     "nilai-property",
		"applies":      "nilai-applies",
	})

	if got := p.GetResourceType(); got != "ChargeItemDefinition" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetUrl(); got != "nilai-url" {
		t.Errorf("url: dapat %v", got)
	}
	if got := p.GetVersion(); got != "nilai-version" {
		t.Errorf("version: dapat %v", got)
	}
	if got := p.GetName(); got != "nilai-name" {
		t.Errorf("name: dapat %v", got)
	}
	if got := p.GetTitle(); got != "nilai-title" {
		t.Errorf("title: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetDate(); got != "nilai-date" {
		t.Errorf("date: dapat %v", got)
	}
	if got := p.GetPublisher(); got != "nilai-publisher" {
		t.Errorf("publisher: dapat %v", got)
	}
	if got := p.GetDescription(); got != "nilai-description" {
		t.Errorf("description: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetInstance(); got != "nilai-instance" {
		t.Errorf("instance: dapat %v", got)
	}
	if got := p.GetProperty(); got != "nilai-property" {
		t.Errorf("property: dapat %v", got)
	}
	if got := p.GetApplies(); got != "nilai-applies" {
		t.Errorf("applies: dapat %v", got)
	}
}
func TestConsentParser(t *testing.T) {
	p := parser.NewConsentParser(map[string]any{
		"resourceType": "Consent",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "Consent" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestContractParser(t *testing.T) {
	p := parser.NewContractParser(map[string]any{
		"resourceType": "Contract",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "Contract" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestDetectedIssueParser(t *testing.T) {
	p := parser.NewDetectedIssueParser(map[string]any{
		"resourceType": "DetectedIssue",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "DetectedIssue" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestDeviceDefinitionParser(t *testing.T) {
	p := parser.NewDeviceDefinitionParser(map[string]any{
		"resourceType": "DeviceDefinition",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "DeviceDefinition" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestDeviceMetricParser(t *testing.T) {
	p := parser.NewDeviceMetricParser(map[string]any{
		"resourceType":      "DeviceMetric",
		"id":                "nilai-id",
		"identifier":        "nilai-identifier",
		"status":            "nilai-status",
		"type":              "nilai-type",
		"category":          "nilai-category",
		"measurementPeriod": "nilai-measurementPeriod",
	})

	if got := p.GetResourceType(); got != "DeviceMetric" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetType(); got != "nilai-type" {
		t.Errorf("type: dapat %v", got)
	}
	if got := p.GetCategory(); got != "nilai-category" {
		t.Errorf("category: dapat %v", got)
	}
	if got := p.GetMeasurementPeriod(); got != "nilai-measurementPeriod" {
		t.Errorf("measurementPeriod: dapat %v", got)
	}
}
func TestDocumentManifestParser(t *testing.T) {
	p := parser.NewDocumentManifestParser(map[string]any{
		"resourceType": "DocumentManifest",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"type":         "nilai-type",
		"subject":      "nilai-subject",
		"created":      "nilai-created",
		"source":       "nilai-source",
	})

	if got := p.GetResourceType(); got != "DocumentManifest" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetType(); got != "nilai-type" {
		t.Errorf("type: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetCreated(); got != "nilai-created" {
		t.Errorf("created: dapat %v", got)
	}
	if got := p.GetSource(); got != "nilai-source" {
		t.Errorf("source: dapat %v", got)
	}
}
func TestEffectEvidenceSynthesisParser(t *testing.T) {
	p := parser.NewEffectEvidenceSynthesisParser(map[string]any{
		"resourceType": "EffectEvidenceSynthesis",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "EffectEvidenceSynthesis" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestEndpointParser(t *testing.T) {
	p := parser.NewEndpointParser(map[string]any{
		"resourceType":         "Endpoint",
		"id":                   "nilai-id",
		"identifier":           "nilai-identifier",
		"status":               "nilai-status",
		"connectionType":       "nilai-connectionType",
		"name":                 "nilai-name",
		"managingOrganization": "nilai-managingOrganization",
		"contact":              "nilai-contact",
		"period":               "nilai-period",
		"payloadType":          "nilai-payloadType",
		"payloadMimeType":      "nilai-payloadMimeType",
		"address":              "nilai-address",
		"header":               "nilai-header",
	})

	if got := p.GetResourceType(); got != "Endpoint" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetConnectionType(); got != "nilai-connectionType" {
		t.Errorf("connectionType: dapat %v", got)
	}
	if got := p.GetName(); got != "nilai-name" {
		t.Errorf("name: dapat %v", got)
	}
	if got := p.GetManagingOrganization(); got != "nilai-managingOrganization" {
		t.Errorf("managingOrganization: dapat %v", got)
	}
	if got := p.GetContact(); got != "nilai-contact" {
		t.Errorf("contact: dapat %v", got)
	}
	if got := p.GetPeriod(); got != "nilai-period" {
		t.Errorf("period: dapat %v", got)
	}
	if got := p.GetPayloadType(); got != "nilai-payloadType" {
		t.Errorf("payloadType: dapat %v", got)
	}
	if got := p.GetPayloadMimeType(); got != "nilai-payloadMimeType" {
		t.Errorf("payloadMimeType: dapat %v", got)
	}
	if got := p.GetAddress(); got != "nilai-address" {
		t.Errorf("address: dapat %v", got)
	}
	if got := p.GetHeader(); got != "nilai-header" {
		t.Errorf("header: dapat %v", got)
	}
}
func TestEnrollmentRequestParser(t *testing.T) {
	p := parser.NewEnrollmentRequestParser(map[string]any{
		"resourceType": "EnrollmentRequest",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "EnrollmentRequest" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestEnrollmentResponseParser(t *testing.T) {
	p := parser.NewEnrollmentResponseParser(map[string]any{
		"resourceType": "EnrollmentResponse",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"request":      "nilai-request",
		"outcome":      "nilai-outcome",
		"disposition":  "nilai-disposition",
		"created":      "nilai-created",
	})

	if got := p.GetResourceType(); got != "EnrollmentResponse" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetRequest(); got != "nilai-request" {
		t.Errorf("request: dapat %v", got)
	}
	if got := p.GetOutcome(); got != "nilai-outcome" {
		t.Errorf("outcome: dapat %v", got)
	}
	if got := p.GetDisposition(); got != "nilai-disposition" {
		t.Errorf("disposition: dapat %v", got)
	}
	if got := p.GetCreated(); got != "nilai-created" {
		t.Errorf("created: dapat %v", got)
	}
}
func TestEventDefinitionParser(t *testing.T) {
	p := parser.NewEventDefinitionParser(map[string]any{
		"resourceType": "EventDefinition",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "EventDefinition" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestEvidenceParser(t *testing.T) {
	p := parser.NewEvidenceParser(map[string]any{
		"resourceType": "Evidence",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "Evidence" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestEvidenceVariableParser(t *testing.T) {
	p := parser.NewEvidenceVariableParser(map[string]any{
		"resourceType": "EvidenceVariable",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "EvidenceVariable" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestExampleScenarioParser(t *testing.T) {
	p := parser.NewExampleScenarioParser(map[string]any{
		"resourceType": "ExampleScenario",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "ExampleScenario" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestExplanationOfBenefitParser(t *testing.T) {
	p := parser.NewExplanationOfBenefitParser(map[string]any{
		"resourceType": "ExplanationOfBenefit",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"type":         "nilai-type",
		"use":          "nilai-use",
		"patient":      "nilai-patient",
		"created":      "nilai-created",
		"provider":     "nilai-provider",
	})

	if got := p.GetResourceType(); got != "ExplanationOfBenefit" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetType(); got != "nilai-type" {
		t.Errorf("type: dapat %v", got)
	}
	if got := p.GetUse(); got != "nilai-use" {
		t.Errorf("use: dapat %v", got)
	}
	if got := p.GetPatient(); got != "nilai-patient" {
		t.Errorf("patient: dapat %v", got)
	}
	if got := p.GetCreated(); got != "nilai-created" {
		t.Errorf("created: dapat %v", got)
	}
	if got := p.GetProvider(); got != "nilai-provider" {
		t.Errorf("provider: dapat %v", got)
	}
}
func TestFlagParser(t *testing.T) {
	p := parser.NewFlagParser(map[string]any{
		"resourceType": "Flag",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "Flag" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestGuidanceResponseParser(t *testing.T) {
	p := parser.NewGuidanceResponseParser(map[string]any{
		"resourceType": "GuidanceResponse",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "GuidanceResponse" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestHealthcareServiceParser(t *testing.T) {
	p := parser.NewHealthcareServiceParser(map[string]any{
		"resourceType": "HealthcareService",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"name":         "nilai-name",
		"type":         "nilai-type",
		"location":     "nilai-location",
		"providedBy":   "nilai-providedBy",
	})

	if got := p.GetResourceType(); got != "HealthcareService" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetName(); got != "nilai-name" {
		t.Errorf("name: dapat %v", got)
	}
	if got := p.GetType(); got != "nilai-type" {
		t.Errorf("type: dapat %v", got)
	}
	if got := p.GetLocation(); got != "nilai-location" {
		t.Errorf("location: dapat %v", got)
	}
	if got := p.GetProvidedBy(); got != "nilai-providedBy" {
		t.Errorf("providedBy: dapat %v", got)
	}
}
func TestImagingManifestParser(t *testing.T) {
	p := parser.NewImagingManifestParser(map[string]any{
		"resourceType": "ImagingManifest",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "ImagingManifest" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestImagingSelectionParser(t *testing.T) {
	p := parser.NewImagingSelectionParser(map[string]any{
		"resourceType": "ImagingSelection",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "ImagingSelection" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestImmunizationEvaluationParser(t *testing.T) {
	p := parser.NewImmunizationEvaluationParser(map[string]any{
		"resourceType": "ImmunizationEvaluation",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "ImmunizationEvaluation" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestImmunizationRecommendationParser(t *testing.T) {
	p := parser.NewImmunizationRecommendationParser(map[string]any{
		"resourceType": "ImmunizationRecommendation",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "ImmunizationRecommendation" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestInsurancePlanParser(t *testing.T) {
	p := parser.NewInsurancePlanParser(map[string]any{
		"resourceType":   "InsurancePlan",
		"id":             "nilai-id",
		"identifier":     "nilai-identifier",
		"status":         "nilai-status",
		"name":           "nilai-name",
		"type":           "nilai-type",
		"administeredBy": "nilai-administeredBy",
		"coverage":       "nilai-coverage",
	})

	if got := p.GetResourceType(); got != "InsurancePlan" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetName(); got != "nilai-name" {
		t.Errorf("name: dapat %v", got)
	}
	if got := p.GetType(); got != "nilai-type" {
		t.Errorf("type: dapat %v", got)
	}
	if got := p.GetAdministeredBy(); got != "nilai-administeredBy" {
		t.Errorf("administeredBy: dapat %v", got)
	}
	if got := p.GetCoverage(); got != "nilai-coverage" {
		t.Errorf("coverage: dapat %v", got)
	}
}
func TestLibraryParser(t *testing.T) {
	p := parser.NewLibraryParser(map[string]any{
		"resourceType": "Library",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "Library" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestLinkageParser(t *testing.T) {
	p := parser.NewLinkageParser(map[string]any{
		"resourceType": "Linkage",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "Linkage" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestListResourceParser(t *testing.T) {
	p := parser.NewListResourceParser(map[string]any{
		"resourceType": "ListResource",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"mode":         "nilai-mode",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"date":         "nilai-date",
		"title":        "nilai-title",
	})

	if got := p.GetResourceType(); got != "ListResource" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetMode(); got != "nilai-mode" {
		t.Errorf("mode: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetDate(); got != "nilai-date" {
		t.Errorf("date: dapat %v", got)
	}
	if got := p.GetTitle(); got != "nilai-title" {
		t.Errorf("title: dapat %v", got)
	}
}
func TestMeasureParser(t *testing.T) {
	p := parser.NewMeasureParser(map[string]any{
		"resourceType": "Measure",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "Measure" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestMeasureReportParser(t *testing.T) {
	p := parser.NewMeasureReportParser(map[string]any{
		"resourceType": "MeasureReport",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "MeasureReport" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestMediaParser(t *testing.T) {
	p := parser.NewMediaParser(map[string]any{
		"resourceType": "Media",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "Media" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestMedicationKnowledgeParser(t *testing.T) {
	p := parser.NewMedicationKnowledgeParser(map[string]any{
		"resourceType": "MedicationKnowledge",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"manufacturer": "nilai-manufacturer",
		"doseForm":     "nilai-doseForm",
	})

	if got := p.GetResourceType(); got != "MedicationKnowledge" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetManufacturer(); got != "nilai-manufacturer" {
		t.Errorf("manufacturer: dapat %v", got)
	}
	if got := p.GetDoseForm(); got != "nilai-doseForm" {
		t.Errorf("doseForm: dapat %v", got)
	}
}
func TestMedicinalProductParser(t *testing.T) {
	p := parser.NewMedicinalProductParser(map[string]any{
		"resourceType": "MedicinalProduct",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"name":         "nilai-name",
		"manufacturer": "nilai-manufacturer",
		"productClass": "nilai-productClass",
	})

	if got := p.GetResourceType(); got != "MedicinalProduct" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetName(); got != "nilai-name" {
		t.Errorf("name: dapat %v", got)
	}
	if got := p.GetManufacturer(); got != "nilai-manufacturer" {
		t.Errorf("manufacturer: dapat %v", got)
	}
	if got := p.GetProductClass(); got != "nilai-productClass" {
		t.Errorf("productClass: dapat %v", got)
	}
}
func TestMedicinalProductAuthorizationParser(t *testing.T) {
	p := parser.NewMedicinalProductAuthorizationParser(map[string]any{
		"resourceType":   "MedicinalProductAuthorization",
		"id":             "nilai-id",
		"identifier":     "nilai-identifier",
		"status":         "nilai-status",
		"country":        "nilai-country",
		"regulator":      "nilai-regulator",
		"validityPeriod": "nilai-validityPeriod",
	})

	if got := p.GetResourceType(); got != "MedicinalProductAuthorization" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCountry(); got != "nilai-country" {
		t.Errorf("country: dapat %v", got)
	}
	if got := p.GetRegulator(); got != "nilai-regulator" {
		t.Errorf("regulator: dapat %v", got)
	}
	if got := p.GetValidityPeriod(); got != "nilai-validityPeriod" {
		t.Errorf("validityPeriod: dapat %v", got)
	}
}
func TestMedicinalProductContraindicationParser(t *testing.T) {
	p := parser.NewMedicinalProductContraindicationParser(map[string]any{
		"resourceType": "MedicinalProductContraindication",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"subject":      "nilai-subject",
		"disease":      "nilai-disease",
		"comorbidity":  "nilai-comorbidity",
	})

	if got := p.GetResourceType(); got != "MedicinalProductContraindication" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetDisease(); got != "nilai-disease" {
		t.Errorf("disease: dapat %v", got)
	}
	if got := p.GetComorbidity(); got != "nilai-comorbidity" {
		t.Errorf("comorbidity: dapat %v", got)
	}
}
func TestMedicinalProductIndicationParser(t *testing.T) {
	p := parser.NewMedicinalProductIndicationParser(map[string]any{
		"resourceType": "MedicinalProductIndication",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"subject":      "nilai-subject",
		"disease":      "nilai-disease",
		"comorbidity":  "nilai-comorbidity",
	})

	if got := p.GetResourceType(); got != "MedicinalProductIndication" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetDisease(); got != "nilai-disease" {
		t.Errorf("disease: dapat %v", got)
	}
	if got := p.GetComorbidity(); got != "nilai-comorbidity" {
		t.Errorf("comorbidity: dapat %v", got)
	}
}
func TestMedicinalProductIngredientParser(t *testing.T) {
	p := parser.NewMedicinalProductIngredientParser(map[string]any{
		"resourceType": "MedicinalProductIngredient",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"role":         "nilai-role",
		"substance":    "nilai-substance",
		"quantity":     "nilai-quantity",
	})

	if got := p.GetResourceType(); got != "MedicinalProductIngredient" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetRole(); got != "nilai-role" {
		t.Errorf("role: dapat %v", got)
	}
	if got := p.GetSubstance(); got != "nilai-substance" {
		t.Errorf("substance: dapat %v", got)
	}
	if got := p.GetQuantity(); got != "nilai-quantity" {
		t.Errorf("quantity: dapat %v", got)
	}
}
func TestMedicinalProductInteractionParser(t *testing.T) {
	p := parser.NewMedicinalProductInteractionParser(map[string]any{
		"resourceType": "MedicinalProductInteraction",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"subject":      "nilai-subject",
		"description":  "nilai-description",
		"interactant":  "nilai-interactant",
	})

	if got := p.GetResourceType(); got != "MedicinalProductInteraction" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetDescription(); got != "nilai-description" {
		t.Errorf("description: dapat %v", got)
	}
	if got := p.GetInteractant(); got != "nilai-interactant" {
		t.Errorf("interactant: dapat %v", got)
	}
}
func TestMedicinalProductManufacturedParser(t *testing.T) {
	p := parser.NewMedicinalProductManufacturedParser(map[string]any{
		"resourceType": "MedicinalProductManufactured",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"quantity":     "nilai-quantity",
		"manufacturer": "nilai-manufacturer",
		"ingredient":   "nilai-ingredient",
	})

	if got := p.GetResourceType(); got != "MedicinalProductManufactured" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetQuantity(); got != "nilai-quantity" {
		t.Errorf("quantity: dapat %v", got)
	}
	if got := p.GetManufacturer(); got != "nilai-manufacturer" {
		t.Errorf("manufacturer: dapat %v", got)
	}
	if got := p.GetIngredient(); got != "nilai-ingredient" {
		t.Errorf("ingredient: dapat %v", got)
	}
}
func TestMedicinalProductPackagedParser(t *testing.T) {
	p := parser.NewMedicinalProductPackagedParser(map[string]any{
		"resourceType": "MedicinalProductPackaged",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"packageItem":  "nilai-packageItem",
		"description":  "nilai-description",
		"manufacturer": "nilai-manufacturer",
	})

	if got := p.GetResourceType(); got != "MedicinalProductPackaged" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetPackageItem(); got != "nilai-packageItem" {
		t.Errorf("packageItem: dapat %v", got)
	}
	if got := p.GetDescription(); got != "nilai-description" {
		t.Errorf("description: dapat %v", got)
	}
	if got := p.GetManufacturer(); got != "nilai-manufacturer" {
		t.Errorf("manufacturer: dapat %v", got)
	}
}
func TestMedicinalProductPharmaceuticalParser(t *testing.T) {
	p := parser.NewMedicinalProductPharmaceuticalParser(map[string]any{
		"resourceType": "MedicinalProductPharmaceutical",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"route":        "nilai-route",
		"doseForm":     "nilai-doseForm",
		"ingredient":   "nilai-ingredient",
	})

	if got := p.GetResourceType(); got != "MedicinalProductPharmaceutical" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetRoute(); got != "nilai-route" {
		t.Errorf("route: dapat %v", got)
	}
	if got := p.GetDoseForm(); got != "nilai-doseForm" {
		t.Errorf("doseForm: dapat %v", got)
	}
	if got := p.GetIngredient(); got != "nilai-ingredient" {
		t.Errorf("ingredient: dapat %v", got)
	}
}
func TestMedicinalProductUndesirableEffectParser(t *testing.T) {
	p := parser.NewMedicinalProductUndesirableEffectParser(map[string]any{
		"resourceType":           "MedicinalProductUndesirableEffect",
		"id":                     "nilai-id",
		"identifier":             "nilai-identifier",
		"status":                 "nilai-status",
		"subject":                "nilai-subject",
		"symptomConditionEffect": "nilai-symptomConditionEffect",
		"classification":         "nilai-classification",
	})

	if got := p.GetResourceType(); got != "MedicinalProductUndesirableEffect" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetSymptomConditionEffect(); got != "nilai-symptomConditionEffect" {
		t.Errorf("symptomConditionEffect: dapat %v", got)
	}
	if got := p.GetClassification(); got != "nilai-classification" {
		t.Errorf("classification: dapat %v", got)
	}
}
func TestMessageDefinitionParser(t *testing.T) {
	p := parser.NewMessageDefinitionParser(map[string]any{
		"resourceType": "MessageDefinition",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "MessageDefinition" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestMessageHeaderParser(t *testing.T) {
	p := parser.NewMessageHeaderParser(map[string]any{
		"resourceType": "MessageHeader",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "MessageHeader" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestObservationDefinitionParser(t *testing.T) {
	p := parser.NewObservationDefinitionParser(map[string]any{
		"resourceType": "ObservationDefinition",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"category":     "nilai-category",
		"method":       "nilai-method",
	})

	if got := p.GetResourceType(); got != "ObservationDefinition" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetCategory(); got != "nilai-category" {
		t.Errorf("category: dapat %v", got)
	}
	if got := p.GetMethod(); got != "nilai-method" {
		t.Errorf("method: dapat %v", got)
	}
}
func TestOperationOutcomeParser(t *testing.T) {
	p := parser.NewOperationOutcomeParser(map[string]any{
		"resourceType": "OperationOutcome",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "OperationOutcome" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestOrganizationAffiliationParser(t *testing.T) {
	p := parser.NewOrganizationAffiliationParser(map[string]any{
		"resourceType":              "OrganizationAffiliation",
		"id":                        "nilai-id",
		"identifier":                "nilai-identifier",
		"status":                    "nilai-status",
		"organization":              "nilai-organization",
		"participatingOrganization": "nilai-participatingOrganization",
		"network":                   "nilai-network",
		"code":                      "nilai-code",
	})

	if got := p.GetResourceType(); got != "OrganizationAffiliation" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetOrganization(); got != "nilai-organization" {
		t.Errorf("organization: dapat %v", got)
	}
	if got := p.GetParticipatingOrganization(); got != "nilai-participatingOrganization" {
		t.Errorf("participatingOrganization: dapat %v", got)
	}
	if got := p.GetNetwork(); got != "nilai-network" {
		t.Errorf("network: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
}
func TestPersonParser(t *testing.T) {
	p := parser.NewPersonParser(map[string]any{
		"resourceType": "Person",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "Person" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestPlanDefinitionParser(t *testing.T) {
	p := parser.NewPlanDefinitionParser(map[string]any{
		"resourceType": "PlanDefinition",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "PlanDefinition" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestPractitionerRoleParser(t *testing.T) {
	p := parser.NewPractitionerRoleParser(map[string]any{
		"resourceType":           "PractitionerRole",
		"id":                     "nilai-id",
		"identifier":             "nilai-identifier",
		"active":                 "nilai-active",
		"practitioner":           "nilai-practitioner",
		"organization":           "nilai-organization",
		"code":                   "nilai-code",
		"specialty":              "nilai-specialty",
		"location":               "nilai-location",
		"healthcareService":      "nilai-healthcareService",
		"telecom":                "nilai-telecom",
		"availableTime":          "nilai-availableTime",
		"notAvailable":           "nilai-notAvailable",
		"endpoint":               "nilai-endpoint",
		"availabilityExceptions": "nilai-availabilityExceptions",
		"extension":              "nilai-extension",
	})

	if got := p.GetResourceType(); got != "PractitionerRole" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetActive(); got != "nilai-active" {
		t.Errorf("active: dapat %v", got)
	}
	if got := p.GetPractitioner(); got != "nilai-practitioner" {
		t.Errorf("practitioner: dapat %v", got)
	}
	if got := p.GetOrganization(); got != "nilai-organization" {
		t.Errorf("organization: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSpecialty(); got != "nilai-specialty" {
		t.Errorf("specialty: dapat %v", got)
	}
	if got := p.GetLocation(); got != "nilai-location" {
		t.Errorf("location: dapat %v", got)
	}
	if got := p.GetHealthcareService(); got != "nilai-healthcareService" {
		t.Errorf("healthcareService: dapat %v", got)
	}
	if got := p.GetTelecom(); got != "nilai-telecom" {
		t.Errorf("telecom: dapat %v", got)
	}
	if got := p.GetAvailableTime(); got != "nilai-availableTime" {
		t.Errorf("availableTime: dapat %v", got)
	}
	if got := p.GetNotAvailable(); got != "nilai-notAvailable" {
		t.Errorf("notAvailable: dapat %v", got)
	}
	if got := p.GetEndpoint(); got != "nilai-endpoint" {
		t.Errorf("endpoint: dapat %v", got)
	}
	if got := p.GetAvailabilityExceptions(); got != "nilai-availabilityExceptions" {
		t.Errorf("availabilityExceptions: dapat %v", got)
	}
	if got := p.GetExtension(); got != "nilai-extension" {
		t.Errorf("extension: dapat %v", got)
	}
}
func TestProvenanceParser(t *testing.T) {
	p := parser.NewProvenanceParser(map[string]any{
		"resourceType": "Provenance",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "Provenance" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestPurificationDecisionParser(t *testing.T) {
	p := parser.NewPurificationDecisionParser(map[string]any{
		"resourceType":  "PurificationDecision",
		"id":            "nilai-id",
		"identifier":    "nilai-identifier",
		"status":        "nilai-status",
		"insurer":       "nilai-insurer",
		"provider":      "nilai-provider",
		"claimResponse": "nilai-claimResponse",
		"created":       "nilai-created",
	})

	if got := p.GetResourceType(); got != "PurificationDecision" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetInsurer(); got != "nilai-insurer" {
		t.Errorf("insurer: dapat %v", got)
	}
	if got := p.GetProvider(); got != "nilai-provider" {
		t.Errorf("provider: dapat %v", got)
	}
	if got := p.GetClaimResponse(); got != "nilai-claimResponse" {
		t.Errorf("claimResponse: dapat %v", got)
	}
	if got := p.GetCreated(); got != "nilai-created" {
		t.Errorf("created: dapat %v", got)
	}
}
func TestQuestionnaireParser(t *testing.T) {
	p := parser.NewQuestionnaireParser(map[string]any{
		"resourceType": "Questionnaire",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "Questionnaire" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestRequestGroupParser(t *testing.T) {
	p := parser.NewRequestGroupParser(map[string]any{
		"resourceType": "RequestGroup",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "RequestGroup" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestResearchDefinitionParser(t *testing.T) {
	p := parser.NewResearchDefinitionParser(map[string]any{
		"resourceType": "ResearchDefinition",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "ResearchDefinition" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestResearchElementDefinitionParser(t *testing.T) {
	p := parser.NewResearchElementDefinitionParser(map[string]any{
		"resourceType": "ResearchElementDefinition",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "ResearchElementDefinition" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestResearchStudyParser(t *testing.T) {
	p := parser.NewResearchStudyParser(map[string]any{
		"resourceType": "ResearchStudy",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"title":        "nilai-title",
		"protocol":     "nilai-protocol",
		"sponsor":      "nilai-sponsor",
		"phase":        "nilai-phase",
	})

	if got := p.GetResourceType(); got != "ResearchStudy" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetTitle(); got != "nilai-title" {
		t.Errorf("title: dapat %v", got)
	}
	if got := p.GetProtocol(); got != "nilai-protocol" {
		t.Errorf("protocol: dapat %v", got)
	}
	if got := p.GetSponsor(); got != "nilai-sponsor" {
		t.Errorf("sponsor: dapat %v", got)
	}
	if got := p.GetPhase(); got != "nilai-phase" {
		t.Errorf("phase: dapat %v", got)
	}
}
func TestResearchSubjectParser(t *testing.T) {
	p := parser.NewResearchSubjectParser(map[string]any{
		"resourceType": "ResearchSubject",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "ResearchSubject" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestResourceGuideParser(t *testing.T) {
	p := parser.NewResourceGuideParser(map[string]any{
		"resourceType": "ResourceGuide",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"name":         "nilai-name",
		"description":  "nilai-description",
		"version":      "nilai-version",
		"publisher":    "nilai-publisher",
	})

	if got := p.GetResourceType(); got != "ResourceGuide" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetName(); got != "nilai-name" {
		t.Errorf("name: dapat %v", got)
	}
	if got := p.GetDescription(); got != "nilai-description" {
		t.Errorf("description: dapat %v", got)
	}
	if got := p.GetVersion(); got != "nilai-version" {
		t.Errorf("version: dapat %v", got)
	}
	if got := p.GetPublisher(); got != "nilai-publisher" {
		t.Errorf("publisher: dapat %v", got)
	}
}
func TestRiskEvidenceSynthesisParser(t *testing.T) {
	p := parser.NewRiskEvidenceSynthesisParser(map[string]any{
		"resourceType": "RiskEvidenceSynthesis",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "RiskEvidenceSynthesis" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestScheduleParser(t *testing.T) {
	p := parser.NewScheduleParser(map[string]any{
		"resourceType": "Schedule",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "Schedule" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestSearchParameterParser(t *testing.T) {
	p := parser.NewSearchParameterParser(map[string]any{
		"resourceType": "SearchParameter",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "SearchParameter" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestSlotParser(t *testing.T) {
	p := parser.NewSlotParser(map[string]any{
		"resourceType": "Slot",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "Slot" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestSpecimenDefinitionParser(t *testing.T) {
	p := parser.NewSpecimenDefinitionParser(map[string]any{
		"resourceType": "SpecimenDefinition",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"type":         "nilai-type",
		"collection":   "nilai-collection",
		"preparation":  "nilai-preparation",
	})

	if got := p.GetResourceType(); got != "SpecimenDefinition" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetType(); got != "nilai-type" {
		t.Errorf("type: dapat %v", got)
	}
	if got := p.GetCollection(); got != "nilai-collection" {
		t.Errorf("collection: dapat %v", got)
	}
	if got := p.GetPreparation(); got != "nilai-preparation" {
		t.Errorf("preparation: dapat %v", got)
	}
}
func TestSubscriptionParser(t *testing.T) {
	p := parser.NewSubscriptionParser(map[string]any{
		"resourceType": "Subscription",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "Subscription" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestSubstanceParser(t *testing.T) {
	p := parser.NewSubstanceParser(map[string]any{
		"resourceType": "Substance",
		"id":           "nilai-id",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"text":         "nilai-text",
	})

	if got := p.GetResourceType(); got != "Substance" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetText(); got != "nilai-text" {
		t.Errorf("text: dapat %v", got)
	}
}
func TestSubstanceNucleicAcidParser(t *testing.T) {
	p := parser.NewSubstanceNucleicAcidParser(map[string]any{
		"resourceType": "SubstanceNucleicAcid",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "SubstanceNucleicAcid" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestSubstancePolymerParser(t *testing.T) {
	p := parser.NewSubstancePolymerParser(map[string]any{
		"resourceType": "SubstancePolymer",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "SubstancePolymer" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestSubstanceProteinParser(t *testing.T) {
	p := parser.NewSubstanceProteinParser(map[string]any{
		"resourceType": "SubstanceProtein",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "SubstanceProtein" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestSubstanceReferenceInformationParser(t *testing.T) {
	p := parser.NewSubstanceReferenceInformationParser(map[string]any{
		"resourceType": "SubstanceReferenceInformation",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"comment":      "nilai-comment",
		"gene":         "nilai-gene",
		"geneElement":  "nilai-geneElement",
	})

	if got := p.GetResourceType(); got != "SubstanceReferenceInformation" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetComment(); got != "nilai-comment" {
		t.Errorf("comment: dapat %v", got)
	}
	if got := p.GetGene(); got != "nilai-gene" {
		t.Errorf("gene: dapat %v", got)
	}
	if got := p.GetGeneElement(); got != "nilai-geneElement" {
		t.Errorf("geneElement: dapat %v", got)
	}
}
func TestSubstanceSourceMaterialParser(t *testing.T) {
	p := parser.NewSubstanceSourceMaterialParser(map[string]any{
		"resourceType": "SubstanceSourceMaterial",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "SubstanceSourceMaterial" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestSubstanceSpecificationParser(t *testing.T) {
	p := parser.NewSubstanceSpecificationParser(map[string]any{
		"resourceType": "SubstanceSpecification",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "SubstanceSpecification" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestTestReportParser(t *testing.T) {
	p := parser.NewTestReportParser(map[string]any{
		"resourceType": "TestReport",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "TestReport" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestTestScriptParser(t *testing.T) {
	p := parser.NewTestScriptParser(map[string]any{
		"resourceType": "TestScript",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "TestScript" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
func TestVerificationResultParser(t *testing.T) {
	p := parser.NewVerificationResultParser(map[string]any{
		"resourceType": "VerificationResult",
		"id":           "nilai-id",
		"identifier":   "nilai-identifier",
		"status":       "nilai-status",
		"code":         "nilai-code",
		"subject":      "nilai-subject",
		"encounter":    "nilai-encounter",
	})

	if got := p.GetResourceType(); got != "VerificationResult" {
		t.Errorf("resourceType: dapat %v", got)
	}
	if got := p.GetId(); got != "nilai-id" {
		t.Errorf("id: dapat %v", got)
	}
	if got := p.GetIdentifier(); got != "nilai-identifier" {
		t.Errorf("identifier: dapat %v", got)
	}
	if got := p.GetStatus(); got != "nilai-status" {
		t.Errorf("status: dapat %v", got)
	}
	if got := p.GetCode(); got != "nilai-code" {
		t.Errorf("code: dapat %v", got)
	}
	if got := p.GetSubject(); got != "nilai-subject" {
		t.Errorf("subject: dapat %v", got)
	}
	if got := p.GetEncounter(); got != "nilai-encounter" {
		t.Errorf("encounter: dapat %v", got)
	}
}
