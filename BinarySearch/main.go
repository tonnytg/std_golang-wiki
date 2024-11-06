package main

import (
	"fmt"
	"log"
	"sort"
)

func BinarySearch(x int, listElements []int) (int, error) {
	if len(listElements) == 0 {
		return 0, fmt.Errorf("valor não encontrado")
	}

	mid := len(listElements) / 2

	if listElements[mid] == x {
		return x, nil
	}

	if x < listElements[mid] {
		return BinarySearch(x, listElements[:mid]) // Busca na metade esquerda
	}

	return BinarySearch(x, listElements[mid+1:]) // Busca na metade direita
}

func main() {
	x := 3

	listNumbers := []int{1, 2, 3, 7, 9, 21, 30, 19, 23, 59}
	sort.Ints(listNumbers)

	result, err := BinarySearch(x, listNumbers)
	if err != nil {
		log.Println("Erro:", err)
		return
	}
	fmt.Println("Resultado encontrado:", result)
}
