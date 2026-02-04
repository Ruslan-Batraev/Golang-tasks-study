package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main started")
	go func() {
		fmt.Println("Gorutine finished")

	}()
	fmt.Println("Main finished")
	time.Sleep(1 * time.Second)
}
