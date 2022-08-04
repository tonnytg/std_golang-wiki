package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Parameters struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type Labels struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type CheckParameters struct {
	JobID      string       `json:"JobID"`
	Parameters []Parameters `json:"Parameters"`
	Labels     []Labels     `json:"Labels"`
}

type Fields struct {
	ID         int          `json:"Id"`
	Name       string       `json:"Name"`
	Parameters []Parameters `json:"Parameters"`
	Labels     []Labels     `json:"Labels"`
}

func GetParameters() error {
	file := "./test.json"
	o, _ := os.Open(file)
	defer o.Close()

	content, err := io.ReadAll(o)
	if err != nil {
		panic(err)
	}

	p := CheckParameters{}
	err = json.Unmarshal(content, &p)
	if err != nil {
		panic(err)
	}
	for i, v := range p.Parameters {
		fmt.Println(i, v.Name, v.Value)
	}

	return fmt.Errorf("Not implemented")
}

func main() {
	fmt.Println("Test Dynamic fields read by outside JSON")
	var f Fields
	f.ID = 1
	f.Name = "Test"
	for i, v := range f.Parameters {
		fmt.Println("Check", i, v)
	}
	fmt.Println(f)

	err := GetParameters()
	if err != nil {
		log.Println(err)
	}
}
