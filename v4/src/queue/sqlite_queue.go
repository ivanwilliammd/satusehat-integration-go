package queue

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

const queueTimeFormat = "2006-01-02 15:04:05"

// SqliteQueue implements QueueInterface with zero-CGo thread-safe persistence
// (supports ":memory:" or file-backed durable storage).
type SqliteQueue struct {
	mu       sync.Mutex
	filePath string
	nextID   int
	records  []map[string]any
}

type persistedQueueState struct {
	NextID  int              `json:"next_id"`
	Records []map[string]any `json:"records"`
}

// NewSqliteQueue creates a new SqliteQueue instance.
// Pass ":memory:" or "" for in-memory storage, or a file path for durable persistence.
func NewSqliteQueue(dsn ...string) *SqliteQueue {
	path := ":memory:"
	if len(dsn) > 0 && dsn[0] != "" {
		path = strings.TrimPrefix(dsn[0], "sqlite:")
	}

	q := &SqliteQueue{
		filePath: path,
		nextID:   1,
		records:  []map[string]any{},
	}

	if path != ":memory:" {
		q.loadFromDisk()
	}

	return q
}

func (q *SqliteQueue) loadFromDisk() {
	raw, err := os.ReadFile(q.filePath)
	if err != nil || len(raw) == 0 {
		return
	}
	var state persistedQueueState
	if err := json.Unmarshal(raw, &state); err == nil {
		q.nextID = state.NextID
		if q.nextID < 1 {
			q.nextID = 1
		}
		q.records = state.Records
		for i := range q.records {
			if idFloat, ok := q.records[i]["id"].(float64); ok {
				q.records[i]["id"] = int(idFloat)
			}
			if attFloat, ok := q.records[i]["attempts"].(float64); ok {
				q.records[i]["attempts"] = int(attFloat)
			}
			if maxFloat, ok := q.records[i]["max_attempts"].(float64); ok {
				q.records[i]["max_attempts"] = int(maxFloat)
			}
		}
	}
}

func (q *SqliteQueue) saveToDisk() {
	if q.filePath == "" || q.filePath == ":memory:" {
		return
	}
	state := persistedQueueState{
		NextID:  q.nextID,
		Records: q.records,
	}
	if raw, err := json.Marshal(state); err == nil {
		_ = os.WriteFile(q.filePath, raw, 0644)
	}
}

// EnqueueBundle enqueues a FHIR Bundle transaction or batch payload.
func (q *SqliteQueue) EnqueueBundle(
	bundlePayload map[string]any,
	bundleType string,
	userId string,
	maxAttempts int,
	metadata map[string]any,
) map[string]any {
	q.mu.Lock()
	defer q.mu.Unlock()

	if bundleType == "" {
		bundleType = "transaction"
	}
	if userId == "" {
		userId = "system"
	}
	if maxAttempts <= 0 {
		maxAttempts = DefaultMaxAttempts
	}

	nowStr := time.Now().UTC().Format(queueTimeFormat)
	id := q.nextID
	q.nextID++

	rec := map[string]any{
		"id":              id,
		"uuid":            generateUUID(),
		"bundle_type":     bundleType,
		"bundle_payload":  cloneMap(bundlePayload),
		"resource_type":   "Bundle",
		"method":          "POST",
		"url":             "Bundle",
		"status":          StatusPending,
		"attempts":        0,
		"max_attempts":    maxAttempts,
		"last_error":      nil,
		"last_response":   nil,
		"etag":            nil,
		"idempotency_key": nil,
		"scheduled_at":    nil,
		"completed_at":    nil,
		"dlq_reason":      nil,
		"user_id":         userId,
		"metadata":        cloneMap(metadata),
		"created_at":      nowStr,
		"updated_at":      nowStr,
	}

	q.records = append(q.records, rec)
	q.saveToDisk()
	return cloneRecord(rec)
}

// Enqueue enqueues a single FHIR resource operation.
func (q *SqliteQueue) Enqueue(
	method string,
	resourceType string,
	url string,
	payload map[string]any,
	etag string,
	idempotencyKey string,
	userId string,
	maxAttempts int,
	metadata map[string]any,
) map[string]any {
	q.mu.Lock()
	defer q.mu.Unlock()

	if method == "" {
		method = "POST"
	}
	if userId == "" {
		userId = "system"
	}
	if maxAttempts <= 0 {
		maxAttempts = DefaultMaxAttempts
	}

	jobUUID := idempotencyKey
	if jobUUID == "" {
		jobUUID = generateUUID()
	}

	var etagVal any
	if etag != "" {
		etagVal = etag
	}
	var idemVal any
	if idempotencyKey != "" {
		idemVal = idempotencyKey
	}

	nowStr := time.Now().UTC().Format(queueTimeFormat)
	id := q.nextID
	q.nextID++

	rec := map[string]any{
		"id":              id,
		"uuid":            jobUUID,
		"bundle_type":     "single",
		"bundle_payload":  cloneMap(payload),
		"resource_type":   resourceType,
		"method":          strings.ToUpper(method),
		"url":             url,
		"status":          StatusPending,
		"attempts":        0,
		"max_attempts":    maxAttempts,
		"last_error":      nil,
		"last_response":   nil,
		"etag":            etagVal,
		"idempotency_key": idemVal,
		"scheduled_at":    nil,
		"completed_at":    nil,
		"dlq_reason":      nil,
		"user_id":         userId,
		"metadata":        cloneMap(metadata),
		"created_at":      nowStr,
		"updated_at":      nowStr,
	}

	q.records = append(q.records, rec)
	q.saveToDisk()
	return cloneRecord(rec)
}

// Dequeue atomically claims the oldest ready-to-process pending job and marks it processing.
func (q *SqliteQueue) Dequeue() map[string]any {
	q.mu.Lock()
	defer q.mu.Unlock()

	now := time.Now().UTC()
	nowStr := now.Format(queueTimeFormat)

	for i := range q.records {
		rec := q.records[i]
		status, _ := rec["status"].(string)
		if status != StatusPending {
			continue
		}
		if !isReadyToRun(rec["scheduled_at"], now) {
			continue
		}

		// Snapshot row before increment (matches PHP SqliteQueue which fetches row then increments attempts in DB)
		snapshot := cloneRecord(rec)

		attempts := intFromAny(rec["attempts"])
		rec["status"] = StatusProcessing
		rec["attempts"] = attempts + 1
		rec["updated_at"] = nowStr
		q.saveToDisk()

		return snapshot
	}

	return nil
}

// MarkSuccess marks a job as succeeded.
func (q *SqliteQueue) MarkSuccess(id int, response map[string]any, location string) bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	rec := q.findByIDLocked(id)
	if rec == nil {
		return false
	}

	respCopy := cloneMap(response)
	if respCopy == nil {
		respCopy = map[string]any{}
	}
	respCopy["location"] = location

	nowStr := time.Now().UTC().Format(queueTimeFormat)
	rec["status"] = StatusSuccess
	rec["last_response"] = respCopy
	rec["completed_at"] = nowStr
	rec["last_error"] = nil
	rec["updated_at"] = nowStr
	q.saveToDisk()
	return true
}

// MarkFailed marks a job as failed and schedules retry with exponential backoff, or moves to DLQ if attempts >= max_attempts.
func (q *SqliteQueue) MarkFailed(id int, errMsg string, response map[string]any) bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	rec := q.findByIDLocked(id)
	if rec == nil {
		return false
	}

	attempts := intFromAny(rec["attempts"])
	maxAttempts := intFromAny(rec["max_attempts"])
	if attempts >= maxAttempts {
		return q.markDlqLocked(rec, errMsg, response)
	}

	backoffSeconds := int(math.Pow(float64(attempts), 2)) * 30
	if backoffSeconds > 300 {
		backoffSeconds = 300
	}

	now := time.Now().UTC()
	scheduledAt := now.Add(time.Duration(backoffSeconds) * time.Second).Format(queueTimeFormat)
	nowStr := now.Format(queueTimeFormat)

	rec["status"] = StatusPending
	rec["last_error"] = errMsg
	rec["last_response"] = cloneMap(response)
	rec["scheduled_at"] = scheduledAt
	rec["updated_at"] = nowStr
	q.saveToDisk()
	return true
}

// MarkDlq moves a job to the Dead-Letter Queue.
func (q *SqliteQueue) MarkDlq(id int, reason string, response map[string]any) bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	rec := q.findByIDLocked(id)
	if rec == nil {
		return false
	}
	return q.markDlqLocked(rec, reason, response)
}

func (q *SqliteQueue) markDlqLocked(rec map[string]any, reason string, response map[string]any) bool {
	nowStr := time.Now().UTC().Format(queueTimeFormat)
	rec["status"] = StatusDlq
	rec["dlq_reason"] = reason
	rec["last_error"] = reason
	rec["last_response"] = cloneMap(response)
	rec["updated_at"] = nowStr
	q.saveToDisk()
	return true
}

// Reset resets a single job back to pending with 0 attempts.
func (q *SqliteQueue) Reset(id int) bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	rec := q.findByIDLocked(id)
	if rec == nil {
		return false
	}

	nowStr := time.Now().UTC().Format(queueTimeFormat)
	rec["status"] = StatusPending
	rec["attempts"] = 0
	rec["scheduled_at"] = nil
	rec["last_error"] = nil
	rec["dlq_reason"] = nil
	rec["updated_at"] = nowStr
	q.saveToDisk()
	return true
}

// SetScheduledAt overrides scheduled_at for a pending job (used by Worker for Retry-After).
func (q *SqliteQueue) SetScheduledAt(id int, scheduledAt time.Time) bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	rec := q.findByIDLocked(id)
	if rec == nil {
		return false
	}
	if status, _ := rec["status"].(string); status != StatusPending {
		return false
	}
	rec["scheduled_at"] = scheduledAt.UTC().Format(queueTimeFormat)
	rec["updated_at"] = time.Now().UTC().Format(queueTimeFormat)
	q.saveToDisk()
	return true
}

// Get retrieves a job record by ID.
func (q *SqliteQueue) Get(id int) map[string]any {
	q.mu.Lock()
	defer q.mu.Unlock()

	rec := q.findByIDLocked(id)
	if rec == nil {
		return nil
	}
	return cloneRecord(rec)
}

// PendingCount returns the count of pending jobs ready to process.
func (q *SqliteQueue) PendingCount() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	now := time.Now().UTC()
	cnt := 0
	for _, rec := range q.records {
		status, _ := rec["status"].(string)
		if status == StatusPending && isReadyToRun(rec["scheduled_at"], now) {
			cnt++
		}
	}
	return cnt
}

// DlqCount returns the count of DLQ jobs.
func (q *SqliteQueue) DlqCount() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	cnt := 0
	for _, rec := range q.records {
		if status, _ := rec["status"].(string); status == StatusDlq {
			cnt++
		}
	}
	return cnt
}

// Stats returns counts grouped by status.
func (q *SqliteQueue) Stats() map[string]int {
	q.mu.Lock()
	defer q.mu.Unlock()

	out := map[string]int{
		StatusPending:    0,
		StatusProcessing: 0,
		StatusSuccess:    0,
		StatusFailed:     0,
		StatusDlq:        0,
	}
	for _, rec := range q.records {
		if status, _ := rec["status"].(string); status != "" {
			out[status]++
		}
	}
	return out
}

// PendingJobs returns up to limit ready-to-process pending jobs ordered by ID ascending.
func (q *SqliteQueue) PendingJobs(limit int) []map[string]any {
	q.mu.Lock()
	defer q.mu.Unlock()

	if limit <= 0 {
		limit = 20
	}
	now := time.Now().UTC()
	out := []map[string]any{}
	for _, rec := range q.records {
		status, _ := rec["status"].(string)
		if status == StatusPending && isReadyToRun(rec["scheduled_at"], now) {
			out = append(out, cloneRecord(rec))
			if len(out) >= limit {
				break
			}
		}
	}
	return out
}

// DlqJobs returns up to limit DLQ jobs ordered by ID descending.
func (q *SqliteQueue) DlqJobs(limit int) []map[string]any {
	q.mu.Lock()
	defer q.mu.Unlock()

	if limit <= 0 {
		limit = 20
	}
	out := []map[string]any{}
	for i := len(q.records) - 1; i >= 0; i-- {
		rec := q.records[i]
		if status, _ := rec["status"].(string); status == StatusDlq {
			out = append(out, cloneRecord(rec))
			if len(out) >= limit {
				break
			}
		}
	}
	return out
}

// ResetAll resets all failed and dlq jobs back to pending.
func (q *SqliteQueue) ResetAll() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	nowStr := time.Now().UTC().Format(queueTimeFormat)
	cnt := 0
	for _, rec := range q.records {
		status, _ := rec["status"].(string)
		if status == StatusFailed || status == StatusDlq {
			rec["status"] = StatusPending
			rec["attempts"] = 0
			rec["scheduled_at"] = nil
			rec["last_error"] = nil
			rec["dlq_reason"] = nil
			rec["updated_at"] = nowStr
			cnt++
		}
	}
	if cnt > 0 {
		q.saveToDisk()
	}
	return cnt
}

// PurgeSuccess deletes all success jobs completed older than days.
func (q *SqliteQueue) PurgeSuccess(days int) int {
	q.mu.Lock()
	defer q.mu.Unlock()

	cutoff := time.Now().UTC().Add(-time.Duration(days) * 24 * time.Hour)
	kept := make([]map[string]any, 0, len(q.records))
	purged := 0

	for _, rec := range q.records {
		status, _ := rec["status"].(string)
		if status == StatusSuccess {
			if compStr, ok := rec["completed_at"].(string); ok && compStr != "" {
				if compTime, err := time.Parse(queueTimeFormat, compStr); err == nil {
					if !compTime.After(cutoff) {
						purged++
						continue
					}
				}
			}
		}
		kept = append(kept, rec)
	}

	q.records = kept
	if purged > 0 {
		q.saveToDisk()
	}
	return purged
}

// AllJobs returns a cloned slice of all jobs (used by QueueMonitor).
func (q *SqliteQueue) AllJobs() []map[string]any {
	q.mu.Lock()
	defer q.mu.Unlock()

	out := make([]map[string]any, 0, len(q.records))
	for _, rec := range q.records {
		out = append(out, cloneRecord(rec))
	}
	sort.Slice(out, func(i, j int) bool {
		return intFromAny(out[i]["id"]) < intFromAny(out[j]["id"])
	})
	return out
}

func (q *SqliteQueue) findByIDLocked(id int) map[string]any {
	for _, rec := range q.records {
		if intFromAny(rec["id"]) == id {
			return rec
		}
	}
	return nil
}

func isReadyToRun(scheduledAt any, now time.Time) bool {
	if scheduledAt == nil {
		return true
	}
	s, ok := scheduledAt.(string)
	if !ok || s == "" {
		return true
	}
	t, err := time.Parse(queueTimeFormat, s)
	if err != nil {
		return true
	}
	return !t.After(now)
}

func intFromAny(v any) int {
	switch n := v.(type) {
	case int:
		return n
	case float64:
		return int(n)
	default:
		return 0
	}
}

func cloneMap(m map[string]any) map[string]any {
	if len(m) == 0 {
		return nil
	}
	raw, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	var out map[string]any
	_ = json.Unmarshal(raw, &out)
	return out
}

func generateUUID() string {
	var b [16]byte
	_, _ = rand.Read(b[:])
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

func cloneRecord(rec map[string]any) map[string]any {
	if rec == nil {
		return nil
	}
	out := make(map[string]any, len(rec))
	for k, v := range rec {
		if m, ok := v.(map[string]any); ok {
			out[k] = cloneMap(m)
		} else {
			out[k] = v
		}
	}
	return out
}
