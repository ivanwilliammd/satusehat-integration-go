package queue

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Error category constants matching PHP ErrorClassifier.
const (
	CatSuccess         = "success"
	CatInvalidRequest  = "invalid_request"
	CatUnauthorized    = "unauthorized"
	CatForbidden       = "forbidden"
	CatNotFound        = "not_found"
	CatConflict        = "conflict"
	CatValidationError = "validation_error"
	CatRateLimited     = "rate_limited"
	CatServerError     = "server_error"
	CatNetworkError    = "network_error"
	CatUnexpected      = "unexpected"

	CAT_SUCCESS          = CatSuccess
	CAT_INVALID_REQUEST  = CatInvalidRequest
	CAT_UNAUTHORIZED     = CatUnauthorized
	CAT_FORBIDDEN        = CatForbidden
	CAT_NOT_FOUND        = CatNotFound
	CAT_CONFLICT         = CatConflict
	CAT_VALIDATION_ERROR = CatValidationError
	CAT_RATE_LIMITED     = CatRateLimited
	CAT_SERVER_ERROR     = CatServerError
	CAT_NETWORK_ERROR    = CatNetworkError
	CAT_UNEXPECTED       = CatUnexpected
)

// ErrorClassification represents the result of classifying an HTTP response.
type ErrorClassification struct {
	Category  string
	Retryable bool
	Status    string
	Detail    string
}

// ToMap converts ErrorClassification to a map matching PHP's associative array.
func (e ErrorClassification) ToMap() map[string]any {
	return map[string]any{
		"category":  e.Category,
		"retryable": e.Retryable,
		"status":    e.Status,
		"detail":    e.Detail,
	}
}

// ErrorClassifier provides static-like classification methods.
type ErrorClassifier struct{}

// Classify classifies an HTTP status code and optional response body into an ErrorClassification.
func Classify(httpCode int, responseBody ...any) ErrorClassification {
	var body any
	if len(responseBody) > 0 {
		body = responseBody[0]
	}
	return ClassifyWithBody(httpCode, body)
}

// ClassifyWithBody classifies an HTTP status code and response body.
func ClassifyWithBody(httpCode int, responseBody any) ErrorClassification {
	// 2xx — always success
	if httpCode >= 200 && httpCode < 300 {
		return ErrorClassification{
			Category:  CatSuccess,
			Retryable: false,
			Status:    StatusSuccess,
			Detail:    fmt.Sprintf("HTTP %d", httpCode),
		}
	}

	if httpCode == 400 {
		detail := extractDetail(responseBody)
		if detail == "" {
			detail = "Bad request (HTTP 400)"
		}
		return ErrorClassification{
			Category:  CatInvalidRequest,
			Retryable: false,
			Status:    StatusDlq,
			Detail:    detail,
		}
	}

	if httpCode == 401 {
		detail := extractDetail(responseBody)
		if detail == "" {
			detail = "Unauthorized (HTTP 401)"
		}
		return ErrorClassification{
			Category:  CatUnauthorized,
			Retryable: true,
			Status:    StatusPending,
			Detail:    detail,
		}
	}

	if httpCode == 403 {
		detail := extractDetail(responseBody)
		if detail == "" {
			detail = "Forbidden (HTTP 403)"
		}
		return ErrorClassification{
			Category:  CatForbidden,
			Retryable: false,
			Status:    StatusDlq,
			Detail:    detail,
		}
	}

	if httpCode == 404 {
		detail := extractDetail(responseBody)
		if detail == "" {
			detail = "Resource not found (HTTP 404)"
		}
		return ErrorClassification{
			Category:  CatNotFound,
			Retryable: false,
			Status:    StatusDlq,
			Detail:    detail,
		}
	}

	if httpCode == 408 {
		return ErrorClassification{
			Category:  CatNetworkError,
			Retryable: true,
			Status:    StatusPending,
			Detail:    "Request timeout (HTTP 408)",
		}
	}

	if httpCode == 409 {
		detail := extractDetail(responseBody)
		if detail == "" {
			detail = "Conflict (HTTP 409)"
		}
		return ErrorClassification{
			Category:  CatConflict,
			Retryable: false,
			Status:    StatusDlq,
			Detail:    detail,
		}
	}

	if httpCode == 412 {
		return ErrorClassification{
			Category:  CatConflict,
			Retryable: false,
			Status:    StatusDlq,
			Detail:    "Precondition failed (HTTP 412) — resource changed",
		}
	}

	if httpCode == 422 {
		detail := extractFhirIssues(responseBody)
		if detail == "" {
			detail = extractDetail(responseBody)
		}
		if detail == "" {
			detail = "Validation error (HTTP 422)"
		}
		return ErrorClassification{
			Category:  CatValidationError,
			Retryable: false,
			Status:    StatusDlq,
			Detail:    detail,
		}
	}

	if httpCode == 429 {
		return ErrorClassification{
			Category:  CatRateLimited,
			Retryable: true,
			Status:    StatusPending,
			Detail:    "Rate limited (HTTP 429)",
		}
	}

	if httpCode >= 500 && httpCode < 600 {
		return ErrorClassification{
			Category:  CatServerError,
			Retryable: true,
			Status:    StatusPending,
			Detail:    fmt.Sprintf("Server error (HTTP %d)", httpCode),
		}
	}

	if httpCode >= 400 && httpCode < 500 {
		detail := extractDetail(responseBody)
		if detail == "" {
			detail = fmt.Sprintf("Client error (HTTP %d)", httpCode)
		}
		return ErrorClassification{
			Category:  CatInvalidRequest,
			Retryable: false,
			Status:    StatusDlq,
			Detail:    detail,
		}
	}

	return ErrorClassification{
		Category:  CatUnexpected,
		Retryable: false,
		Status:    StatusDlq,
		Detail:    fmt.Sprintf("Unexpected HTTP %d", httpCode),
	}
}

func extractDetail(body any) string {
	if body == nil {
		return ""
	}
	if m, ok := body.(map[string]any); ok {
		if issues := toMapSliceQueue(m["issue"]); len(issues) > 0 {
			if diag, ok := issues[0]["diagnostics"].(string); ok && diag != "" {
				return truncateStr(diag, 200)
			}
		}
		if msg, ok := m["message"].(string); ok && msg != "" {
			return truncateStr(msg, 200)
		}
		if errDesc, ok := m["error_description"].(string); ok && errDesc != "" {
			return truncateStr(errDesc, 200)
		}
	}
	if s, ok := body.(string); ok && s != "" {
		var data map[string]any
		if err := json.Unmarshal([]byte(s), &data); err == nil && data != nil {
			return extractDetail(data)
		}
		return truncateStr(s, 200)
	}
	return ""
}

func extractFhirIssues(body any) string {
	var m map[string]any
	switch v := body.(type) {
	case map[string]any:
		m = v
	case string:
		if v != "" {
			_ = json.Unmarshal([]byte(v), &m)
		}
	}
	if m == nil {
		return ""
	}
	issues := toMapSliceQueue(m["issue"])
	if len(issues) == 0 {
		return ""
	}
	limit := len(issues)
	if limit > 5 {
		limit = 5
	}
	lines := make([]string, 0, limit)
	for _, issue := range issues[:limit] {
		msg, _ := issue["diagnostics"].(string)
		if msg == "" {
			msg, _ = issue["location"].(string)
		}
		if msg == "" {
			msg = "Unknown"
		}
		sev, _ := issue["severity"].(string)
		if sev == "" {
			sev = "ERROR"
		} else {
			sev = strings.ToUpper(sev)
		}
		lines = append(lines, fmt.Sprintf("[%s] %s", sev, truncateStr(msg, 100)))
	}
	return strings.Join(lines, " | ")
}

func truncateStr(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen]
}

func toMapSliceQueue(val any) []map[string]any {
	if val == nil {
		return nil
	}
	if s, ok := val.([]map[string]any); ok {
		return s
	}
	if s, ok := val.([]any); ok {
		out := make([]map[string]any, 0, len(s))
		for _, item := range s {
			if m, ok := item.(map[string]any); ok {
				out = append(out, m)
			}
		}
		return out
	}
	return nil
}
