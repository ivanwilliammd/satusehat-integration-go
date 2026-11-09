package ssresponse

import (
	"testing"
)

func TestSSResponse_Methods(t *testing.T) {
	r := NewSSResponse(201, map[string]any{
		"resourceType": "Encounter",
		"id":           "enc-123",
	}, `{"resourceType":"Encounter","id":"enc-123"}`)

	if !r.IsSuccess() || r.IsError() {
		t.Fatal("expected 201 to be success")
	}
	if r.GetResourceId() != "enc-123" || r.Get("Encounter") != "enc-123" {
		t.Fatalf("unexpected resource id: %s", r.GetResourceId())
	}

	errResp := NewSSResponse(422, map[string]any{
		"resourceType": "OperationOutcome",
		"issue": []any{
			map[string]any{
				"severity": "error",
				"code":     "invalid",
				"details":  map[string]any{"text": "RuleNumber: 10003 - Invalid Condition.code"},
				"expression": []any{
					"Condition.code",
				},
			},
		},
	}, `{}`)

	if !errResp.IsError() || errResp.IsSuccess() {
		t.Fatal("expected 422 to be error")
	}
	msgs := errResp.GetErrorMessages()
	if len(msgs) != 1 || msgs[0] != "RuleNumber: 10003 - Invalid Condition.code" {
		t.Fatalf("unexpected error messages: %v", msgs)
	}
}

func TestBundleResponse_Methods(t *testing.T) {
	bundle := map[string]any{
		"resourceType": "Bundle",
		"type":         "transaction-response",
		"total":        2,
		"entry": []any{
			map[string]any{
				"fullUrl": "urn:uuid:1111",
				"response": map[string]any{
					"status":   "201 Created",
					"location": "https://api-satusehat-dev.dto.kemkes.go.id/fhir-r4/v1/Encounter/enc-999/_history/1",
					"etag":     "W/\"1\"",
				},
			},
			map[string]any{
				"fullUrl": "urn:uuid:2222",
				"response": map[string]any{
					"status":   "200 OK",
					"location": "https://api-satusehat-dev.dto.kemkes.go.id/fhir-r4/v1/Patient/pat-888",
				},
				"resource": map[string]any{
					"resourceType": "Patient",
					"id":           "pat-888",
				},
			},
		},
	}

	reqEntries := []map[string]any{
		{"method": "POST", "url": "Encounter", "fullUrl": "urn:uuid:1111"},
		{"method": "GET", "url": "Patient/pat-888", "fullUrl": "urn:uuid:2222"},
	}

	br := NewBundleResponse(200, bundle, reqEntries)
	if !br.IsSuccess() || !br.IsAllSuccess() {
		t.Fatal("expected bundle response to be all success")
	}
	if br.GetTotal() != 2 {
		t.Fatalf("expected total 2, got %d", br.GetTotal())
	}
	if len(br.GetSuccessEntries()) != 2 || len(br.GetFailedEntries()) != 0 {
		t.Fatal("unexpected success/failed entry counts")
	}
	if br.GetEntryResourceId(0) != "_history" && br.GetEntryResourceId(1) != "pat-888" {
		t.Fatalf("unexpected entry resource id: %s", br.GetEntryResourceId(1))
	}

	created := br.Created()
	if len(created) != 1 || created[0]["fullUrl"] != "urn:uuid:1111" {
		t.Fatalf("unexpected created entries: %v", created)
	}
	read := br.Read()
	if len(read) != 1 || read[0]["resource"] == nil {
		t.Fatalf("unexpected read entries: %v", read)
	}
	idMap := br.IdMap()
	if idMap["urn:uuid:1111"] == "" {
		t.Fatalf("expected urn:uuid:1111 in IdMap, got %v", idMap)
	}
	found := br.Find("Patient", "pat-888")
	if found == nil || found["resourceId"] != "pat-888" {
		t.Fatalf("expected Find to locate Patient/pat-888, got %v", found)
	}
}

func TestSSValidationError_FullRuleMapAndMethods(t *testing.T) {
	if len(RuleMap) != 558 {
		t.Fatalf("expected 558 rules in RuleMap, got %d", len(RuleMap))
	}
	if len(ResponseTypes) != 14 {
		t.Fatalf("expected 14 entries in ResponseTypes, got %d", len(ResponseTypes))
	}

	avail := GetAvailableRuleNumbers()
	if len(avail) != 558 || avail[0] != 10003 {
		t.Fatalf("unexpected available rule numbers: count=%d first=%d", len(avail), avail[0])
	}

	condRules := GetRulesForResource("Condition")
	if len(condRules) == 0 || condRules[10003]["path"] != "Condition.code" {
		t.Fatalf("unexpected Condition rules: %v", condRules[10003])
	}

	loincRules := SearchRules("loinc")
	if len(loincRules) == 0 {
		t.Fatal("expected SearchRules('loinc') to return rules")
	}

	rt := GetResponseType("duplicate")
	if rt == nil || rt["title"] != "Found duplicate resource" {
		t.Fatalf("unexpected duplicate response type: %v", rt)
	}

	resp := NewSSResponse(400, map[string]any{
		"resourceType": "OperationOutcome",
		"issue": []any{
			map[string]any{
				"severity":   "error",
				"code":       "duplicate",
				"details":    map[string]any{"text": "RuleNumber: 10003 Found duplicate resource with invalid date time format"},
				"expression": []any{"Condition.code"},
			},
		},
	}, `{}`)

	valErr := FromResponse(resp)
	nums := valErr.GetRuleNumbers()
	if len(nums) != 1 || nums[0] != 10003 {
		t.Fatalf("expected rule number [10003], got %v", nums)
	}
	if !valErr.IsDuplicate() {
		t.Fatal("expected IsDuplicate() to be true")
	}
	if !valErr.IsDateTimeError() {
		t.Fatal("expected IsDateTimeError() to be true")
	}
	paths := valErr.GetAffectedPaths()
	if len(paths) != 1 || paths[0] != "Condition.code" {
		t.Fatalf("unexpected affected paths: %v", paths)
	}
	summary := valErr.GetSummary()
	if summary == "" {
		t.Fatal("expected non-empty summary")
	}
}
