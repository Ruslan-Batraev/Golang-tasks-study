package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	data map[string]int
	mu   sync.RWMutex
}

func (sm *SafeMap) Set(key string, val int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = val
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	v, ok := sm.data[key]
	return v, ok
}

func main() {
	var wg sync.WaitGroup
	sm := SafeMap{data: make(map[string]int), mu: sync.RWMutex{}}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sm.Set(fmt.Sprintf("key_%d", i), i)
		}()
	}
	wg.Wait()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(sm.Get(fmt.Sprintf("key_%d", i)))
		}()
	}
	wg.Wait()
}
