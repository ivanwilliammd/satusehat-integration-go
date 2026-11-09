package queue

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestErrorClassifier_AllCategories(t *testing.T) {
	if c := Classify(200); c.Category != CatSuccess || c.Retryable {
		t.Fatalf("unexpected 200 classification: %+v", c)
	}
	if c := Classify(400); c.Category != CatInvalidRequest || c.Retryable || c.Status != StatusDlq {
		t.Fatalf("unexpected 400 classification: %+v", c)
	}
	if c := Classify(401); c.Category != CatUnauthorized || !c.Retryable || c.Status != StatusPending {
		t.Fatalf("unexpected 401 classification: %+v", c)
	}
	if c := Classify(403); c.Category != CatForbidden || c.Retryable {
		t.Fatalf("unexpected 403 classification: %+v", c)
	}
	if c := Classify(404); c.Category != CatNotFound || c.Retryable {
		t.Fatalf("unexpected 404 classification: %+v", c)
	}
	if c := Classify(408); c.Category != CatNetworkError || !c.Retryable {
		t.Fatalf("unexpected 408 classification: %+v", c)
	}
	if c := Classify(409); c.Category != CatConflict || c.Retryable {
		t.Fatalf("unexpected 409 classification: %+v", c)
	}
	if c := Classify(412); c.Category != CatConflict || c.Retryable {
		t.Fatalf("unexpected 412 classification: %+v", c)
	}
	if c := ClassifyWithBody(422, map[string]any{
		"issue": []any{map[string]any{"severity": "error", "diagnostics": "Missing subject"}},
	}); c.Category != CatValidationError || c.Detail != "[ERROR] Missing subject" {
		t.Fatalf("unexpected 422 classification: %+v", c)
	}
	if c := Classify(429); c.Category != CatRateLimited || !c.Retryable {
		t.Fatalf("unexpected 429 classification: %+v", c)
	}
	if c := Classify(502); c.Category != CatServerError || !c.Retryable {
		t.Fatalf("unexpected 502 classification: %+v", c)
	}
}

func TestSqliteQueue_LifecycleAndPersistence(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "queue.json")
	q := NewSqliteQueue(dbPath)

	j1 := q.Enqueue("POST", "Patient", "Patient", map[string]any{"resourceType": "Patient"}, "", "idem-1", "u1", 2, map[string]any{"source": "unit-test"})
	if intFromAny(j1["id"]) != 1 || j1["uuid"] != "idem-1" {
		t.Fatalf("unexpected enqueued job 1: %v", j1)
	}

	j2 := q.EnqueueBundle(map[string]any{"resourceType": "Bundle"}, "transaction", "u1", 3, nil)
	if intFromAny(j2["id"]) != 2 || j2["resource_type"] != "Bundle" {
		t.Fatalf("unexpected enqueued bundle job 2: %v", j2)
	}

	if q.PendingCount() != 2 {
		t.Fatalf("expected 2 pending jobs, got %d", q.PendingCount())
	}

	// Reload from disk to verify durability
	q2 := NewSqliteQueue(dbPath)
	if q2.PendingCount() != 2 {
		t.Fatalf("expected 2 pending jobs after disk reload, got %d", q2.PendingCount())
	}

	deq1 := q2.Dequeue()
	if deq1 == nil || intFromAny(deq1["id"]) != 1 {
		t.Fatalf("expected job 1 dequeued, got %v", deq1)
	}
	q2.MarkSuccess(1, map[string]any{"status": "ok"}, "Patient/P1")

	deq2 := q2.Dequeue()
	if deq2 == nil || intFromAny(deq2["id"]) != 2 {
		t.Fatalf("expected job 2 dequeued, got %v", deq2)
	}
	q2.MarkDlq(2, "Validation failed", map[string]any{})

	if q2.DlqCount() != 1 {
		t.Fatalf("expected 1 DLQ job, got %d", q2.DlqCount())
	}

	mon := NewQueueMonitor(q2)
	sum := mon.Summary()
	if intFromAny(sum["total"]) != 2 || intFromAny(sum["dlq"]) != 1 || intFromAny(sum["success"]) != 1 {
		t.Fatalf("unexpected monitor summary: %v", sum)
	}
	hc := mon.HealthCheck()
	if hc["status"] != "critical" {
		t.Fatalf("expected critical health status when dlq >= success, got %v", hc)
	}
	topErr := mon.TopErrors(5)
	if len(topErr) != 1 || topErr[0]["last_error"] != "Validation failed" {
		t.Fatalf("unexpected top errors: %v", topErr)
	}

	resetCount := q2.ResetAll()
	if resetCount != 1 || q2.PendingCount() != 1 {
		t.Fatalf("expected ResetAll to reset 1 job, got %d", resetCount)
	}
}

func TestWorker_ProcessJobs(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/oauth2/v1/accesstoken":
			_ = json.NewEncoder(w).Encode(map[string]any{"access_token": "worker-token", "expires_in": 3600})
		case "/fhir-r4/v1/Patient":
			if r.Header.Get("If-None-Exist") != "uuid=idem-patient-1" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusCreated)
			_ = json.NewEncoder(w).Encode(map[string]any{
				"fullUrl":  "Patient/PAT-OK",
				"resource": map[string]any{"resourceType": "Patient", "id": "PAT-OK"},
			})
		case "/fhir-r4/v1/Invalid":
			w.WriteHeader(http.StatusUnprocessableEntity)
			_ = json.NewEncoder(w).Encode(map[string]any{
				"resourceType": "OperationOutcome",
				"issue":        []any{map[string]any{"severity": "error", "diagnostics": "Invalid payload"}},
			})
		}
	}))
	defer srv.Close()

	q := NewSqliteQueue(":memory:")
	q.Enqueue("POST", "Patient", "Patient", map[string]any{"resourceType": "Patient"}, "", "idem-patient-1", "system", 3, nil)
	q.Enqueue("POST", "Invalid", "Invalid", map[string]any{"resourceType": "Invalid"}, "", "", "system", 3, nil)

	worker := NewWorker(q, map[string]string{
		"base_url":      srv.URL,
		"fhir_url":      srv.URL + "/fhir-r4/v1",
		"client_id":     "cid",
		"client_secret": "csec",
	}, NewRateLimiter(60000), srv.Client())

	res := worker.Process(10)
	if res["processed"] != 2 || res["succeeded"] != 1 || res["dlq"] != 1 {
		t.Fatalf("unexpected worker result: %v", res)
	}
}

func TestRateLimiter_ParseRetryAfter(t *testing.T) {
	if v := ParseRetryAfter("15"); v == nil || *v != 15 {
		t.Fatalf("expected 15s, got %v", v)
	}
	if v := ParseRetryAfter("Wed, 31 Oct 2099 12:00:00 GMT"); v == nil || *v <= 0 {
		t.Fatalf("expected positive seconds for future HTTP-date, got %v", v)
	}
	if v := ParseRetryAfter(""); v != nil {
		t.Fatalf("expected nil for empty Retry-After, got %v", v)
	}
}
