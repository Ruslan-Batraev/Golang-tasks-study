package main

import (
	"fmt"
	"sync"
	"time"
)

var counterr int

var wg sync.WaitGroup

var mu sync.Mutex

func main() {

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counterr++
			mu.Unlock()
		}()
	}
	wg.Wait()
	time.Sleep(1 * time.Second)
	fmt.Println(counterr)

}
