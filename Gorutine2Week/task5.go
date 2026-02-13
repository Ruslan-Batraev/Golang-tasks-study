package main

import (
	"fmt"
	"sync"
)

//func Merge(ch1, ch2 <-chan int) <-chan int {
//	out := make(chan int)
//	go func() {
//		defer close(out)
//
//		for ch1 != nil || ch2 != nil {
//			select {
//			case v, ok := <-ch1:
//				if ok {
//					out <- v
//				} else {
//					ch1 = nil
//				}
//			case v, ok := <-ch2:
//				if ok {
//					out <- v
//				} else {
//					ch2 = nil
//				}
//			}
//		}
//	}()
//	return out
//}

func Merge(ch1, ch2 <-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	f1 := func(ch <-chan int) {

		defer wg.Done()
		for v := range ch {
			out <- v
		}

	}
	wg.Add(2)
	go f1(ch1)
	go f1(ch2)

	go func() {
		wg.Wait()
		defer close(out)
	}()

	return out
}

func main() {
	mu := &sync.Mutex{}
	ch1 := make(chan int)

	go func() {
		mu.Lock()
		defer close(ch1)
		for i := 1; i <= 5; i++ {
			ch1 <- i
		}
		mu.Unlock()
	}()

	ch2 := make(chan int)

	go func() {
		mu.Lock()
		defer close(ch2)
		for i := 1; i <= 5; i++ {
			ch2 <- i
		}
		mu.Unlock()
	}()

	for v := range Merge(ch1, ch2) {
		fmt.Println(v)
	}
}
