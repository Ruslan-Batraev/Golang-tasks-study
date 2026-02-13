package main

import (
	"fmt"
	"sync"
)

type OnceFlag struct {
	mu   sync.Mutex
	done bool
}

func (o *OnceFlag) Do(f func()) {
	o.mu.Lock()
	if o.done {
		o.mu.Unlock()
		return
	}
	o.done = true
	o.mu.Unlock()
	f()
}

func main() {
	var once OnceFlag
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			once.Do(func() {
				fmt.Printf("Функция вызвана горутиной %d\n", id)
			})
		}(i)
	}

	wg.Wait()
	fmt.Println("Все горутины завершены")
}
