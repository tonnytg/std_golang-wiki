package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Criando uma variável de tipo desconhecido em tempo de compilação
	var x interface{}
	x = 42

	// Obtendo o tipo da variável em tempo de execução
	fmt.Println("Tipo de x:", reflect.TypeOf(x))

	// Obtendo o valor da variável em tempo de execução
	fmt.Println("Valor de x:", reflect.ValueOf(x))

	// Verificando se o valor de x é do tipo int
	if reflect.TypeOf(x).Kind() == reflect.Int {
		fmt.Println("x é um número inteiro.")
	}

	// Convertendo o valor de x para int
	if v, ok := x.(int); ok {
		fmt.Println("Valor de x como int:", v)
	}
}
