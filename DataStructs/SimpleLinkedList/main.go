package main

import "fmt"

// Node define um nó na lista
type Node struct {
	value int
	next  *Node
}

// LinkedList define a lista encadeada
type LinkedList struct {
	head *Node
}

// Push adiciona um novo valor no início da lista
func (ll *LinkedList) Push(value int) {
	newNode := &Node{value: value, next: ll.head}
	ll.head = newNode
}

// Pop remove e retorna o valor do início da lista
func (ll *LinkedList) Pop() (int, bool) {
	if ll.head == nil {
		return 0, false
	}
	value := ll.head.value
	ll.head = ll.head.next
	return value, true
}

func main() {
	ll := &LinkedList{}

	// Adicionar elementos
	ll.Push(10)
	ll.Push(20)

	// Remover elementos
	val, ok := ll.Pop()
	if ok {
		fmt.Println("Popped:", val)
	}

	val, ok = ll.Pop()
	if ok {
		fmt.Println("Popped:", val)
	}
}

