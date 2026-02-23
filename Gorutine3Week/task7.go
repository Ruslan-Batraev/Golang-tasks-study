package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Stats struct {
	Timestamp  time.Time
	Requests   int64
	Errors     int64
	AvgLatency time.Duration
}

type StatsCache struct {
	value atomic.Value
}

func NewStatsCache(initial Stats) *StatsCache {
	cache := &StatsCache{}
	cache.value.Store(&initial)
	return cache
}

func (c *StatsCache) Get() *Stats {
	return c.value.Load().(*Stats)
}

func (c *StatsCache) Update() {
	newStats := &Stats{
		Timestamp:  time.Now(),
		Requests:   rand.Int63n(1000),
		Errors:     rand.Int63n(10),
		AvgLatency: time.Duration(rand.Int63n(100)) * time.Millisecond,
	}
	c.value.Store(newStats)
}

func main() {
	initial := Stats{Timestamp: time.Now()}
	cache := NewStatsCache(initial)

	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			ticker := time.NewTicker(10 * time.Millisecond)
			for range ticker.C {
				stats := cache.Get()
				fmt.Printf("Reader %d: %+v\n", id, stats)
			}
		}(i)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(1 * time.Minute)
		for range ticker.C {
			cache.Update()
			fmt.Println("Cache updated")
		}
	}()

	time.Sleep(2 * time.Minute)
}
