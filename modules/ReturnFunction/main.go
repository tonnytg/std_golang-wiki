package main

import "fmt"

func main() {
	z := foo()
	z("ola mundo")
}

func foo() func (string) {
	return func (z string) {
		fmt.Println(z)
	}
}