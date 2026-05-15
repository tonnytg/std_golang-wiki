package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Pessoa struct {
	Nome string
	Idade int
	Profissao string
}

func main() {

	pessoa := Pessoa { "Antonio", 32, "Analista"}
	encoder, _ := json.Marshal(pessoa)

	fmt.Printf("JSON to Slice of Bytes[]\n")
	fmt.Println(encoder)

	fmt.Printf("\nSlice of Bytes[] to JSON\n")
	var decoder Pessoa
	err := json.Unmarshal(encoder, &decoder)
	if err != nil {
		log.Println("error:", err)
	}
	fmt.Println("decode", decoder)
}