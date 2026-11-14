package builder

import (
	"encoding/json"
	"fmt"
	"time"
)

// BundleType lists the bundle types accepted by SetType, mirroring VALID_TYPES in PHP.
var BundleTypes = []string{
	"document", "batch", "transaction", "history-collection", "history-document",
	"history-feed", "searchset", "collection", "feed", "wrapper",
}

// BundleBuilder builds a FHIR R4 Bundle, including transaction/batch entries.
//
// Method-for-method counterpart of PayloadBuilderBundle in the PHP SDK. The constructor stamps
// timestamp with the current time, exactly like PHP; call SetTimestamp or DisableAutoTimestamp for
// a reproducible payload. An invalid type is recorded in Err() instead of throwing, on the same
// convention as ImmunizationBuilder.
//
// @link https://www.hl7.org/fhir/bundle.html
type BundleBuilder struct {
	ResourceType string
	Data         map[string]interface{}
	err          error
	autoStamp    bool
}

// NewBundleBuilder creates an empty Bundle. The optional type is validated and written when given,
// matching the nullable constructor argument in PHP.
func NewBundleBuilder(tipe ...string) *BundleBuilder {
	b := &BundleBuilder{ResourceType: "Bundle", Data: map[string]interface{}{}, autoStamp: true}
	b.Data["resourceType"] = b.ResourceType
	if len(tipe) > 0 && tipe[0] != "" {
		b.SetType(tipe[0])
	}
	if b.autoStamp {
		b.Data["timestamp"] = time.Now().Format(time.RFC3339)
	}
	return b
}

// Err returns the first validation error recorded by a setter, if any.
func (b *BundleBuilder) Err() error { return b.err }

func (b *BundleBuilder) set(k string, v interface{}) *BundleBuilder {
	setPath(b.Data, k, v)
	return b
}

func (b *BundleBuilder) push(k string, v interface{}) *BundleBuilder {
	pushPath(b.Data, k, v)
	return b
}

func (b *BundleBuilder) fail(format string, args ...interface{}) *BundleBuilder {
	if b.err == nil {
		b.err = fmt.Errorf(format, args...)
	}
	return b
}

// DisableAutoTimestamp stops the constructor from stamping the payload.
func (b *BundleBuilder) DisableAutoTimestamp() *BundleBuilder {
	b.autoStamp = false
	return b
}

// SetId sets the logical id.
func (b *BundleBuilder) SetId(id string) *BundleBuilder { return b.set("id", id) }

// SetType validates and sets type.
func (b *BundleBuilder) SetType(typ string) *BundleBuilder {
	for _, sah := range BundleTypes {
		if typ == sah {
			return b.set("type", typ)
		}
	}
	return b.fail("Invalid Bundle type %q. Valid: %v", typ, BundleTypes)
}

// SetTimestamp sets timestamp and disables auto-stamping.
func (b *BundleBuilder) SetTimestamp(timestamp string) *BundleBuilder {
	b.autoStamp = false
	return b.set("timestamp", timestamp)
}

// SetTotal sets total.
func (b *BundleBuilder) SetTotal(total int) *BundleBuilder { return b.set("total", total) }

// SetMeta sets meta.
func (b *BundleBuilder) SetMeta(meta map[string]interface{}) *BundleBuilder {
	return b.set("meta", meta)
}

// AddLink appends a link.
func (b *BundleBuilder) AddLink(relation, url string) *BundleBuilder {
	return b.push("link", map[string]interface{}{"relation": relation, "url": url})
}

// AddPaginationLinks appends self, first, last and, when applicable, previous and next.
func (b *BundleBuilder) AddPaginationLinks(baseURL string, page, count, total int) *BundleBuilder {
	offset := (page - 1) * count
	params := fmt.Sprintf("offset=%d&count=%d", offset, count)
	b.AddLink("self", baseURL+"?"+params)
	b.AddLink("first", fmt.Sprintf("%s?offset=0&count=%d", baseURL, count))
	b.AddLink("last", fmt.Sprintf("%s?offset=%d&count=%d", baseURL, (total/count)*count, count))
	if page > 1 {
		b.AddLink("previous", fmt.Sprintf("%s?offset=%d&count=%d", baseURL, (page-2)*count, count))
	}
	if offset+count < total {
		b.AddLink("next", fmt.Sprintf("%s?offset=%d&count=%d", baseURL, page*count, count))
	}
	return b
}

// AddEntry appends an entry with an optional fullUrl.
func (b *BundleBuilder) AddEntry(resource map[string]interface{}, fullURL string) *BundleBuilder {
	entry := map[string]interface{}{"resource": resource}
	if fullURL != "" {
		entry["fullUrl"] = fullURL
	}
	return b.push("entry", entry)
}

// AddSearchEntry appends a search-result entry. searchMode and score are optional.
func (b *BundleBuilder) AddSearchEntry(resource map[string]interface{}, fullURL, searchMode string, score *float64) *BundleBuilder {
	entry := map[string]interface{}{"resource": resource}
	if fullURL != "" {
		entry["fullUrl"] = fullURL
	}
	search := map[string]interface{}{}
	if searchMode != "" {
		search["mode"] = searchMode
	}
	if score != nil {
		search["score"] = *score
	}
	if len(search) > 0 {
		entry["search"] = search
	}
	return b.push("entry", entry)
}

// AddBatchEntry appends a batch/transaction entry with its request block.
func (b *BundleBuilder) AddBatchEntry(resource map[string]interface{}, fullURL, method, url, ifMatch, ifNoneMatch, ifNoneExist string) *BundleBuilder {
	request := map[string]interface{}{"method": method, "url": url}
	if ifMatch != "" {
		request["ifMatch"] = ifMatch
	}
	if ifNoneMatch != "" {
		request["ifNoneMatch"] = ifNoneMatch
	}
	if ifNoneExist != "" {
		request["ifNoneExist"] = ifNoneExist
	}
	entry := map[string]interface{}{"fullUrl": fullURL, "request": request}
	if resource != nil {
		entry["resource"] = resource
	}
	return b.push("entry", entry)
}

// AddTransactionEntry sets type=transaction, then appends the entry.
func (b *BundleBuilder) AddTransactionEntry(resource map[string]interface{}, fullURL, method, url, ifMatch, ifNoneMatch, ifNoneExist string) *BundleBuilder {
	b.SetType("transaction")
	return b.AddBatchEntry(resource, fullURL, method, url, ifMatch, ifNoneMatch, ifNoneExist)
}

// AddGetEntry appends a GET entry.
func (b *BundleBuilder) AddGetEntry(fullURL, url, ifNoneMatch string) *BundleBuilder {
	return b.AddBatchEntry(nil, fullURL, "GET", url, "", ifNoneMatch, "")
}

// AddDeleteEntry appends a DELETE entry.
func (b *BundleBuilder) AddDeleteEntry(fullURL, url, ifMatch string) *BundleBuilder {
	return b.AddBatchEntry(nil, fullURL, "DELETE", url, ifMatch, "", "")
}

// AddHistoryEntry appends a history entry with its response block.
func (b *BundleBuilder) AddHistoryEntry(resource map[string]interface{}, fullURL, status string, statusCode *int, etag, lastModified, location, outcome string) *BundleBuilder {
	response := map[string]interface{}{"status": status}
	if statusCode != nil {
		response["statusCode"] = *statusCode
	}
	if etag != "" {
		response["etag"] = etag
	}
	if lastModified != "" {
		response["lastModified"] = lastModified
	}
	if location != "" {
		response["location"] = location
	}
	if outcome != "" {
		response["outcome"] = outcome
	}
	return b.push("entry", map[string]interface{}{
		"fullUrl":  fullURL,
		"resource": resource,
		"response": response,
	})
}

// AddEntryWithResponse appends an entry with a minimal response block.
func (b *BundleBuilder) AddEntryWithResponse(resource map[string]interface{}, fullURL, status, etag, lastModified string) *BundleBuilder {
	response := map[string]interface{}{"status": status}
	if etag != "" {
		response["etag"] = etag
	}
	if lastModified != "" {
		response["lastModified"] = lastModified
	}
	return b.push("entry", map[string]interface{}{
		"fullUrl":  fullURL,
		"resource": resource,
		"response": response,
	})
}

// AddOutcomeEntry appends an OperationOutcome entry with search.mode=outcome.
func (b *BundleBuilder) AddOutcomeEntry(operationOutcome map[string]interface{}, fullURL string) *BundleBuilder {
	return b.push("entry", map[string]interface{}{
		"fullUrl":  fullURL,
		"resource": operationOutcome,
		"search":   map[string]interface{}{"mode": "outcome"},
	})
}

// SetSignature sets signature. sigAlg, created, who and onBehalfOf are optional.
func (b *BundleBuilder) SetSignature(typ, sigFormat, sigBlob, sigAlg, created, who, onBehalfOf string) *BundleBuilder {
	signature := map[string]interface{}{
		"type":   []interface{}{map[string]interface{}{"code": typ}},
		"format": sigFormat,
		"data":   sigBlob,
	}
	if sigAlg != "" {
		signature["sigAlg"] = sigAlg
	}
	if created != "" {
		signature["created"] = created
	}
	if who != "" {
		signature["who"] = map[string]interface{}{"reference": who}
	}
	if onBehalfOf != "" {
		signature["onBehalfOf"] = map[string]interface{}{"reference": onBehalfOf}
	}
	return b.set("signature", signature)
}

// Build returns the payload.
func (b *BundleBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *BundleBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
