package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("names.txt")
	defer f.Close()
	if err != nil {
		fmt.Println("Deu erro:", err)
		return
	}

	r := strings.NewReader("James Bond")
	io.Copy(f, r)
}
