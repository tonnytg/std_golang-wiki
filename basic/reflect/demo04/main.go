package main

import (
	"fmt"
	"reflect"
)

func add(a, b interface{}) interface{} {
	valA := reflect.ValueOf(a)
	valB := reflect.ValueOf(b)

	// Verificando se os tipos são numéricos
	if valA.Kind() != reflect.Int && valA.Kind() != reflect.Float64 {
		fmt.Println("O primeiro argumento não é numérico.")
		return nil
	}
	if valB.Kind() != reflect.Int && valB.Kind() != reflect.Float64 {
		fmt.Println("O segundo argumento não é numérico.")
		return nil
	}

	// Convertendo valores para float64 para evitar problemas de tipo
	floatA := valA.Float()
	floatB := valB.Float()

	// Realizando a operação de adição
	result := floatA + floatB

	// Retornando o resultado como interface{} para manter a generalidade
	return result
}

func main() {
	fmt.Println(add(10, 20))
	fmt.Println(add(3.14, 2.5))
	fmt.Println(add(5, "abc")) // Irá imprimir mensagens de erro
}
