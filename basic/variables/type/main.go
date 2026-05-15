package main

import "fmt"

type Pessoa string

var aluno1 Pessoa
var x int
var y bool
var z float64

const (
	a = iota
	b = iota
	_ = iota
)

func main() {

	bit  := 1
	bit2 := bit << 1

	aluno1 = "Antonio"

	fmt.Println(aluno1) // imprime o valor que eu criei
	fmt.Println(x) // imprime 0
	fmt.Println(y) // imprime false
	fmt.Println(z) // imprime 0.0

	fmt.Println(a, b) // iota

	fmt.Println(bit, bit2)

}
