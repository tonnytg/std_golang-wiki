package main

import "fmt"

// Calc a value we don't know
func Calc(w, x, y float64) float64 {
	return ( y * x) / w
}

func main() {
	fmt.Println("Result: ", Calc(2, 300, 5))
}
