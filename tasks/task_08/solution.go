package main

import (
	"sync"
	"time"
)

type (
	Clock   interface{ Now() time.Time }
	Limiter struct {
		mu     sync.Mutex
		clock  Clock
		rate   float64
		burst  int
		tokens float64
		last   time.Time
	}
)

func NewLimiter(clock Clock, ratePerSec float64, burst int) *Limiter {
	l := &Limiter{
		clock:  clock,
		rate:   ratePerSec,
		burst:  burst,
		tokens: float64(burst),
	}
	if clock != nil {
		l.last = clock.Now()
	}
	return l
}

func (l *Limiter) AllowN(n int) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.burst <= 0 || l.clock == nil {
		return false
	}
	if n <= 0 || n > l.burst {
		return false
	}

	if l.rate > 0 {
		now := l.clock.Now()
		if elapsed := now.Sub(l.last); elapsed > 0 {
			l.tokens += elapsed.Seconds() * l.rate
			if l.tokens > float64(l.burst) {
				l.tokens = float64(l.burst)
			}
			l.last = now
		}
	}

	if l.tokens >= float64(n) {
		l.tokens -= float64(n)
		return true
	}
	return false
}
