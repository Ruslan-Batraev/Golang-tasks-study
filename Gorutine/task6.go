package main

import "fmt"

func main() {
	ch := make(chan struct{})

	go func() {
		fmt.Println("Worker work")
		close(ch)
		fmt.Println("Worker done")
	}()

	<-ch

	fmt.Println("Main continues")
}
