package main

import "fmt"

func main() {

	mes := make(chan string, 2)

	go func() {
		mes <- "first"
		mes <- "second"

		close(mes)
	}()

	val1 := <-mes
	val2 := <-mes

	fmt.Println("val1:", val1, "val2:", val2)
}
