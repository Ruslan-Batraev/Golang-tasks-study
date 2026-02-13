package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Stats struct {
	mu       sync.Mutex
	requests int
	errors   int
}

func (s *Stats) RecordRequest() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.requests++
}

func (s *Stats) RecordError() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.errors++
}

func main() {
	rand.Seed(time.Now().UnixNano())

	stats := &Stats{}
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			const requestsPerGoroutine = 1000
			for j := 0; j < requestsPerGoroutine; j++ {
				stats.RecordRequest()
				if rand.Float32() < 0.1 {
					stats.RecordError()
				}
			}
		}(i)
	}

	wg.Wait()

	fmt.Printf("Всего запросов: %d\n", stats.requests)
	fmt.Printf("Ошибок: %d\n", stats.errors)
	fmt.Printf("Процент ошибок: %.2f%%\n", float64(stats.errors)/float64(stats.requests)*100)
}
