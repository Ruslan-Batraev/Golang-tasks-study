package main

import (
	"fmt"
	"time"
)

func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func main() {
	go greet("Alice")
	time.Sleep(1 * time.Second)
}
