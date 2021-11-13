package main

import (
	"bufio"
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
	ReadLines(files[0])
}

func ReadLines(file string) {
	f, err := os.Open(file)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    lines := make([]string, 0)
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    fmt.Println(lines)
}