package LIFO

import "fmt"

func Lifo() {
	var stack []string

	stack = append(stack, "world!") // Push
	stack = append(stack, "Hello ")

	for len(stack) > 0 {
		n := len(stack) - 1 // Top element
		fmt.Print(stack[n])

		stack = stack[:n] // Pop
	}
}
