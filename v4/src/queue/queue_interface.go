package queue

// Queue status constants matching PHP QueueInterface.
const (
	StatusPending    = "pending"
	StatusProcessing = "processing"
	StatusSuccess    = "success"
	StatusFailed     = "failed"
	StatusDlq        = "dlq"

	STATUS_PENDING    = StatusPending
	STATUS_PROCESSING = StatusProcessing
	STATUS_SUCCESS    = StatusSuccess
	STATUS_FAILED     = StatusFailed
	STATUS_DLQ        = StatusDlq

	DefaultMaxAttempts   = 5
	DEFAULT_MAX_ATTEMPTS = DefaultMaxAttempts
)

// QueueInterface defines the framework-agnostic contract for the SATUSEHAT queue.
// Status flow: pending -> processing -> success | failed (retry) | dlq
type QueueInterface interface {
	EnqueueBundle(
		bundlePayload map[string]any,
		bundleType string,
		userId string,
		maxAttempts int,
		metadata map[string]any,
	) map[string]any

	Enqueue(
		method string,
		resourceType string,
		url string,
		payload map[string]any,
		etag string,
		idempotencyKey string,
		userId string,
		maxAttempts int,
		metadata map[string]any,
	) map[string]any

	Dequeue() map[string]any

	MarkSuccess(id int, response map[string]any, location string) bool
	MarkFailed(id int, errMsg string, response map[string]any) bool
	MarkDlq(id int, reason string, response map[string]any) bool
	Reset(id int) bool

	Get(id int) map[string]any
	PendingCount() int
	DlqCount() int
	Stats() map[string]int
	PendingJobs(limit int) []map[string]any
	DlqJobs(limit int) []map[string]any

	ResetAll() int
	PurgeSuccess(days int) int
}
