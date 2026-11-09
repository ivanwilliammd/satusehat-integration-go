package ssrequest

import (
	"bytes"
	"encoding/json"
	"io"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/ssresponse"
)

// SSRequestConfig holds optional configuration for SSRequest.
type SSRequestConfig struct {
	Timeout    time.Duration
	MaxRetries *int
	Headers    map[string]string
	HTTPClient *http.Client
	SleepFunc  func(time.Duration)
}

// SSRequest is the SATUSEHAT FHIR R4 HTTP client with automatic token refresh and retry backoff.
type SSRequest struct {
	oauth2         *OAuth2Client
	token          string
	baseURL        string
	defaultHeaders map[string]string
	timeout        time.Duration
	maxRetries     int
	retryCount     int
	httpClient     *http.Client
	sleepFunc      func(time.Duration)
}

// NewSSRequest creates a new SSRequest instance.
func NewSSRequest(oauth2 *OAuth2Client, token string, baseURL string, configs ...SSRequestConfig) *SSRequest {
	var cfg SSRequestConfig
	if len(configs) > 0 {
		cfg = configs[0]
	}

	timeout := cfg.Timeout
	if timeout <= 0 {
		timeout = 30 * time.Second
	}

	maxRetries := 5
	if cfg.MaxRetries != nil {
		maxRetries = *cfg.MaxRetries
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/fhir+json",
	}
	for k, v := range cfg.Headers {
		headers[k] = v
	}

	resolvedBaseURL := ""
	if baseURL != "" {
		resolvedBaseURL = strings.TrimRight(baseURL, "/")
	} else if oauth2 != nil {
		resolvedBaseURL = strings.TrimRight(oauth2.FhirURL, "/")
	}

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: timeout}
	}

	sleepFn := cfg.SleepFunc
	if sleepFn == nil {
		sleepFn = time.Sleep
	}

	return &SSRequest{
		oauth2:         oauth2,
		token:          token,
		baseURL:        resolvedBaseURL,
		defaultHeaders: headers,
		timeout:        timeout,
		maxRetries:     maxRetries,
		httpClient:     httpClient,
		sleepFunc:      sleepFn,
	}
}

// Get sends an HTTP GET request to a FHIR resource endpoint with optional query parameters.
func (s *SSRequest) Get(path string, params ...map[string]string) *ssresponse.SSResponse {
	var p map[string]string
	if len(params) > 0 {
		p = params[0]
	}
	fullURL := s.buildURL(path, p)
	return s.request(http.MethodGet, fullURL, nil, "")
}

// Post sends an HTTP POST request with a JSON body.
func (s *SSRequest) Post(path string, body map[string]any) *ssresponse.SSResponse {
	fullURL := s.buildURL(path, nil)
	return s.request(http.MethodPost, fullURL, body, "")
}

// Put sends an HTTP PUT request with a JSON body.
func (s *SSRequest) Put(path string, body map[string]any) *ssresponse.SSResponse {
	fullURL := s.buildURL(path, nil)
	return s.request(http.MethodPut, fullURL, body, "")
}

// Delete sends an HTTP DELETE request with an optional JSON body.
func (s *SSRequest) Delete(path string, body ...map[string]any) *ssresponse.SSResponse {
	fullURL := s.buildURL(path, nil)
	var b any
	if len(body) > 0 && len(body[0]) > 0 {
		b = body[0]
	}
	return s.request(http.MethodDelete, fullURL, b, "")
}

// Patch sends an HTTP PATCH request with Content-Type: application/json-patch+json.
func (s *SSRequest) Patch(path string, patchPayload any) *ssresponse.SSResponse {
	fullURL := s.buildURL(path, nil)
	return s.request(http.MethodPatch, fullURL, patchPayload, "application/json-patch+json")
}

// PatchFhirPath sends an HTTP PATCH request with Content-Type: application/fhir+json.
func (s *SSRequest) PatchFhirPath(path string, fhirPathPayload any) *ssresponse.SSResponse {
	fullURL := s.buildURL(path, nil)
	s.retryCount = 0
	return s.doRequest(http.MethodPatch, fullURL, fhirPathPayload, "application/fhir+json")
}

// PostBundle sends a FHIR Bundle transaction/batch payload via HTTP POST.
func (s *SSRequest) PostBundle(path string, bundlePayload map[string]any) *ssresponse.SSResponse {
	return s.Post(path, bundlePayload)
}

func (s *SSRequest) getToken() string {
	if s.token != "" {
		return s.token
	}
	if s.oauth2 != nil {
		tok, _ := s.oauth2.Token()
		return tok
	}
	return ""
}

func (s *SSRequest) isTokenAutoRefresh() bool {
	if s.oauth2 != nil {
		return s.oauth2.IsTokenAutoRefresh()
	}
	return false
}

func (s *SSRequest) request(method, fullURL string, body any, contentType string) *ssresponse.SSResponse {
	s.retryCount = 0
	return s.doRequest(method, fullURL, body, contentType)
}

func (s *SSRequest) doRequest(method, fullURL string, body any, contentType string) *ssresponse.SSResponse {
	tok := s.getToken()

	headers := make(map[string]string, len(s.defaultHeaders)+2)
	for k, v := range s.defaultHeaders {
		headers[k] = v
	}
	if tok != "" {
		headers["Authorization"] = "Bearer " + tok
	}

	if method == http.MethodPatch {
		if contentType != "" {
			headers["Content-Type"] = contentType
		} else {
			headers["Content-Type"] = "application/json-patch+json"
		}
	} else if contentType != "" {
		headers["Content-Type"] = contentType
	}

	var bodyReader io.Reader
	if body != nil && (method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch || method == http.MethodDelete) {
		rawBody, err := json.Marshal(body)
		if err != nil {
			return ssresponse.NewSSResponse(
				0,
				map[string]any{
					"resourceType": "OperationOutcome",
					"issue":        []any{map[string]any{"details": map[string]any{"text": err.Error()}}},
				},
				err.Error(),
			)
		}
		bodyReader = bytes.NewReader(rawBody)
	}

	req, err := http.NewRequest(method, fullURL, bodyReader)
	if err != nil {
		return ssresponse.NewSSResponse(
			0,
			map[string]any{
				"resourceType": "OperationOutcome",
				"issue":        []any{map[string]any{"details": map[string]any{"text": err.Error()}}},
			},
			err.Error(),
		)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return ssresponse.NewSSResponse(
			0,
			map[string]any{
				"resourceType": "OperationOutcome",
				"issue":        []any{map[string]any{"details": map[string]any{"text": err.Error()}}},
			},
			err.Error(),
		)
	}
	defer resp.Body.Close()

	rawBytes, _ := io.ReadAll(resp.Body)
	content := string(rawBytes)
	var bodyMap map[string]any
	if content != "" {
		_ = json.Unmarshal(rawBytes, &bodyMap)
	}
	if bodyMap == nil {
		bodyMap = map[string]any{}
	}

	if resp.StatusCode >= 400 {
		return s.handleFailure(method, fullURL, resp, bodyMap, content, body, contentType)
	}

	return ssresponse.NewSSResponse(resp.StatusCode, bodyMap, content)
}

func (s *SSRequest) handleFailure(
	method, fullURL string,
	resp *http.Response,
	bodyMap map[string]any,
	content string,
	reqBody any,
	contentType string,
) *ssresponse.SSResponse {
	statusCode := resp.StatusCode

	// 429 — rate limited: read Retry-After header
	if statusCode == http.StatusTooManyRequests {
		s.retryCount++
		if s.retryCount > s.maxRetries {
			return ssresponse.NewSSResponse(statusCode, bodyMap, content)
		}

		retryAfter := strings.TrimSpace(resp.Header.Get("Retry-After"))
		sleepSeconds := int(math.Pow(2, float64(s.retryCount)))
		if retryAfter != "" {
			if n, err := strconv.Atoi(retryAfter); err == nil {
				sleepSeconds = n
			}
		}

		s.sleep(sleepSeconds)
		return s.doRequest(method, fullURL, reqBody, contentType)
	}

	// 401 — unauthorized: attempt token refresh
	if statusCode == http.StatusUnauthorized {
		if s.isTokenAutoRefresh() && s.oauth2 != nil {
			s.retryCount++
			if s.retryCount > s.maxRetries {
				return ssresponse.NewSSResponse(statusCode, bodyMap, content)
			}

			_, _ = s.oauth2.RefreshToken()
			s.token = ""

			s.sleep(int(math.Pow(2, float64(s.retryCount))))
			return s.doRequest(method, fullURL, reqBody, contentType)
		}
		return ssresponse.NewSSResponse(statusCode, bodyMap, content)
	}

	// 5xx — server error: exponential backoff
	if statusCode >= 500 {
		s.retryCount++
		if s.retryCount > s.maxRetries {
			return ssresponse.NewSSResponse(statusCode, bodyMap, content)
		}

		s.sleep(int(math.Pow(2, float64(s.retryCount))))
		return s.doRequest(method, fullURL, reqBody, contentType)
	}

	return ssresponse.NewSSResponse(statusCode, bodyMap, content)
}

func (s *SSRequest) buildURL(path string, params map[string]string) string {
	var full string
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		full = path
	} else {
		full = s.baseURL + "/" + strings.TrimLeft(path, "/")
	}

	if len(params) == 0 {
		return full
	}

	q := url.Values{}
	for k, v := range params {
		q.Set(k, v)
	}
	sep := "?"
	if strings.Contains(full, "?") {
		sep = "&"
	}
	return full + sep + q.Encode()
}

func (s *SSRequest) sleep(seconds int) {
	if seconds <= 0 {
		return
	}
	if seconds > 60 {
		seconds = 60
	}
	if s.sleepFunc != nil {
		s.sleepFunc(time.Duration(seconds) * time.Second)
	}
}
