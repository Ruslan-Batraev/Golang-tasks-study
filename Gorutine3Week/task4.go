package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type AtomicBool struct {
	value int32
}

func (b *AtomicBool) Set(val bool) {
	var v int32
	if val {
		v = 1
	}
	atomic.StoreInt32(&b.value, v)
}

func (b *AtomicBool) Get() bool {
	return atomic.LoadInt32(&b.value) != 0
}

func main() {
	flag := &AtomicBool{}
	flag.Set(true)

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				fmt.Printf("Горутина %d: флаг = %v\n", id, flag.Get())
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	time.Sleep(200 * time.Millisecond)
	flag.Set(false)
	fmt.Println("Флаг изменён на false")

	wg.Wait()
}
