package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		log.Println("No file name given")
		os.Exit(1)
	}

	files := make(map[int]string)

	for i, filename := range os.Args[1:] {
		fmt.Println(filename)
		files[i] = filename
	}
	fmt.Println(files)
}