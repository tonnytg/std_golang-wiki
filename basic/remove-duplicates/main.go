package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var values []string
	reader := bufio.NewReader(os.Stdin)

	for {
		value, _ := reader.ReadString('\n')

		fmt.Println("Inser:", value)
		for i := 0; i < len(values); i++ {
			fmt.Println("Check:", values[i])
			if value == values[i] {
				fmt.Println("DUB")
			}
		}
		values = append(values, value)

		fmt.Println(values)
	}
}
