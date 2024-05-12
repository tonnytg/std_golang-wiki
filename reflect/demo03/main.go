package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	City string
}

func printFields(data interface{}) {
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	if val.Kind() != reflect.Struct {
		fmt.Println("O argumento fornecido não é uma struct.")
		return
	}

	fmt.Println("Campos da struct:")
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name
		fmt.Printf("%s: %v\n", fieldName, field.Interface())
	}
}

func main() {
	p := Person{"Alice", 30, "New York"}
	printFields(p)
}
