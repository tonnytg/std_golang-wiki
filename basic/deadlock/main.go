package main

import (
	"fmt"
	"time"
)

// Start contain func and wait Time Sleep
func Start(done chan bool) {
	fmt.Println("Start Goroutine")
	time.Sleep(10 * time.Second)
	fmt.Println("Goroutine finished.")
	done <- true
}

func main() {

	// done is a Channel to force main wait for Start() ends.
	done := make(chan bool)

	fmt.Println("Start Main")
	// pass done channel to another scope
	go Start(done)
	time.Sleep(3 * time.Second)
	fmt.Println("Finish Main.")
	// main() waiting channel stay empty to continue, but will be break with time out
	// done <- false // fatal error: all goroutines are asleep - deadlock!

	// clean channel
	<- done

	go Start(done)

	fmt.Println("All Finish")
}
