package main

import (
	"fmt"
)

func main() {
	fmt.Println("Demo Map")

	// create a map wit name of fruits
	fruits := map[string]int{
		"apple":  10,
		"banana": 20,
		"orange": 30,
	}
	// sort fruits
	for k, v := range fruits {
		fmt.Println(k, v)
	}
	// get value of key "apple"
	fmt.Println(fruits["apple"])
	// check if key "apple" is exist
	_, prs := fruits["apple"]
	fmt.Println(prs)

}
