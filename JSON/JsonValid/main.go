package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Test struct {
	Name string `json:"name"`
}

type Mandatory struct {
	Values []struct {
		Name  string `json:"Name"`
		Value string `json:"Value"`
	} `json:"Values"`
}

func Read(file string) []byte {
	// read file
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer f.Close()

	body, _ := io.ReadAll(f)
	return body
}

func main() {

	data := Read("input.json")
	var t Test
	err := json.Unmarshal(data, &t)
	if err != nil {
		fmt.Println(err)
		return
	}

	mandatory := Read("mandatory.json")

	var m Mandatory

	err = json.Unmarshal(mandatory, &m)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("M", m)
	fmt.Println("T", t.Name)

	for _, v := range m.Values {
		fmt.Println(v.Name, v.Value)
		if v.Name == t.Name {
			fmt.Println("Found")
		}
	}
}
