package main

import (
	"fmt"
)

// BubbleSort implementa o algoritmo Bubble Sort
func BubbleSort(arr []int) {
	n := len(arr)

	// Loop para percorrer todo o array
	for i := 0; i < n-1; i++ {
		// Comparar elementos adjacentes
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Trocar se estiver fora de ordem
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{9, 3, 5, 2, 7, 6, 1, 8, 4}

	fmt.Println("Array original:", arr)

	BubbleSort(arr)

	fmt.Println("Array ordenado:", arr)
}

