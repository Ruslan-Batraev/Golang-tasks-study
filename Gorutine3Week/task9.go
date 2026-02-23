package main

import (
	"fmt"
	"sync"
	"time"
)

type TimestampStore struct {
	mu    sync.RWMutex
	store map[string]time.Time
}

func NewTimestampStore() *TimestampStore {
	return &TimestampStore{
		store: make(map[string]time.Time),
	}
}

func (ts *TimestampStore) Get(event string) (time.Time, bool) {
	ts.mu.RLock()
	defer ts.mu.RUnlock()
	t, ok := ts.store[event]
	return t, ok
}

func (ts *TimestampStore) Set(event string, t time.Time) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	ts.store[event] = t
}

func (ts *TimestampStore) GetAll() map[string]time.Time {
	ts.mu.RLock()
	defer ts.mu.RUnlock()
	copyMap := make(map[string]time.Time, len(ts.store))
	for k, v := range ts.store {
		copyMap[k] = v
	}
	return copyMap
}

func main() {
	store := NewTimestampStore()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 20; j++ {
				event := fmt.Sprintf("event_%d", j%5)
				t, ok := store.Get(event)
				if ok {
					fmt.Printf("Читатель %d: событие %s, время %v\n", id, event, t.Format("15:04:05.000"))
				} else {
					fmt.Printf("Читатель %d: событие %s ещё не записано\n", id, event)
				}
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for j := 0; j < 5; j++ {
			event := fmt.Sprintf("event_%d", j%5)
			store.Set(event, time.Now())
			fmt.Printf("Писатель: обновлено событие %s\n", event)
			time.Sleep(50 * time.Millisecond)
		}
	}()

	wg.Wait()
	fmt.Println("Все горутины завершили работу")
}
