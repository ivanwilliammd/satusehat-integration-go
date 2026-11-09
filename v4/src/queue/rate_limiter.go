package queue

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

// RateLimiter implements a per-process sliding window rate limiter.
type RateLimiter struct {
	rpm           float64
	interval      time.Duration
	lastRequestAt time.Time
}

// NewRateLimiter creates a new RateLimiter configured for the given requests per minute (default 300 if <= 0).
func NewRateLimiter(rpm float64) *RateLimiter {
	if rpm < 1 {
		rpm = 1
	}
	interval := time.Duration(float64(time.Minute) / rpm)
	return &RateLimiter{
		rpm:      rpm,
		interval: interval,
	}
}

// Wait blocks until a request slot is available.
func (r *RateLimiter) Wait() {
	if !r.lastRequestAt.IsZero() {
		elapsed := time.Since(r.lastRequestAt)
		if elapsed < r.interval {
			time.Sleep(r.interval - elapsed)
		}
	}
	r.lastRequestAt = time.Now()
}

// ParseRetryAfter parses a Retry-After HTTP header value into seconds to wait.
// Supports both integer seconds and HTTP-date format (e.g., "Wed, 31 Oct 2026 12:00:00 GMT").
func ParseRetryAfter(header string) *float64 {
	trimmed := strings.TrimSpace(header)
	if trimmed == "" {
		return nil
	}

	// Integer seconds
	if secs, err := strconv.Atoi(trimmed); err == nil {
		f := float64(secs)
		return &f
	}

	// HTTP-date format
	if t, err := http.ParseTime(trimmed); err == nil {
		diff := time.Until(t).Seconds()
		if diff < 0 {
			diff = 0
		}
		return &diff
	}

	return nil
}
