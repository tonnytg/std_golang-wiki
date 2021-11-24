package main

import "fmt"

// Calc a value we don't know
func Calc(a, b, c float64) float64 {
	return (b * c) / a
}

func main() {
	fmt.Println("Result: ", Calc(33, 100, 15))
}
