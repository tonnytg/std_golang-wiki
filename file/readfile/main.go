package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Read File")

	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	fileLines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for i, v := range fileLines {
		fmt.Println(i, v)
	}
}
