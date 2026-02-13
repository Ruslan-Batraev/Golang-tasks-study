package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	totalgor := 10
	curr := 3

	sem := make(chan struct{}, curr)
	var wg sync.WaitGroup

	for i := 0; i < totalgor; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sem <- struct{}{}
			fmt.Println(id, len(sem))
			time.Sleep(100 * time.Millisecond)

			<-sem
			fmt.Println("Горутина зваершила работу", id)
		}(i)
	}
	wg.Wait()
	close(sem)
	fmt.Println("Конец")
}
