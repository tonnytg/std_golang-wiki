package main

import "fmt"

func main() {
	x := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("x: %v\n", x)
	fmt.Printf("x: %T\n", x)
	fmt.Printf("x: %v\n", len(x))
	for i := range x {
		fmt.Printf("%v ", x[i])
	}
	fmt.Println("\n")

	xx := make([]int, 2)
	fmt.Printf("xx: %v\n", xx)
	fmt.Printf("xx: %T\n", xx)
	fmt.Printf("xx: %v\n", len(xx))
	for i := range x {
		xx = append(xx, x[i])
	}

	fmt.Println("x + xx:", xx)

}
