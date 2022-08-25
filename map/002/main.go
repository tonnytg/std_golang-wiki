package main

import "fmt"

var pessoa map[string]string

func main() {
	fmt.Println("vim-go")

	pessoa = map[string]string{"nome": "João", "idade": "20"}

	pessoa1 := make(map[string]string)
	pessoa1["nome"] = "João"
	pessoa1["sobrenome"] = "Silva"

	pessoa2 := map[string]string{
		"nome":      "João",
		"sobrenome": "Silva",
	}

	fmt.Println(pessoa)
	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
}
