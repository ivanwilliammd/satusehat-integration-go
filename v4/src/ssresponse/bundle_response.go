package ssresponse

import (
	"fmt"
	"strconv"
	"strings"
)

// BundleResponse parses FHIR Bundle responses (batch-response / transaction-response).
type BundleResponse struct {
	httpCode       int
	bundle         map[string]any
	requestEntries []map[string]any
}

// NewBundleResponse creates a new BundleResponse instance.
func NewBundleResponse(httpCode int, bundle map[string]any, requestEntries ...[]map[string]any) *BundleResponse {
	if bundle == nil {
		bundle = map[string]any{}
	}
	var reqEntries []map[string]any
	if len(requestEntries) > 0 {
		reqEntries = requestEntries[0]
	}
	return &BundleResponse{
		httpCode:       httpCode,
		bundle:         bundle,
		requestEntries: reqEntries,
	}
}

// GetHttpCode returns the HTTP status code.
func (b *BundleResponse) GetHttpCode() int {
	return b.httpCode
}

// GetBundle returns the raw bundle map.
func (b *BundleResponse) GetBundle() map[string]any {
	return b.bundle
}

// IsSuccess returns true if the HTTP status code is 2xx.
func (b *BundleResponse) IsSuccess() bool {
	return b.httpCode >= 200 && b.httpCode < 300
}

// GetTotal returns the total count from the bundle.
func (b *BundleResponse) GetTotal() int {
	if b.bundle == nil {
		return 0
	}
	switch v := b.bundle["total"].(type) {
	case int:
		return v
	case float64:
		return int(v)
	case string:
		n, _ := strconv.Atoi(v)
		return n
	default:
		return 0
	}
}

// GetEntries returns normalized response entries from the bundle.
func (b *BundleResponse) GetEntries() []map[string]any {
	entries := []map[string]any{}
	rawEntries := toMapSlice(b.bundle["entry"])
	for _, entry := range rawEntries {
		resp, _ := entry["response"].(map[string]any)
		if resp == nil {
			resp = map[string]any{}
		}
		status, _ := resp["status"].(string)
		location, _ := resp["location"].(string)
		etag, _ := resp["etag"].(string)
		lastModified, _ := resp["lastModified"].(string)
		outcome := resp["outcome"]
		fullURL, _ := entry["fullUrl"].(string)

		entries = append(entries, map[string]any{
			"status":       status,
			"location":     location,
			"etag":         etag,
			"lastModified": lastModified,
			"outcome":      outcome,
			"fullUrl":      fullURL,
		})
	}
	return entries
}

// GetEntryStatuses returns the status strings for all entries.
func (b *BundleResponse) GetEntryStatuses() []string {
	entries := b.GetEntries()
	statuses := make([]string, 0, len(entries))
	for _, entry := range entries {
		s, _ := entry["status"].(string)
		statuses = append(statuses, s)
	}
	return statuses
}

// GetSuccessEntries returns all entries with a 2xx status code.
func (b *BundleResponse) GetSuccessEntries() []map[string]any {
	success := []map[string]any{}
	for idx, entry := range b.GetEntries() {
		status, _ := entry["status"].(string)
		code := parseStatusCode(status)
		if code >= 200 && code < 300 {
			copied := copyMap(entry)
			copied["index"] = idx
			success = append(success, copied)
		}
	}
	return success
}

// GetFailedEntries returns all entries with a 4xx/5xx or 0 status code.
func (b *BundleResponse) GetFailedEntries() []map[string]any {
	failed := []map[string]any{}
	for idx, entry := range b.GetEntries() {
		status, _ := entry["status"].(string)
		code := parseStatusCode(status)
		if code >= 400 || code == 0 {
			copied := copyMap(entry)
			copied["index"] = idx
			failed = append(failed, copied)
		}
	}
	return failed
}

// GetEntryResourceId extracts the resource ID from the entry at index.
func (b *BundleResponse) GetEntryResourceId(index int) string {
	rawEntries := toMapSlice(b.bundle["entry"])
	if index < 0 || index >= len(rawEntries) {
		return ""
	}
	resp, _ := rawEntries[index]["response"].(map[string]any)
	if resp == nil {
		return ""
	}
	location, _ := resp["location"].(string)
	return parseIdFromUrl(location)
}

// GetEntryErrors extracts OperationOutcome error messages for the entry at index.
func (b *BundleResponse) GetEntryErrors(index int) []string {
	rawEntries := toMapSlice(b.bundle["entry"])
	if index < 0 || index >= len(rawEntries) {
		return []string{}
	}
	resp, _ := rawEntries[index]["response"].(map[string]any)
	if resp == nil {
		return []string{}
	}
	outcome, _ := resp["outcome"].(map[string]any)
	if outcome == nil {
		return []string{}
	}
	if rt, _ := outcome["resourceType"].(string); rt != "OperationOutcome" {
		return []string{}
	}
	messages := []string{}
	for _, issue := range toMapSlice(outcome["issue"]) {
		if text := extractIssueText(issue); text != "" {
			messages = append(messages, text)
		}
	}
	return messages
}

// GetResourceIds returns a slice of maps with fullUrl, id, and location for all entries with an ID.
func (b *BundleResponse) GetResourceIds() []map[string]string {
	ids := []map[string]string{}
	for idx, entry := range b.GetEntries() {
		id := b.GetEntryResourceId(idx)
		if id != "" {
			fullURL, _ := entry["fullUrl"].(string)
			location, _ := entry["location"].(string)
			ids = append(ids, map[string]string{
				"fullUrl":  fullURL,
				"id":       id,
				"location": location,
			})
		}
	}
	return ids
}

// IsAllSuccess returns true if there is at least one entry and all entries have 2xx status.
func (b *BundleResponse) IsAllSuccess() bool {
	entries := b.GetEntries()
	if len(entries) == 0 {
		return false
	}
	for _, entry := range entries {
		status, _ := entry["status"].(string)
		code := parseStatusCode(status)
		if code < 200 || code >= 300 {
			return false
		}
	}
	return true
}

// ByMethod filters response entries by their corresponding request HTTP method(s).
func (b *BundleResponse) ByMethod(methods []string) []map[string]any {
	if len(b.requestEntries) == 0 {
		return []map[string]any{}
	}

	methodSet := make(map[string]bool, len(methods))
	for _, m := range methods {
		methodSet[strings.ToUpper(m)] = true
	}

	results := []map[string]any{}
	responseEntries := toMapSlice(b.bundle["entry"])

	for reqIdx, reqEntry := range b.requestEntries {
		method := extractReqMethod(reqEntry)
		if !methodSet[method] {
			continue
		}

		reqURL := extractReqURL(reqEntry)
		fullURL, _ := reqEntry["fullUrl"].(string)

		var respEntry map[string]any
		if reqIdx < len(responseEntries) {
			respEntry = responseEntries[reqIdx]
		}
		var resp map[string]any
		if respEntry != nil {
			resp, _ = respEntry["response"].(map[string]any)
		}
		if resp == nil {
			resp = map[string]any{}
		}

		status, _ := resp["status"].(string)
		statusCode := parseStatusCode(status)
		location, _ := resp["location"].(string)
		resourceID := parseIdFromUrl(location)
		etag, _ := resp["etag"].(string)
		lastModified, _ := resp["lastModified"].(string)
		outcome := resp["outcome"]

		errors := []string{}
		if outcomeMap, ok := outcome.(map[string]any); ok {
			if rt, _ := outcomeMap["resourceType"].(string); rt == "OperationOutcome" {
				for _, issue := range toMapSlice(outcomeMap["issue"]) {
					if text := extractIssueText(issue); text != "" {
						errors = append(errors, text)
					}
				}
			}
		}

		var resource any
		if method == "GET" && respEntry != nil {
			resource = respEntry["resource"]
		}

		var resIDVal any
		if resourceID != "" {
			resIDVal = resourceID
		}

		results = append(results, map[string]any{
			"index":        reqIdx,
			"method":       method,
			"url":          reqURL,
			"fullUrl":      fullURL,
			"status":       status,
			"statusCode":   statusCode,
			"location":     location,
			"resourceId":   resIDVal,
			"etag":         etag,
			"lastModified": lastModified,
			"outcome":      outcome,
			"errors":       errors,
			"isSuccess":    statusCode >= 200 && statusCode < 300,
			"resource":     resource,
		})
	}

	return results
}

// Created returns entries that were POST (created resources).
func (b *BundleResponse) Created() []map[string]any {
	return b.ByMethod([]string{"POST"})
}

// Updated returns entries that were PUT (updated resources).
func (b *BundleResponse) Updated() []map[string]any {
	return b.ByMethod([]string{"PUT"})
}

// Patched returns entries that were PATCH (partially updated resources).
func (b *BundleResponse) Patched() []map[string]any {
	return b.ByMethod([]string{"PATCH"})
}

// Read returns entries that were GET (read results — includes full resource).
func (b *BundleResponse) Read() []map[string]any {
	return b.ByMethod([]string{"GET"})
}

// Deleted returns entries that were DELETE (deleted resources).
func (b *BundleResponse) Deleted() []map[string]any {
	return b.ByMethod([]string{"DELETE"})
}

// IdMap builds a map of request fullUrl -> server-assigned resource ID for created entries.
func (b *BundleResponse) IdMap() map[string]string {
	m := map[string]string{}
	for _, entry := range b.Created() {
		fullURL, _ := entry["fullUrl"].(string)
		resID := ""
		if entry["resourceId"] != nil {
			resID = fmt.Sprintf("%v", entry["resourceId"])
		}
		if fullURL != "" && resID != "" {
			m[fullURL] = resID
		}
	}
	return m
}

// Find looks up a response entry by resource type and ID from the request URL.
func (b *BundleResponse) Find(resourceType, resourceId string) map[string]any {
	rtLower := strings.ToLower(resourceType)
	idLower := strings.ToLower(resourceId)
	responseEntries := toMapSlice(b.bundle["entry"])

	for reqIdx, reqEntry := range b.requestEntries {
		reqURL := extractReqURL(reqEntry)
		urlLower := strings.ToLower(reqURL)
		if strings.Contains(urlLower, rtLower) && strings.Contains(urlLower, idLower) {
			var resp map[string]any
			if reqIdx < len(responseEntries) {
				resp, _ = responseEntries[reqIdx]["response"].(map[string]any)
			}
			if resp == nil {
				resp = map[string]any{}
			}
			status, _ := resp["status"].(string)
			statusCode := parseStatusCode(status)
			location, _ := resp["location"].(string)
			resID := b.GetEntryResourceId(reqIdx)
			var resIDVal any
			if resID != "" {
				resIDVal = resID
			}
			return map[string]any{
				"index":      reqIdx,
				"method":     extractReqMethod(reqEntry),
				"status":     status,
				"statusCode": statusCode,
				"location":   location,
				"resourceId": resIDVal,
				"isSuccess":  statusCode >= 200,
			}
		}
	}
	return nil
}

func extractReqMethod(reqEntry map[string]any) string {
	if m, ok := reqEntry["method"].(string); ok && m != "" {
		return strings.ToUpper(m)
	}
	if req, ok := reqEntry["request"].(map[string]any); ok {
		if m, ok := req["method"].(string); ok && m != "" {
			return strings.ToUpper(m)
		}
	}
	return ""
}

func extractReqURL(reqEntry map[string]any) string {
	if u, ok := reqEntry["url"].(string); ok && u != "" {
		return u
	}
	if req, ok := reqEntry["request"].(map[string]any); ok {
		if u, ok := req["url"].(string); ok && u != "" {
			return u
		}
	}
	return ""
}

func parseStatusCode(status string) int {
	if status == "" {
		return 0
	}
	parts := strings.Split(strings.TrimSpace(status), " ")
	code, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0
	}
	return code
}

func copyMap(m map[string]any) map[string]any {
	out := make(map[string]any, len(m)+1)
	for k, v := range m {
		out[k] = v
	}
	return out
}
