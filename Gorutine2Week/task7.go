package main

import (
	"fmt"
	"os"
	"sync"
)

type SafeLogger struct {
	mu sync.Mutex
}

func (l *SafeLogger) Log(msg string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Fprintln(os.Stdout, msg)
}

func main() {
	var wg sync.WaitGroup
	logger := &SafeLogger{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				logger.Log(fmt.Sprintf("Горутина %d, сообщение %d", id, j))
			}
		}(i)
	}

	wg.Wait()

}
