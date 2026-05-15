package LIFO

import "fmt"

func Lifo(words ...string) {
	var stack []string

	for _, word := range words {
		//total += word
		stack = append(stack, word)
	}

	for len(stack) > 0 {
		n := len(stack) - 1 // Top element
		fmt.Print(stack[n], " ")

		stack = stack[:n] // Pop
	}
}
