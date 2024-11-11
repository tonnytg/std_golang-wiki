package main

import (
	"fmt"
)

// QuickSort implementa o algoritmo QuickSort
func QuickSort(arr []int) []int {
	// Caso base: se o array tiver 0 ou 1 elemento, já está ordenado
	if len(arr) <= 1 {
		return arr
	}

	// Escolhe um pivô (usamos o primeiro elemento, mas pode ser qualquer um)
	pivot := arr[0]
	var left, right []int

	// Divide o array em dois subarrays: menores e maiores que o pivô
	for _, value := range arr[1:] {
		if value < pivot {
			left = append(left, value)
		} else {
			right = append(right, value)
		}
	}

	// Recursivamente ordena os subarrays e junta tudo
	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}

func main() {
	arr := []int{9, 3, 5, 2, 7, 6, 1, 8, 4}

	fmt.Println("Array original:", arr)

	sorted := QuickSort(arr)

	fmt.Println("Array ordenado:", sorted)
}

