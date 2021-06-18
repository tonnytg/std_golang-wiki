package main

import "fmt"

type Pessoa string

var aluno1 Pessoa
var x int
var y bool
var z float64

func main() {

	aluno1 = "Antonio"

	fmt.Println(aluno1) // imprime o valor que eu criei
	fmt.Println(x) // imprime 0
	fmt.Println(y) // imprime false
	fmt.Println(z) // imprime 0.0
}
