package list

import (
	"container/list"
	"fmt"
)

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
