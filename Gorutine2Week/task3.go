package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	counter int
	mut     sync.RWMutex
}

func (s *SafeCounter) Inc() {
	s.mut.Lock()
	defer s.mut.Unlock()
	s.counter++
}

func (s *SafeCounter) Value() int {
	s.mut.Lock()
	defer s.mut.Unlock()
	return s.counter
}

func main() {
	s := SafeCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.Inc()
		}()
	}
	wg.Wait()
	fmt.Println(s.Value)
}
