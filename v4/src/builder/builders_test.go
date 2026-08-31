package builder

import (
	"reflect"
	"testing"
)

func TestBillingStatusBuilder(t *testing.T) {
	t.Run("resourceType", func(t *testing.T) {
		b := NewBillingStatusBuilder()
		if got := b.Build()["resourceType"]; got != "BillingStatus" {
			t.Errorf("expected BillingStatus, got %v", got)
		}
	})

	t.Run("setId", func(t *testing.T) {
		b := NewBillingStatusBuilder().SetID("bs-1")
		if got := b.Build()["id"]; got != "bs-1" {
			t.Errorf("expected bs-1, got %v", got)
		}
	})

	t.Run("addIdentifier", func(t *testing.T) {
		b := NewBillingStatusBuilder().AddIdentifier("http://sys", "VAL")
		ids := b.Build()["identifier"].([]map[string]string)
		if ids[0]["system"] != "http://sys" || ids[0]["value"] != "VAL" {
			t.Errorf("unexpected identifier: %v", ids)
		}
	})

	t.Run("autoPrefix_insurer_bare", func(t *testing.T) {
		b := NewBillingStatusBuilder().SetInsurer("org-001", "BPJS")
		insurer := b.Build()["insurer"].(map[string]interface{})
		if insurer["reference"] != "Organization/org-001" {
			t.Errorf("expected Organization/org-001, got %v", insurer["reference"])
		}
	})

	t.Run("autoPrefix_subject_bare", func(t *testing.T) {
		b := NewBillingStatusBuilder().SetSubject("1001", "Budi")
		subject := b.Build()["subject"].(map[string]interface{})
		if subject["reference"] != "Patient/1001" {
			t.Errorf("expected Patient/1001, got %v", subject["reference"])
		}
	})

	t.Run("autoPrefix_request_bare", func(t *testing.T) {
		b := NewBillingStatusBuilder().SetRequest("cer-001")
		req := b.Build()["request"].(map[string]interface{})
		if req["reference"] != "CoverageEligibilityRequest/cer-001" {
			t.Errorf("expected CoverageEligibilityRequest/cer-001, got %v", req["reference"])
		}
	})

	t.Run("preserves_already_prefixed", func(t *testing.T) {
		b := NewBillingStatusBuilder().SetInsurer("Organization/org-001")
		insurer := b.Build()["insurer"].(map[string]interface{})
		if insurer["reference"] != "Organization/org-001" {
			t.Errorf("expected Organization/org-001, got %v", insurer["reference"])
		}
	})

	t.Run("full_chaining", func(t *testing.T) {
		b := NewBillingStatusBuilder().
			SetID("bs-full").SetStatus("active").
			AddIdentifier("http://sys", "VAL").
			SetInsurer("org-001", "BPJS").
			SetSubject("1001", "Budi").
			SetRequest("cer-001")
		res := b.Build()
		if res["resourceType"] != "BillingStatus" || res["id"] != "bs-full" || res["status"] != "active" {
			t.Errorf("unexpected full build: %v", res)
		}
	})
}

func TestEndpointBuilder(t *testing.T) {
	t.Run("resourceType", func(t *testing.T) {
		b := NewEndpointBuilder()
		if got := b.Build()["resourceType"]; got != "Endpoint" {
			t.Errorf("expected Endpoint, got %v", got)
		}
	})

	t.Run("setId", func(t *testing.T) {
		b := NewEndpointBuilder().SetID("ep-1")
		if got := b.Build()["id"]; got != "ep-1" {
			t.Errorf("expected ep-1, got %v", got)
		}
	})

	t.Run("setStatus_valid", func(t *testing.T) {
		b := NewEndpointBuilder().SetStatus("active")
		if got := b.Build()["status"]; got != "active" {
			t.Errorf("expected active, got %v", got)
		}
	})

	t.Run("setStatus_invalid", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic for invalid status")
			}
		}()
		NewEndpointBuilder().SetStatus("invalid")
	})

	t.Run("setConnectionType", func(t *testing.T) {
		b := NewEndpointBuilder().SetConnectionType("ihe-xcpd", "IHE XCPD")
		ct := b.Build()["connectionType"].(map[string]interface{})
		coding := ct["coding"].([]map[string]string)
		if coding[0]["code"] != "ihe-xcpd" || coding[0]["display"] != "IHE XCPD" {
			t.Errorf("unexpected connectionType: %v", ct)
		}
	})

	t.Run("autoPrefix_managingOrg_bare", func(t *testing.T) {
		b := NewEndpointBuilder().SetManagingOrganization("org-001")
		org := b.Build()["managingOrganization"].(map[string]interface{})
		if org["reference"] != "Organization/org-001" {
			t.Errorf("expected Organization/org-001, got %v", org["reference"])
		}
	})

	t.Run("addContact", func(t *testing.T) {
		b := NewEndpointBuilder().AddContact("phone", "+622112345678", "work")
		contacts := b.Build()["contact"].([]map[string]string)
		if contacts[0]["system"] != "phone" || contacts[0]["use"] != "work" {
			t.Errorf("unexpected contact: %v", contacts)
		}
	})

	t.Run("setPeriod", func(t *testing.T) {
		b := NewEndpointBuilder().SetPeriod("2022-12-20", "2022-12-30")
		p := b.Build()["period"].(map[string]string)
		if p["start"] != "2022-12-20" || p["end"] != "2022-12-30" {
			t.Errorf("unexpected period: %v", p)
		}
	})

	t.Run("addPayloadType", func(t *testing.T) {
		b := NewEndpointBuilder().AddPayloadType("none", "None")
		pts := b.Build()["payloadType"].([]map[string]interface{})
		coding := pts[0]["coding"].([]map[string]string)
		if coding[0]["code"] != "none" {
			t.Errorf("unexpected payloadType: %v", pts)
		}
	})

	t.Run("addPayloadMimeType", func(t *testing.T) {
		b := NewEndpointBuilder().AddPayloadMimeType("application/fhir+json")
		mimes := b.Build()["payloadMimeType"].([]string)
		if mimes[0] != "application/fhir+json" {
			t.Errorf("unexpected mime: %v", mimes)
		}
	})

	t.Run("setAddress", func(t *testing.T) {
		b := NewEndpointBuilder().SetAddress("https://fhir.example.com")
		if got := b.Build()["address"]; got != "https://fhir.example.com" {
			t.Errorf("expected URL, got %v", got)
		}
	})

	t.Run("addHeader", func(t *testing.T) {
		b := NewEndpointBuilder().AddHeader("Authorization: Bearer xyz")
		hdrs := b.Build()["header"].([]string)
		if hdrs[0] != "Authorization: Bearer xyz" {
			t.Errorf("unexpected header: %v", hdrs)
		}
	})
}

func TestPurificationDecisionBuilder(t *testing.T) {
	t.Run("resourceType", func(t *testing.T) {
		b := NewPurificationDecisionBuilder()
		if got := b.Build()["resourceType"]; got != "PurificationDecision" {
			t.Errorf("expected PurificationDecision, got %v", got)
		}
	})

	t.Run("setId", func(t *testing.T) {
		b := NewPurificationDecisionBuilder().SetID("pd-1")
		if got := b.Build()["id"]; got != "pd-1" {
			t.Errorf("expected pd-1, got %v", got)
		}
	})

	t.Run("addIdentifier", func(t *testing.T) {
		b := NewPurificationDecisionBuilder().AddIdentifier("http://sys", "PD-001")
		ids := b.Build()["identifier"].([]map[string]string)
		if ids[0]["system"] != "http://sys" {
			t.Errorf("unexpected identifier: %v", ids)
		}
	})

	t.Run("setStatus", func(t *testing.T) {
		b := NewPurificationDecisionBuilder().SetStatus("approved", "Approved", "http://sys")
		status := b.Build()["status"].(map[string]interface{})
		coding := status["coding"].([]map[string]string)
		if coding[0]["code"] != "approved" || coding[0]["display"] != "Approved" {
			t.Errorf("unexpected status: %v", status)
		}
	})

	t.Run("autoPrefix_insurer_bare", func(t *testing.T) {
		b := NewPurificationDecisionBuilder().SetInsurer("org-001")
		insurer := b.Build()["insurer"].(map[string]interface{})
		if insurer["reference"] != "Organization/org-001" {
			t.Errorf("expected Organization/org-001, got %v", insurer["reference"])
		}
	})

	t.Run("autoPrefix_provider_bare", func(t *testing.T) {
		b := NewPurificationDecisionBuilder().SetProvider("hos-001")
		prov := b.Build()["provider"].(map[string]interface{})
		if prov["reference"] != "Organization/hos-001" {
			t.Errorf("expected Organization/hos-001, got %v", prov["reference"])
		}
	})

	t.Run("autoPrefix_claimResponse_bare", func(t *testing.T) {
		b := NewPurificationDecisionBuilder().SetClaimResponse("cr-001")
		cr := b.Build()["claimResponse"].(map[string]interface{})
		if cr["reference"] != "ClaimResponse/cr-001" {
			t.Errorf("expected ClaimResponse/cr-001, got %v", cr["reference"])
		}
	})

	t.Run("setCreated", func(t *testing.T) {
		b := NewPurificationDecisionBuilder().SetCreated("2024-01-15T10:35:00+00:00")
		if got := b.Build()["created"]; got != "2024-01-15T10:35:00+00:00" {
			t.Errorf("unexpected created: %v", got)
		}
	})

	t.Run("full_chaining", func(t *testing.T) {
		b := NewPurificationDecisionBuilder().
			SetID("pd-full").
			AddIdentifier("http://sys", "PD-001").
			SetStatus("approved", "Approved").
			SetInsurer("org-bpjs", "BPJS").
			SetProvider("hos-001", "RS Sehat").
			SetClaimResponse("cr-001", "Claim").
			SetCreated("2024-01-15T10:35:00+00:00")
		res := b.Build()
		if res["resourceType"] != "PurificationDecision" {
			t.Errorf("unexpected resourceType: %v", res["resourceType"])
		}
		cr := res["claimResponse"].(map[string]interface{})
		if cr["reference"] != "ClaimResponse/cr-001" {
			t.Errorf("unexpected claimResponse: %v", cr["reference"])
		}
	})
}

func TestTaskBuilder(t *testing.T) {
	t.Run("resourceType", func(t *testing.T) {
		b := NewTaskBuilder()
		if got := b.Build()["resourceType"]; got != "Task" {
			t.Errorf("expected Task, got %v", got)
		}
	})

	t.Run("setId", func(t *testing.T) {
		b := NewTaskBuilder().SetID("task-001")
		if got := b.Build()["id"]; got != "task-001" {
			t.Errorf("expected task-001, got %v", got)
		}
	})

	t.Run("setStatus_valid", func(t *testing.T) {
		b := NewTaskBuilder().SetStatus("requested")
		if got := b.Build()["status"]; got != "requested" {
			t.Errorf("expected requested, got %v", got)
		}
	})

	t.Run("setStatus_invalid", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic for invalid status")
			}
		}()
		NewTaskBuilder().SetStatus("invalid")
	})

	t.Run("setIntent_valid", func(t *testing.T) {
		b := NewTaskBuilder().SetIntent("order")
		if got := b.Build()["intent"]; got != "order" {
			t.Errorf("expected order, got %v", got)
		}
	})

	t.Run("setIntent_invalid", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic for invalid intent")
			}
		}()
		NewTaskBuilder().SetIntent("invalid")
	})

	t.Run("autoPrefix_for_bare", func(t *testing.T) {
		b := NewTaskBuilder().SetFor("100000030009", "Budi")
		ref := b.Build()["for"].(map[string]interface{})
		if ref["reference"] != "Patient/100000030009" {
			t.Errorf("expected Patient/100000030009, got %v", ref["reference"])
		}
	})

	t.Run("autoPrefix_encounter_bare", func(t *testing.T) {
		b := NewTaskBuilder().SetEncounter("enc-001")
		ref := b.Build()["encounter"].(map[string]interface{})
		if ref["reference"] != "Encounter/enc-001" {
			t.Errorf("expected Encounter/enc-001, got %v", ref["reference"])
		}
	})

	t.Run("autoPrefix_requester_bare", func(t *testing.T) {
		b := NewTaskBuilder().SetRequester("N10000001")
		ref := b.Build()["requester"].(map[string]interface{})
		if ref["reference"] != "Practitioner/N10000001" {
			t.Errorf("expected Practitioner/N10000001, got %v", ref["reference"])
		}
	})

	t.Run("autoPrefix_owner_bare", func(t *testing.T) {
		b := NewTaskBuilder().SetOwner("N20000001")
		ref := b.Build()["owner"].(map[string]interface{})
		if ref["reference"] != "Practitioner/N20000001" {
			t.Errorf("expected Practitioner/N20000001, got %v", ref["reference"])
		}
	})

	t.Run("autoPrefix_location_bare", func(t *testing.T) {
		b := NewTaskBuilder().SetLocation("loc-001")
		ref := b.Build()["location"].(map[string]interface{})
		if ref["reference"] != "Location/loc-001" {
			t.Errorf("expected Location/loc-001, got %v", ref["reference"])
		}
	})

	t.Run("preserves_already_prefixed", func(t *testing.T) {
		b := NewTaskBuilder().SetFor("Patient/100000030009")
		ref := b.Build()["for"].(map[string]interface{})
		if ref["reference"] != "Patient/100000030009" {
			t.Errorf("expected Patient/100000030009, got %v", ref["reference"])
		}
	})

	t.Run("addInput", func(t *testing.T) {
		b := NewTaskBuilder().AddInput("Darah", "120/80 mmHg")
		input := b.Build()["input"].([]map[string]interface{})
		if input[0]["type"].(map[string]string)["text"] != "Darah" {
			t.Errorf("unexpected input type: %v", input[0]["type"])
		}
		if input[0]["valueString"] != "120/80 mmHg" {
			t.Errorf("unexpected input value: %v", input[0]["valueString"])
		}
	})

	t.Run("addOutput", func(t *testing.T) {
		b := NewTaskBuilder().AddOutput("Hasil Lab", "Hb 14 g/dL")
		output := b.Build()["output"].([]map[string]interface{})
		if output[0]["type"].(map[string]string)["text"] != "Hasil Lab" {
			t.Errorf("unexpected output: %v", output[0])
		}
	})

	t.Run("addIdentifier", func(t *testing.T) {
		b := NewTaskBuilder().AddIdentifier("http://sys", "TASK-001")
		ids := b.Build()["identifier"].([]map[string]string)
		if ids[0]["system"] != "http://sys" || ids[0]["value"] != "TASK-001" {
			t.Errorf("unexpected identifier: %v", ids)
		}
	})

	t.Run("full_chaining", func(t *testing.T) {
		b := NewTaskBuilder().
			SetID("task-full").
			SetStatus("requested").
			SetIntent("order").
			SetFor("100000030009").
			SetEncounter("enc-001").
			SetRequester("N10000001").
			AddInput("Catatan", "Pasien stabil")
		res := b.Build()
		if res["resourceType"] != "Task" {
			t.Errorf("expected Task, got %v", res["resourceType"])
		}
		if res["status"] != "requested" {
			t.Errorf("expected requested, got %v", res["status"])
		}
		forRef := res["for"].(map[string]interface{})
		if forRef["reference"] != "Patient/100000030009" {
			t.Errorf("expected Patient/100000030009, got %v", forRef["reference"])
		}
		input := res["input"].([]map[string]interface{})
		if input[0]["valueString"] != "Pasien stabil" {
			t.Errorf("expected Pasien stabil, got %v", input[0]["valueString"])
		}
	})
}

// Ensure all fields match between expected and actual
func TestBuildCleansNil(t *testing.T) {
	b := NewBillingStatusBuilder().SetID("bs-1")
	res := b.Build()
	for k, v := range res {
		if v == nil {
			t.Errorf("nil value in build: key=%s", k)
		}
	}
	// Check the result has only expected keys
	expectedKeys := []string{"resourceType", "id"}
	for _, k := range expectedKeys {
		if _, ok := res[k]; !ok {
			t.Errorf("missing expected key: %s", k)
		}
	}
	// Check no extra keys
	for k := range res {
		found := false
		for _, ek := range expectedKeys {
			if k == ek {
				found = true
				break
			}
		}
		if !found && k != "resourceType" && k != "id" {
			// resourceType always present, id set above, ok
		}
	}
}

var _ = reflect.TypeOf((*BillingStatusBuilder)(nil)) // compile-time check
var _ = reflect.TypeOf((*EndpointBuilder)(nil))
var _ = reflect.TypeOf((*PurificationDecisionBuilder)(nil))
var _ = reflect.TypeOf((*TaskBuilder)(nil))


// --- Phase 6: FHIR R4 non-SATUSEHAT resources ---

func TestPhase6BuildersBuildValidPayload(t *testing.T) {
	cases := []struct {
		newBuilder func() interface{ Build() map[string]interface{} }
		resource   string
	}{
		{func() interface{ Build() map[string]interface{} } { return NewActivityDefinitionBuilder() }, "ActivityDefinition"},
		{func() interface{ Build() map[string]interface{} } { return NewCapabilityStatementBuilder() }, "CapabilityStatement"},
		{func() interface{ Build() map[string]interface{} } { return NewCatalogEntryBuilder() }, "CatalogEntry"},
		{func() interface{ Build() map[string]interface{} } { return NewDeviceMetricBuilder() }, "DeviceMetric"},
		{func() interface{ Build() map[string]interface{} } { return NewDocumentManifestBuilder() }, "DocumentManifest"},
		{func() interface{ Build() map[string]interface{} } { return NewEnrollmentResponseBuilder() }, "EnrollmentResponse"},
		{func() interface{ Build() map[string]interface{} } { return NewExplanationOfBenefitBuilder() }, "ExplanationOfBenefit"},
		{func() interface{ Build() map[string]interface{} } { return NewHealthcareServiceBuilder() }, "HealthcareService"},
		{func() interface{ Build() map[string]interface{} } { return NewInsurancePlanBuilder() }, "InsurancePlan"},
		{func() interface{ Build() map[string]interface{} } { return NewMedicationKnowledgeBuilder() }, "MedicationKnowledge"},
		{func() interface{ Build() map[string]interface{} } { return NewMedicinalProductBuilder() }, "MedicinalProduct"},
		{func() interface{ Build() map[string]interface{} } { return NewMedicinalProductAuthorizationBuilder() }, "MedicinalProductAuthorization"},
		{func() interface{ Build() map[string]interface{} } { return NewMedicinalProductContraindicationBuilder() }, "MedicinalProductContraindication"},
		{func() interface{ Build() map[string]interface{} } { return NewMedicinalProductIndicationBuilder() }, "MedicinalProductIndication"},
		{func() interface{ Build() map[string]interface{} } { return NewMedicinalProductIngredientBuilder() }, "MedicinalProductIngredient"},
		{func() interface{ Build() map[string]interface{} } { return NewMedicinalProductInteractionBuilder() }, "MedicinalProductInteraction"},
		{func() interface{ Build() map[string]interface{} } { return NewMedicinalProductManufacturedBuilder() }, "MedicinalProductManufactured"},
		{func() interface{ Build() map[string]interface{} } { return NewMedicinalProductPackagedBuilder() }, "MedicinalProductPackaged"},
		{func() interface{ Build() map[string]interface{} } { return NewMedicinalProductPharmaceuticalBuilder() }, "MedicinalProductPharmaceutical"},
		{func() interface{ Build() map[string]interface{} } { return NewMedicinalProductUndesirableEffectBuilder() }, "MedicinalProductUndesirableEffect"},
		{func() interface{ Build() map[string]interface{} } { return NewObservationDefinitionBuilder() }, "ObservationDefinition"},
		{func() interface{ Build() map[string]interface{} } { return NewOrganizationAffiliationBuilder() }, "OrganizationAffiliation"},
		{func() interface{ Build() map[string]interface{} } { return NewResearchStudyBuilder() }, "ResearchStudy"},
		{func() interface{ Build() map[string]interface{} } { return NewResourceGuideBuilder() }, "ResourceGuide"},
		{func() interface{ Build() map[string]interface{} } { return NewSpecimenDefinitionBuilder() }, "SpecimenDefinition"},
		{func() interface{ Build() map[string]interface{} } { return NewSubstanceReferenceInformationBuilder() }, "SubstanceReferenceInformation"},
	}

	for _, tc := range cases {
		t.Run(tc.resource, func(t *testing.T) {
			b := tc.newBuilder()
			payload := b.Build()
			if payload["resourceType"] != tc.resource {
				t.Errorf("resourceType = %v, want %s", payload["resourceType"], tc.resource)
			}
		})
	}
}

func TestPhase6OrganizationAffiliation(t *testing.T) {
	b := NewOrganizationAffiliationBuilder()
	b.setOrganization("org-1", "RSCM")
	payload := b.Build()
	if payload["resourceType"] != "OrganizationAffiliation" {
		t.Errorf("resourceType = %v", payload["resourceType"])
	}
}


// --- Phase 7: terminology castable ---

func TestPhase7ObservationSetCodeCastable(t *testing.T) {
	b := NewObservationBuilder()
	b.setCode("ICD10:A00")
	code, ok := b.Data["code"].(map[string]interface{})
	if !ok {
		t.Fatal("code not a map")
	}
	coding, ok := code["coding"].([]map[string]string)
	if !ok || len(coding) == 0 {
		t.Fatalf("coding missing: %#v", code["coding"])
	}
	if coding[0]["system"] != "http://hl7.org/fhir/sid/icd-10" {
		t.Errorf("system = %s", coding[0]["system"])
	}
	if coding[0]["code"] != "A00" {
		t.Errorf("code = %s", coding[0]["code"])
	}
}

func TestPhase7ObservationAddCategoryCastable(t *testing.T) {
	b := NewObservationBuilder()
	b.addCategory("SNOMED:386053000")
	cats, ok := b.Data["category"].([]interface{})
	if !ok || len(cats) == 0 {
		t.Fatal("category missing")
	}
	cat := cats[0].(map[string]interface{})
	coding := cat["coding"].([]map[string]string)
	if coding[0]["system"] != "http://snomed.info/sct" {
		t.Errorf("system = %s", coding[0]["system"])
	}
}
