package main

import "fmt"

func main() {
	fmt.Println("Start Go")

	var pessoa = map[string]string{
		"nome":  "João",
		"idade": "30",
	}

	for i, v := range pessoa {
		fmt.Println(i, v)
	}
}
