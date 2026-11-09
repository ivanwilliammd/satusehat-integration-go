package ssrequest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestOAuth2Client_EnvValidationAndEndpoints(t *testing.T) {
	t.Setenv("SATUSEHAT_ENV", "")
	if _, err := NewOAuth2Client(); err == nil {
		t.Fatal("expected error when SATUSEHAT_ENV is missing")
	}

	t.Setenv("SATUSEHAT_ENV", "INVALID")
	if _, err := NewOAuth2Client(); err == nil {
		t.Fatal("expected error when SATUSEHAT_ENV is invalid")
	}

	t.Setenv("SATUSEHAT_ENV", "DEV")
	t.Setenv("CLIENTID_DEV", "")
	if _, err := NewOAuth2Client(); err == nil {
		t.Fatal("expected error when DEV credentials are missing")
	}

	t.Setenv("CLIENTID_DEV", "cid-dev")
	t.Setenv("CLIENTSECRET_DEV", "csec-dev")
	t.Setenv("ORGID_DEV", "org-dev")
	client, err := NewOAuth2Client()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if client.BaseURL != "https://api-satusehat-dev.dto.kemkes.go.id" {
		t.Fatalf("unexpected BaseURL: %s", client.BaseURL)
	}
	if client.AuthURL != "https://api-satusehat-dev.dto.kemkes.go.id/oauth2/v1" {
		t.Fatalf("unexpected AuthURL: %s", client.AuthURL)
	}
	if client.FhirURL != "https://api-satusehat-dev.dto.kemkes.go.id/fhir-r4/v1" {
		t.Fatalf("unexpected FhirURL: %s", client.FhirURL)
	}
}

func TestOAuth2Client_TokenAndHelperRequests(t *testing.T) {
	var tokenCount int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/oauth2/v1/accesstoken":
			n := atomic.AddInt32(&tokenCount, 1)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{
				"access_token": "tok-" + string(rune('0'+n)),
			})
		case r.URL.Path == "/fhir-r4/v1/Patient/P123":
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{
				"resourceType": "Patient",
				"id":           "P123",
			})
		case r.URL.Path == "/fhir-r4/v1/Practitioner" && r.URL.RawQuery != "":
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{
				"resourceType": "Bundle",
				"total":        1,
				"entry": []any{
					map[string]any{"resource": map[string]any{"id": "PRAC-1"}},
				},
			})
		case r.URL.Path == "/kfa-v2/products":
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{"result": map[string]any{"kfa_code": "93004944"}})
		case r.URL.Path == "/fhir-r4/v1/Encounter" && r.Method == http.MethodPost:
			w.WriteHeader(http.StatusCreated)
			_ = json.NewEncoder(w).Encode(map[string]any{"resourceType": "Encounter", "id": "ENC-1"})
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer srv.Close()

	c, err := NewOAuth2Client(OAuth2Config{
		Env:            "DEV",
		BaseURL:        srv.URL,
		ClientID:       "cid",
		ClientSecret:   "csec",
		OrganizationID: "org",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tok1, err := c.Token()
	if err != nil || tok1 != "tok-1" {
		t.Fatalf("expected tok-1, got %q (%v)", tok1, err)
	}
	// Cached token
	tok2, err := c.Token()
	if err != nil || tok2 != "tok-1" {
		t.Fatalf("expected cached tok-1, got %q (%v)", tok2, err)
	}
	// Forced refresh
	tok3, err := c.RefreshToken()
	if err != nil || tok3 != "tok-2" {
		t.Fatalf("expected refreshed tok-2, got %q (%v)", tok3, err)
	}

	code, res, err := c.GetByID("Patient", "P123")
	if err != nil || code != 200 || res["id"] != "P123" {
		t.Fatalf("GetByID failed: %d %v %v", code, res, err)
	}

	code, res, err = c.GetByNIK("Practitioner", "1234567890123456")
	if err != nil || code != 200 {
		t.Fatalf("GetByNIK failed: %d %v %v", code, res, err)
	}

	code, res, err = c.KfaGet("products", "?code=93004944")
	if err != nil || code != 200 {
		t.Fatalf("KfaGet failed: %d %v %v", code, res, err)
	}

	code, res, err = c.Post("Encounter", `{"resourceType":"Encounter"}`)
	if err != nil || code != 201 || res["id"] != "ENC-1" {
		t.Fatalf("Post failed: %d %v %v", code, res, err)
	}
}

func TestSSRequest_VerbMethodsAndAutoRetry(t *testing.T) {
	var attempts int32
	var tokenRefreshCount int32

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/oauth2/v1/accesstoken":
			n := atomic.AddInt32(&tokenRefreshCount, 1)
			_ = json.NewEncoder(w).Encode(map[string]any{"access_token": "fresh-tok-" + string(rune('0'+n))})
		case "/fhir-r4/v1/Retry429":
			cur := atomic.AddInt32(&attempts, 1)
			if cur == 1 {
				w.Header().Set("Retry-After", "1")
				w.WriteHeader(http.StatusTooManyRequests)
				_, _ = w.Write([]byte(`{"resourceType":"OperationOutcome"}`))
				return
			}
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"resourceType":"Patient","id":"P-429-OK"}`))
		case "/fhir-r4/v1/Retry401":
			cur := atomic.AddInt32(&attempts, 1)
			if cur == 1 {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte(`{"resourceType":"OperationOutcome"}`))
				return
			}
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"resourceType":"Patient","id":"P-401-OK"}`))
		case "/fhir-r4/v1/Retry503":
			cur := atomic.AddInt32(&attempts, 1)
			if cur <= 2 {
				w.WriteHeader(http.StatusServiceUnavailable)
				_, _ = w.Write([]byte(`{"resourceType":"OperationOutcome"}`))
				return
			}
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"resourceType":"Patient","id":"P-503-OK"}`))
		case "/fhir-r4/v1/Encounter/E1":
			if r.Method == http.MethodPatch {
				ct := r.Header.Get("Content-Type")
				w.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(w).Encode(map[string]any{"id": "E1", "contentType": ct})
				return
			}
			if r.Method == http.MethodPut || r.Method == http.MethodDelete {
				w.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(w).Encode(map[string]any{"id": "E1"})
				return
			}
		}
	}))
	defer srv.Close()

	oauth, err := NewOAuth2Client(OAuth2Config{
		Env:            "DEV",
		BaseURL:        srv.URL,
		ClientID:       "cid",
		ClientSecret:   "csec",
		OrganizationID: "org",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var slept []time.Duration
	req := NewSSRequest(oauth, "", "", SSRequestConfig{
		SleepFunc: func(d time.Duration) {
			slept = append(slept, d)
		},
	})

	// Test 429 retry
	atomic.StoreInt32(&attempts, 0)
	res429 := req.Get("Retry429")
	if !res429.IsSuccess() || res429.GetResourceId() != "P-429-OK" {
		t.Fatalf("expected 429 retry to succeed, got %d %v", res429.GetHttpCode(), res429.GetBody())
	}

	// Test 401 token refresh and retry
	atomic.StoreInt32(&attempts, 0)
	res401 := req.Get("Retry401")
	if !res401.IsSuccess() || res401.GetResourceId() != "P-401-OK" {
		t.Fatalf("expected 401 retry to succeed, got %d %v", res401.GetHttpCode(), res401.GetBody())
	}

	// Test 503 exponential backoff retry
	atomic.StoreInt32(&attempts, 0)
	res503 := req.Get("Retry503")
	if !res503.IsSuccess() || res503.GetResourceId() != "P-503-OK" {
		t.Fatalf("expected 503 retry to succeed, got %d %v", res503.GetHttpCode(), res503.GetBody())
	}

	// Test Patch vs PatchFhirPath content types
	resPatch := req.Patch("Encounter/E1", []map[string]any{{"op": "replace", "path": "/status", "value": "finished"}})
	if resPatch.GetBody()["contentType"] != "application/json-patch+json" {
		t.Fatalf("expected json-patch+json, got %v", resPatch.GetBody()["contentType"])
	}

	resFhirPatch := req.PatchFhirPath("Encounter/E1", map[string]any{"resourceType": "Parameters"})
	if resFhirPatch.GetBody()["contentType"] != "application/fhir+json" {
		t.Fatalf("expected fhir+json, got %v", resFhirPatch.GetBody()["contentType"])
	}

	if len(slept) == 0 {
		t.Fatal("expected sleepFunc to be invoked during retries")
	}
}
