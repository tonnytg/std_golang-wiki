package main

import (
	"fmt"
	"sync"
	"time"
)

// Create a VAR to storage how many process need wait
// I made this globally to func call at everywhere
var wg sync.WaitGroup

func HelloFriend(seconds int) {
	fmt.Printf("Hi I need to wait %d seconds\n", seconds)
	d := time.Duration(seconds)
	time.Sleep(d * time.Second)
	fmt.Printf( "Bye, I waited %d seconds\n", seconds)

	// Now after all commands you need clear a Wait Group
	wg.Done()
}

func main() {
	for i := 0; i <=5; i++ {
		wg.Add(1)
		go HelloFriend(i)

		// See this lines print without HelloFriend return
		// This happens because Main() don't wait Goroutines end by default
		fmt.Println("---")
	}

	// Say to Main don't stop running because exist some goroutine running
	wg.Wait()
}
