package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool)

	//stop := false
	//
	//go func() {
	//
	//	for !stop {
	//
	//		fmt.Println("Tick")
	//
	//		time.Sleep(time.Millisecond * 300)
	//
	//	}
	//	fmt.Println("Stopped")
	//	done <- true
	//}()
	//
	//time.Sleep(time.Millisecond * 1500)
	//stop = true
	//time.Sleep(time.Millisecond * 100)

	go func(done chan bool) {

		for {
			select {
			case <-done:
				return
			default:
				fmt.Println("TicTok")
				time.Sleep(time.Millisecond * 300)
			}
		}

	}(done)

	time.Sleep(1500 * time.Millisecond)
	done <- true
	fmt.Println("Stopped")
}
