package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type RateLimiter struct {
	limit     int64
	interval  int64
	counter   int64
	lastReset int64
}

func NewRateLimiter(limit int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:     int64(limit),
		interval:  interval.Nanoseconds(),
		lastReset: time.Now().UnixNano(),
	}
}

func (rl *RateLimiter) Allow() bool {
	now := time.Now().UnixNano()
	interval := atomic.LoadInt64(&rl.interval)

	for {
		last := atomic.LoadInt64(&rl.lastReset)
		if now-last > interval {
			if atomic.CompareAndSwapInt64(&rl.lastReset, last, now) {
				// Сброс счётчика
				atomic.StoreInt64(&rl.counter, 0)
			} else {
				continue
			}
		}
		break
	}

	current := atomic.AddInt64(&rl.counter, 1)
	return current <= rl.limit
}

func main() {
	limiter := NewRateLimiter(5, time.Second)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				if limiter.Allow() {
					fmt.Printf("Горутина %d: запрос %d разрешён\n", id, j)
				} else {
					fmt.Printf("Горутина %d: запрос %d отклонён (лимит)\n", id, j)
				}
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
}
