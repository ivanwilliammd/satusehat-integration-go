package builder

import (
	"time"
)

// BundleBuilder builds FHIR Bundle payload
type BundleBuilder struct {
	ResourceType    string
	Data            map[string]interface{}
	autoTimestamp   bool
}

const (
	BundleTypeDocument          = "document"
	BundleTypeBatch             = "batch"
	BundleTypeTransaction       = "transaction"
	BundleTypeHistoryCollection = "history-collection"
	BundleTypeHistoryDocument   = "history-document"
	BundleTypeHistoryFeed       = "history-feed"
	BundleTypeSearchset         = "searchset"
	BundleTypeCollection       = "collection"
	BundleTypeFeed             = "feed"
	BundleTypeWrapper          = "wrapper"
)

func NewBundleBuilder() *BundleBuilder {
	b := &BundleBuilder{
		ResourceType:  "Bundle",
		Data:          make(map[string]interface{}),
		autoTimestamp: true,
	}
	return b
}

func (b *BundleBuilder) SetId(id string) *BundleBuilder {
	b.Data["id"] = id
	return b
}

func (b *BundleBuilder) SetType(bundleType string) *BundleBuilder {
	b.Data["type"] = bundleType
	return b
}

func (b *BundleBuilder) SetTimestamp(timestamp string) *BundleBuilder {
	b.autoTimestamp = false
	b.Data["timestamp"] = timestamp
	return b
}

func (b *BundleBuilder) SetTotal(total int) *BundleBuilder {
	b.Data["total"] = total
	return b
}

func (b *BundleBuilder) SetMeta(meta map[string]interface{}) *BundleBuilder {
	b.Data["meta"] = meta
	return b
}

func (b *BundleBuilder) AddLink(relation string, url string) *BundleBuilder {
	if _, ok := b.Data["link"]; !ok {
		b.Data["link"] = make([]interface{}, 0)
	}
	b.Data["link"] = append(b.Data["link"].([]interface{}), map[string]interface{}{
		"relation": relation,
		"url":      url,
	})
	return b
}

func (b *BundleBuilder) AddEntry(resource map[string]interface{}, fullUrl string) *BundleBuilder {
	entry := map[string]interface{}{"resource": resource}
	if fullUrl != "" {
		entry["fullUrl"] = fullUrl
	}
	if _, ok := b.Data["entry"]; !ok {
		b.Data["entry"] = make([]interface{}, 0)
	}
	b.Data["entry"] = append(b.Data["entry"].([]interface{}), entry)
	return b
}

func (b *BundleBuilder) AddSearchEntry(resource map[string]interface{}, fullUrl string, score float64, searchMode string) *BundleBuilder {
	entry := map[string]interface{}{"resource": resource}
	if fullUrl != "" {
		entry["fullUrl"] = fullUrl
	}
	search := map[string]interface{}{}
	if searchMode != "" {
		search["mode"] = searchMode
	}
	if score != 0 {
		search["score"] = score
	}
	if len(search) > 0 {
		entry["search"] = search
	}
	if _, ok := b.Data["entry"]; !ok {
		b.Data["entry"] = make([]interface{}, 0)
	}
	b.Data["entry"] = append(b.Data["entry"].([]interface{}), entry)
	return b
}

func (b *BundleBuilder) AddBatchEntry(resource map[string]interface{}, fullUrl string, method string, url string, ifMatch string, ifNoneMatch string, ifNoneExist string) *BundleBuilder {
	request := map[string]interface{}{
		"method": method,
		"url":    url,
	}
	if ifMatch != "" {
		request["ifMatch"] = ifMatch
	}
	if ifNoneMatch != "" {
		request["ifNoneMatch"] = ifNoneMatch
	}
	if ifNoneExist != "" {
		request["ifNoneExist"] = ifNoneExist
	}
	entry := map[string]interface{}{
		"fullUrl": fullUrl,
		"request": request,
	}
	if resource != nil {
		entry["resource"] = resource
	}
	if _, ok := b.Data["entry"]; !ok {
		b.Data["entry"] = make([]interface{}, 0)
	}
	b.Data["entry"] = append(b.Data["entry"].([]interface{}), entry)
	return b
}

func (b *BundleBuilder) AddTransactionEntry(resource map[string]interface{}, fullUrl string, method string, url string, ifMatch string, ifNoneMatch string, ifNoneExist string) *BundleBuilder {
	b.SetType("transaction")
	return b.AddBatchEntry(resource, fullUrl, method, url, ifMatch, ifNoneMatch, ifNoneExist)
}

func (b *BundleBuilder) AddGetEntry(fullUrl string, url string, ifNoneMatch string) *BundleBuilder {
	return b.AddBatchEntry(nil, fullUrl, "GET", url, "", ifNoneMatch, "")
}

func (b *BundleBuilder) AddDeleteEntry(fullUrl string, url string, ifMatch string) *BundleBuilder {
	return b.AddBatchEntry(nil, fullUrl, "DELETE", url, ifMatch, "", "")
}

func (b *BundleBuilder) Build() map[string]interface{} {
	if b.autoTimestamp {
		b.Data["timestamp"] = time.Now().UTC().Format(time.RFC3339)
	}
	return b.Data
}
