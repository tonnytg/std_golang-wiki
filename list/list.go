// Package list create a queue
package list

import (
	"container/list"
	"fmt"
)

// MyList it is a function create a queue like a FIFO queue
// ties func depends on container/list
func MyList(itens ...string){
	queue := list.New()

	for _, item := range itens {
		queue.PushBack(item)
	}

	for queue.Len() > 0 {
		e := queue.Front() // First element
		fmt.Println(e.Value)

		queue.Remove(e) // Dequeue
	}
}
