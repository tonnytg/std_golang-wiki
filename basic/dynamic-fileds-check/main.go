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

type ValidValues struct {
	JobID      string       `json:"JobID"`
	Parameters []Parameters `json:"Parameters"`
	Labels     []Labels     `json:"Labels"`
}

type Fields struct {
	ID         int          `json:"Id"`
	Job        string       `json:"Job"`
	Name       string       `json:"Name"`
	Parameters []Parameters `json:"Parameters"`
	Labels     []Labels     `json:"Labels"`
}

func GetSkel() (ValidValues, error) {
	o, _ := os.Open("./skel.json")
	defer o.Close()

	content, err := io.ReadAll(o)
	if err != nil {
		panic(err)
	}

	p := ValidValues{}
	err = json.Unmarshal(content, &p)
	if err != nil {
		panic(err)
	}

	return p, nil
}

func GetValues() (Fields, error) {
	o, _ := os.Open("./values.json")
	defer o.Close()

	content, err := io.ReadAll(o)
	if err != nil {
		panic(err)
	}

	f := Fields{}
	err = json.Unmarshal(content, &f)
	if err != nil {
		panic(err)
	}

	return f, nil
}

func main() {
	fmt.Println("Test Dynamic fields read by outside JSON")

	skel, err := GetSkel()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Skeleton:", skel)

	values, err := GetValues()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Values:", values)
}
