package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type IDGenerator struct {
	counter uint64
}

func NewIDGenerator() *IDGenerator {
	return &IDGenerator{}
}

func (g *IDGenerator) Next() uint64 {
	return atomic.AddUint64(&g.counter, 1)
}

func main() {
	gen := NewIDGenerator()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				id := gen.Next()
				fmt.Printf("Горутина получила ID: %d\n", id)
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Всего сгенерировано ID: %d\n", gen.Next()-1)
}
