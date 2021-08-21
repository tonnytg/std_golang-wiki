package FIFO

import "fmt"

func Fifo() {
	var queue []string

	queue = append(queue, "Hello ") // Enqueue
	queue = append(queue, "world!")

	for len(queue) > 0 {
		fmt.Print(queue[0]) // First element
		queue = queue[1:]   // Dequeue
	}
}
