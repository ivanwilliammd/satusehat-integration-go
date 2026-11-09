package queue

import (
	"fmt"
	"math"
	"sort"
	"time"
)

// QueueStats holds summary counts for queue statuses.
type QueueStats struct {
	Pending    int
	Processing int
	Success    int
	Failed     int
	DLQ        int
}

// QueueMonitor provides queue statistics, health checks, and error analytics.
type QueueMonitor struct {
	queue *SqliteQueue
}

// NewQueueMonitor creates a new QueueMonitor for the given SqliteQueue.
func NewQueueMonitor(queue *SqliteQueue) *QueueMonitor {
	return &QueueMonitor{queue: queue}
}

// Summary returns high-level queue counts, success rate, and oldest pending timestamp.
func (m *QueueMonitor) Summary() map[string]any {
	stats := m.queue.Stats()
	total := 0
	for _, cnt := range stats {
		total += cnt
	}

	successRate := 0.0
	if total > 0 {
		successRate = math.Round((float64(stats[StatusSuccess])/float64(total))*1000) / 10
	}

	oldest := m.oldestPending()
	var oldestPendingAt any
	if oldest != nil {
		if createdAt, ok := oldest["created_at"].(string); ok && createdAt != "" {
			oldestPendingAt = createdAt
		}
	}

	return map[string]any{
		"total":             total,
		"pending":           stats[StatusPending],
		"processing":        stats[StatusProcessing],
		"success":           stats[StatusSuccess],
		"failed":            stats[StatusFailed],
		"dlq":               stats[StatusDlq],
		"success_rate":      successRate,
		"oldest_pending_at": oldestPendingAt,
		"dlq_count":         stats[StatusDlq],
	}
}

// HealthCheck evaluates queue health and returns status ("healthy", "degraded", "critical"), issues, and summary.
func (m *QueueMonitor) HealthCheck() map[string]any {
	stats := m.queue.Stats()
	status := "healthy"
	issues := []string{}

	if stats[StatusDlq] > 0 && stats[StatusDlq] >= stats[StatusSuccess] {
		status = "critical"
		issues = append(issues, "DLQ count >= success count — investigate failed jobs")
	} else if stats[StatusDlq] > 20 {
		status = "degraded"
		issues = append(issues, fmt.Sprintf("%d jobs in DLQ — review and requeue or fix payload", stats[StatusDlq]))
	}

	if stats[StatusPending] > 500 {
		if status == "healthy" {
			status = "degraded"
		}
		issues = append(issues, fmt.Sprintf("%d pending jobs — worker may be idle or under-provisioned", stats[StatusPending]))
	}

	oldest := m.oldestPending()
	if oldest != nil {
		if createdStr, ok := oldest["created_at"].(string); ok && createdStr != "" {
			if createdTime, err := time.Parse(queueTimeFormat, createdStr); err == nil {
				ageMinutes := time.Since(createdTime).Minutes()
				if ageMinutes > 30 && stats[StatusProcessing] == 0 {
					if status == "healthy" {
						status = "degraded"
					}
					issues = append(issues, fmt.Sprintf("Oldest pending job is %.0fmin old — worker may be down", ageMinutes))
				}
			}
		}
	}

	return map[string]any{
		"status":  status,
		"issues":  issues,
		"summary": m.Summary(),
	}
}

// TopErrors returns the top error messages from failed and DLQ jobs ordered by count descending.
func (m *QueueMonitor) TopErrors(limit ...int) []map[string]any {
	maxItems := 10
	if len(limit) > 0 && limit[0] > 0 {
		maxItems = limit[0]
	}

	counts := map[string]int{}
	for _, job := range m.queue.AllJobs() {
		status, _ := job["status"].(string)
		if status != StatusFailed && status != StatusDlq {
			continue
		}
		lastErr, _ := job["last_error"].(string)
		if lastErr == "" {
			continue
		}
		counts[lastErr]++
	}

	type errEntry struct {
		err string
		cnt int
	}
	entries := make([]errEntry, 0, len(counts))
	for k, v := range counts {
		entries = append(entries, errEntry{err: k, cnt: v})
	}
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].cnt != entries[j].cnt {
			return entries[i].cnt > entries[j].cnt
		}
		return entries[i].err < entries[j].err
	})

	if len(entries) > maxItems {
		entries = entries[:maxItems]
	}

	out := make([]map[string]any, 0, len(entries))
	for _, e := range entries {
		out = append(out, map[string]any{
			"last_error": e.err,
			"cnt":        e.cnt,
		})
	}
	return out
}

func (m *QueueMonitor) oldestPending() map[string]any {
	jobs := m.queue.AllJobs()
	var oldest map[string]any
	for _, job := range jobs {
		status, _ := job["status"].(string)
		if status != StatusPending && status != StatusFailed {
			continue
		}
		if oldest == nil {
			oldest = job
			continue
		}
		c1, _ := job["created_at"].(string)
		c2, _ := oldest["created_at"].(string)
		if c1 < c2 {
			oldest = job
		}
	}
	return oldest
}
