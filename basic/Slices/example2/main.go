package main

import "fmt"

func main() {

	fmt.Println("Slice com 3 posições definidas")
	mySlice := make([][]int, 3)
	fmt.Println(len(mySlice))
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		mySlice[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			mySlice[i][j] = i + j
		}
	}
	fmt.Println("Slice: ", mySlice)
}
