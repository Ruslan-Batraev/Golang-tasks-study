package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		arr := []int{1, 2, 3, 4, 5}

		for i := 0; i < len(arr); i++ {
			fmt.Println(arr[i])
			time.Sleep(200 * time.Millisecond)
		}
	}()
	wg.Wait()
}
