package queue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type cachedOAuthToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ExpiresAt   time.Time
}

// Worker processes pending jobs from a QueueInterface using an HTTP client.
type Worker struct {
	queue        QueueInterface
	oauth2Config map[string]string
	tokenCache   *cachedOAuthToken
	processed    int
	succeeded    int
	failed       int
	dlq          int
	rateLimiter  *RateLimiter
	httpClient   *http.Client
}

// NewWorker creates a new queue Worker.
func NewWorker(
	queue QueueInterface,
	oauth2Config map[string]string,
	rateLimiter *RateLimiter,
	httpClients ...*http.Client,
) *Worker {
	if rateLimiter == nil {
		rateLimiter = NewRateLimiter(300)
	}
	var client *http.Client
	if len(httpClients) > 0 && httpClients[0] != nil {
		client = httpClients[0]
	} else {
		client = &http.Client{Timeout: 30 * time.Second}
	}
	return &Worker{
		queue:        queue,
		oauth2Config: oauth2Config,
		rateLimiter:  rateLimiter,
		httpClient:   client,
	}
}

// Process dequeues and processes up to limit pending jobs (default 50).
func (w *Worker) Process(limit ...int) map[string]int {
	maxJobs := 50
	if len(limit) > 0 && limit[0] > 0 {
		maxJobs = limit[0]
	}

	w.processed = 0
	w.succeeded = 0
	w.failed = 0
	w.dlq = 0

	for i := 0; i < maxJobs; i++ {
		job := w.queue.Dequeue()
		if job == nil {
			break
		}

		res := w.HandleJob(job)
		w.processed++

		if res.IsSuccess() {
			w.succeeded++
		} else if res.IsDlq() {
			w.dlq++
		} else {
			w.failed++
		}
	}

	return map[string]int{
		"processed": w.processed,
		"succeeded": w.succeeded,
		"failed":    w.failed,
		"dlq":       w.dlq,
	}
}

// HandleJob handles a single job end-to-end.
func (w *Worker) HandleJob(job map[string]any) *QueueResult {
	jobID := intFromAny(job["id"])
	attempts := intFromAny(job["attempts"])

	w.rateLimiter.Wait()
	respMap, err := w.send(job)
	if err != nil {
		w.queue.MarkFailed(jobID, err.Error(), map[string]any{})
		return NewFailedResult(jobID, err.Error(), map[string]any{}, attempts+1)
	}

	httpCode := intFromAny(respMap["http_code"])
	bodyStr, _ := respMap["body"].(string)
	location, _ := respMap["location"].(string)

	var parsed any = bodyStr
	if bodyStr != "" {
		var m map[string]any
		if jsonErr := json.Unmarshal([]byte(bodyStr), &m); jsonErr == nil && m != nil {
			parsed = m
		}
	}

	classified := ClassifyWithBody(httpCode, parsed)

	// 2xx -> success
	if classified.Category == CatSuccess {
		w.queue.MarkSuccess(jobID, respMap, location)
		return NewSuccessResult(jobID, respMap, attempts+1)
	}

	// 401 -> invalidate token cache and let MarkFailed requeue
	if classified.Category == CatUnauthorized {
		w.tokenCache = nil
	}

	// 429 with Retry-After -> honor it before requeue
	if classified.Category == CatRateLimited {
		retryAfterHeader, _ := respMap["retry_after"].(string)
		if retryAfter := ParseRetryAfter(retryAfterHeader); retryAfter != nil && *retryAfter > 0 {
			w.queue.MarkFailed(jobID, classified.Detail, respMap)
			if sq, ok := w.queue.(*SqliteQueue); ok {
				sq.SetScheduledAt(jobID, time.Now().UTC().Add(time.Duration(*retryAfter)*time.Second))
			}
			return NewFailedResult(jobID, classified.Detail, respMap, attempts+1)
		}
	}

	// Non-retryable -> DLQ
	if !classified.Retryable {
		w.queue.MarkDlq(jobID, classified.Detail, respMap)
		return NewDlqResult(jobID, classified.Detail, respMap, attempts+1)
	}

	// Retryable -> failed (SqliteQueue handles backoff scheduling)
	w.queue.MarkFailed(jobID, classified.Detail, respMap)
	return NewFailedResult(jobID, classified.Detail, respMap, attempts+1)
}

func (w *Worker) send(job map[string]any) (map[string]any, error) {
	token, err := w.getToken()
	if err != nil {
		return nil, err
	}

	method, _ := job["method"].(string)
	if method == "" {
		method = "POST"
	}
	method = strings.ToUpper(method)

	fhirURL := strings.TrimRight(w.oauth2Config["fhir_url"], "/")
	jobURL := strings.TrimLeft(fmt.Sprintf("%v", job["url"]), "/")
	targetURL := fhirURL + "/" + jobURL

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/fhir+json",
		"Accept":        "application/fhir+json",
	}

	if method == "POST" {
		if idem, ok := job["idempotency_key"].(string); ok && idem != "" {
			headers["If-None-Exist"] = "uuid=" + idem
		}
	}
	if method == "PUT" || method == "PATCH" {
		if etag, ok := job["etag"].(string); ok && etag != "" {
			headers["If-Match"] = etag
		}
	}

	var bodyReader io.Reader
	if payload := job["bundle_payload"]; payload != nil && (method == "POST" || method == "PUT" || method == "PATCH") {
		rawBytes, marshalErr := json.Marshal(payload)
		if marshalErr != nil {
			return nil, marshalErr
		}
		bodyReader = bytes.NewReader(rawBytes)
	}

	req, err := http.NewRequest(method, targetURL, bodyReader)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := w.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP error: %w", err)
	}
	defer resp.Body.Close()

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyStr := string(rawBody)
	retryAfter := strings.TrimSpace(resp.Header.Get("Retry-After"))
	location := extractJobLocation(bodyStr)
	if location == "" {
		location = resp.Header.Get("Location")
	}

	var parsedResp any
	if bodyStr != "" {
		_ = json.Unmarshal(rawBody, &parsedResp)
	}

	return map[string]any{
		"body":        bodyStr,
		"http_code":   resp.StatusCode,
		"location":    location,
		"retry_after": retryAfter,
		"response":    parsedResp,
	}, nil
}

func (w *Worker) getToken() (string, error) {
	if w.tokenCache != nil && time.Now().Before(w.tokenCache.ExpiresAt.Add(-60*time.Second)) {
		return w.tokenCache.AccessToken, nil
	}
	tok, err := w.fetchToken()
	if err != nil {
		return "", err
	}
	w.tokenCache = tok
	return tok.AccessToken, nil
}

func (w *Worker) fetchToken() (*cachedOAuthToken, error) {
	tokenURL := strings.TrimRight(w.oauth2Config["base_url"], "/") + "/oauth2/v1/accesstoken"
	form := url.Values{}
	form.Set("grant_type", "client_credentials")
	form.Set("client_id", w.oauth2Config["client_id"])
	form.Set("client_secret", w.oauth2Config["client_secret"])

	req, err := http.NewRequest(http.MethodPost, tokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := w.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rawBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 || len(rawBody) == 0 {
		return nil, fmt.Errorf("OAuth2 token fetch failed (HTTP %d): %s", resp.StatusCode, string(rawBody))
	}

	var data cachedOAuthToken
	if err := json.Unmarshal(rawBody, &data); err != nil || data.AccessToken == "" {
		return nil, fmt.Errorf("Invalid OAuth2 token response: %s", string(rawBody))
	}

	expiresIn := data.ExpiresIn
	if expiresIn <= 0 {
		expiresIn = 3600
	}
	data.ExpiresAt = time.Now().Add(time.Duration(expiresIn) * time.Second)
	return &data, nil
}

func extractJobLocation(body string) string {
	if body == "" {
		return ""
	}
	var data map[string]any
	if err := json.Unmarshal([]byte(body), &data); err != nil || data == nil {
		return ""
	}
	if fullURL, ok := data["fullUrl"].(string); ok && fullURL != "" {
		return fullURL
	}
	if res, ok := data["resource"].(map[string]any); ok {
		if id, ok := res["id"].(string); ok && id != "" {
			rt, _ := res["resourceType"].(string)
			return rt + "/" + id
		}
	}
	return ""
}
