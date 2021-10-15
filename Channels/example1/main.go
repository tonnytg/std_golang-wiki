package main

import (
	"fmt"
)

func main() {
	// Create Chan type int
	c := make(chan int)
	// Pass Chan bi for send type
	go MyLoop(10, c)
	// Pass chan bi for receive type
	MyPrints(c)
}
// MyLoop only send chan type
func MyLoop(t int, s chan<- int) {
	for i := 0; i < t; i++ {
		s <- i
	}
	close(s)
}
// MyPrints only receive chan type
func MyPrints(r <-chan int) {
	for v := range r {
		fmt.Println("Receive of channel:", v)
	}
}