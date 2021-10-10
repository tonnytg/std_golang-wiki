package main

import (
	"fmt"
)

func main() {

	fmt.Println("start Main method")
	// Creating a channel
	ch := make(chan int)

	go myfunc(ch)
	ch <- 23
	fmt.Println("End Main method")
}
func myfunc(ch chan int) {

	fmt.Println(234 + <-ch)
}
