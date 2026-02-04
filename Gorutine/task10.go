package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(2)

	ch := make(chan int)

	send(ch, wg)
	reder(ch, wg)
	wg.Wait()

}

func send(ch chan int, wg *sync.WaitGroup) chan int {
	go func() {
		defer wg.Done()
		for i := 1; i < 4; i++ {
			ch <- i
		}
		close(ch)

	}()
	return ch
}

func reder(ch chan int, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for i := 1; i < 4; i++ {
			fmt.Println(<-ch)
		}
	}()

}
