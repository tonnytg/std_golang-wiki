//https://go.dev/doc/effective_go#maps
package main

import "fmt"

func CheckIndice()  (string, bool){
	return "working", true
}

func main() {
	version := "1.0.0"
	fmt.Printf("Init %s\n", version)

	m := map[int]int{1: 0, 2: 1, 3: 2}

	fmt.Println(m[1])

	if indice, ok := CheckIndice(); ok {
		fmt.Println("it is", indice)
	}
}
