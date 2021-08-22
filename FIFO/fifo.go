package FIFO

import "fmt"

func Fifo(words ...string) {
	var queue []string

	for _, word := range words {
		//total += word
		queue = append(queue, word)
	}

	for len(queue) > 0 {
		fmt.Print(queue[0], " ") // First element
		queue = queue[1:]   // Dequeue
	}
}
