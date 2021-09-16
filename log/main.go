// https://www.honeybadger.io/blog/golang-logging/
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Printf("Error opening")
	}
	// Define o arquivo de saída
	log.SetOutput(file)
	// Define as flags para formatação
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
	// Registra um evento
	log.Println("log inside file")
}
