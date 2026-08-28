package queue

import (
	"math"
	"time"
)

type RateLimiter struct {
	rpm            float64
	interval       time.Duration
	lastRequestAt  time.Time
}

func NewRateLimiter(rpm float64) *RateLimiter {
	if rpm < 1 {
		rpm = 1
	}
	// Interval = 60s / rpm
	interval := time.Duration(float64(time.Minute) / rpm)
	return &RateLimiter{
		rpm:      rpm,
		interval: interval,
	}
}

func (r *RateLimiter) Wait() {
	if !r.lastRequestAt.IsZero() {
		elapsed := time.Since(r.lastRequestAt)
		if elapsed < r.interval {
			time.Sleep(r.interval - elapsed)
		}
	}
	r.lastRequestAt = time.Now()
}
