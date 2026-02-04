package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		for i := 1; i < 4; i++ {
			defer wg.Done()
			fmt.Println(i)
		}
	}()
	wg.Wait()

}
