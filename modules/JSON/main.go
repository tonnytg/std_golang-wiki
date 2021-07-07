package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome string
	Idade int
	Profissao string
}

func main() {
	fmt.Println("Let's Go!")
	pessoa := Pessoa { "Antonio", 32, "Analista"}
	encoder := json.Marshal(pessoa)
	encoder.Encode(pessoa)

	fmt.Printf("JSON to Slice of Bytes[]\n")
	fmt.Println(json.Marshal(pessoa))

	fmt.Printf("\nSlice of Bytes[] to JSON\n")
	fmt.Println(json.Unmarshal(json.Marshal(pessoa)))
}