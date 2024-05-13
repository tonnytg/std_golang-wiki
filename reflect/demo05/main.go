package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name    string
	Age     int
	Country string
}

func getFieldValues(data interface{}, fields ...string) []interface{} {
	val := reflect.ValueOf(data)

	// Verificando se o valor passado é uma estrutura (struct)
	if val.Kind() != reflect.Struct {
		fmt.Println("O argumento fornecido não é uma estrutura (struct).")
		return nil
	}

	// Criando um slice para armazenar os valores dos campos especificados
	fieldValues := make([]interface{}, len(fields))

	// Iterando sobre os campos especificados
	for i, fieldName := range fields {
		// Obtendo o valor do campo
		fieldValue := val.FieldByName(fieldName)

		// Verificando se o campo existe
		if !fieldValue.IsValid() {
			fmt.Printf("Campo %s não encontrado na estrutura.\n", fieldName)
			fieldValues[i] = nil
		} else {
			// Adicionando o valor do campo ao slice de valores
			fieldValues[i] = fieldValue.Interface()
		}
	}

	return fieldValues
}

func main() {
	person := Person{Name: "Alice", Age: 30, Country: "USA"}

	// Obtendo os valores dos campos "Name" e "Age" da estrutura Person
	values := getFieldValues(person, "Name", "Age", "City")

	// Imprimindo os valores dos campos
	for i, fieldName := range []string{"Name", "Age", "City"} {
		fmt.Printf("%s: %v\n", fieldName, values[i])
	}
}
