package main

import "fmt"

// Definição de uma estrutura de lista ligada (LinkedList)
type LinkedList struct {
	value  any
	next   *LinkedList
	before *LinkedList
}

func main() {
	// Inicializa o primeiro nó
	start := LinkedList{
		value: 1,
	}

	// Inicializa o segundo nó e conecta ao primeiro
	next := LinkedList{
		value:  2,
		before: &start,
	}

	// Conecta o próximo nó ao nó inicial
	start.next = &next

	// Exibe os valores para verificar a conexão
	fmt.Printf("Start value: %v\n", start.value)
	fmt.Printf("Next value: %v\n", start.next.value)
	fmt.Printf("Before value: %v\n", next.before.value)
}
