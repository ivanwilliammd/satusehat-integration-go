package ssresponse

import (
	"fmt"
	"strings"
)

// SSResponse wraps an HTTP response from the SATUSEHAT FHIR R4 API.
type SSResponse struct {
	httpCode int
	body     map[string]any
	content  string
}

// NewSSResponse creates a new SSResponse instance.
func NewSSResponse(httpCode int, body map[string]any, content string) *SSResponse {
	if body == nil {
		body = map[string]any{}
	}
	return &SSResponse{
		httpCode: httpCode,
		body:     body,
		content:  content,
	}
}

// GetHttpCode returns the HTTP status code.
func (r *SSResponse) GetHttpCode() int {
	return r.httpCode
}

// GetBody returns the parsed JSON body as a map.
func (r *SSResponse) GetBody() map[string]any {
	return r.body
}

// GetContent returns the raw response body string.
func (r *SSResponse) GetContent() string {
	return r.content
}

// IsSuccess returns true if HTTP status code is 2xx.
func (r *SSResponse) IsSuccess() bool {
	return r.httpCode >= 200 && r.httpCode < 300
}

// IsError returns true if HTTP status code is not 2xx.
func (r *SSResponse) IsError() bool {
	return !r.IsSuccess()
}

// GetErrorMessages extracts FHIR OperationOutcome issue messages from the response body.
func (r *SSResponse) GetErrorMessages() []string {
	messages := []string{}
	if r.body == nil {
		return messages
	}

	rt, _ := r.body["resourceType"].(string)
	if rt != "OperationOutcome" {
		return messages
	}

	issues := toMapSlice(r.body["issue"])
	for _, issue := range issues {
		text := extractIssueText(issue)
		if text != "" {
			messages = append(messages, text)
		}
	}

	return messages
}

// GetResourceId extracts the FHIR resource ID from the response body.
func (r *SSResponse) GetResourceId() string {
	if r.body == nil {
		return ""
	}

	if id, ok := r.body["id"]; ok && id != nil {
		s := fmt.Sprintf("%v", id)
		if s != "" {
			return s
		}
	}

	// Bundle entry link or first entry resource id
	if rt, _ := r.body["resourceType"].(string); rt == "Bundle" {
		if linkMap, ok := r.body["link"].(map[string]any); ok {
			if selfMap, ok := linkMap["self"].(map[string]any); ok {
				if selfURL, ok := selfMap["url"].(string); ok && selfURL != "" {
					if idPart := parseIdFromUrl(selfURL); idPart != "" {
						return idPart
					}
				}
			}
		} else if linkSlice := toMapSlice(r.body["link"]); len(linkSlice) > 0 {
			for _, l := range linkSlice {
				if rel, _ := l["relation"].(string); rel == "self" {
					if selfURL, _ := l["url"].(string); selfURL != "" {
						if idPart := parseIdFromUrl(selfURL); idPart != "" {
							return idPart
						}
					}
				}
			}
		}

		entries := toMapSlice(r.body["entry"])
		if len(entries) > 0 {
			if res, ok := entries[0]["resource"].(map[string]any); ok {
				if id, ok := res["id"]; ok && id != nil {
					s := fmt.Sprintf("%v", id)
					if s != "" {
						return s
					}
				}
			}
		}
	}

	// Patient create response
	if success, ok := r.body["success"].(bool); ok && success {
		if data, ok := r.body["data"].(map[string]any); ok {
			if pid, ok := data["patient_id"]; ok && pid != nil {
				return fmt.Sprintf("%v", pid)
			}
		}
	}

	if cp, ok := r.body["create_patient"].(map[string]any); ok {
		if success, ok := cp["success"].(bool); ok && success {
			if data, ok := cp["data"].(map[string]any); ok {
				if pid, ok := data["patient_id"]; ok && pid != nil {
					return fmt.Sprintf("%v", pid)
				}
			}
		}
	}

	return ""
}

// Get is a shorthand for GetResourceId.
func (r *SSResponse) Get(resourceType string) string {
	_ = resourceType
	return r.GetResourceId()
}

func extractIssueText(issue map[string]any) string {
	if details, ok := issue["details"].(map[string]any); ok {
		if text, ok := details["text"].(string); ok && text != "" {
			return text
		}
	}
	if diag, ok := issue["diagnostics"].(string); ok && diag != "" {
		return diag
	}
	if details, ok := issue["details"].(map[string]any); ok {
		codings := toMapSlice(details["coding"])
		if len(codings) > 0 {
			if disp, ok := codings[0]["display"].(string); ok && disp != "" {
				return disp
			}
		}
	}
	return ""
}

func parseIdFromUrl(rawURL string) string {
	if rawURL == "" {
		return ""
	}
	parts := strings.Split(rawURL, "/")
	idPart := parts[len(parts)-1]
	idPart = strings.Split(idPart, "?")[0]
	return idPart
}

func toMapSlice(val any) []map[string]any {
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
