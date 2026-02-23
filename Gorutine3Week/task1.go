package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu    sync.RWMutex
	store map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{
		store: make(map[string]interface{}),
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.store[key]
	return val, ok
}

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = value
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.store, key)
}

func (c *Cache) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.store)
}

func main() {
	cache := NewCache()

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				key := fmt.Sprintf("key%d", j%5)
				if val, ok := cache.Get(key); ok {
					fmt.Printf("Reader %d: %s = %v\n", id, key, val)
				} else {
					fmt.Printf("Reader %d: %s not found\n", id, key)
				}
				time.Sleep(1 * time.Millisecond)
			}
		}(i)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for j := 0; j < 10; j++ {
			key := fmt.Sprintf("key%d", j%5)
			cache.Set(key, j*100)
			fmt.Printf("Writer: set %s = %d\n", key, j*100)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	wg.Wait()
	fmt.Println("Итоговый размер кэша:", cache.Len())
}
