package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	var sum int

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s := 0
		for v := range ch {
			s += v
		}
		sum = s
	}()

	wg.Wait()
	fmt.Println("Сумма чисел от 1 до 10 равна", sum)
}
