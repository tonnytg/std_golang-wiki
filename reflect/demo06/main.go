package main

import (
	"fmt"
	"reflect"
)

func isNullOrEmpty(value interface{}) bool {
	val := reflect.ValueOf(value)

	// Verificando se o valor é nulo
	if !val.IsValid() || val.IsNil() {
		return true
	}

	// Verificando se o valor é vazio (para tipos de slice, mapa, etc.)
	switch val.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return val.Len() == 0
	default:
		return false
	}
}

func main() {
	var a int
	var b string
	c := []int{}
	d := map[string]int{}

	fmt.Println("a é nulo ou vazio:", isNullOrEmpty(a))
	fmt.Println("b é nulo ou vazio:", isNullOrEmpty(b))
	fmt.Println("c é nulo ou vazio:", isNullOrEmpty(c))
	fmt.Println("d é nulo ou vazio:", isNullOrEmpty(d))
}
