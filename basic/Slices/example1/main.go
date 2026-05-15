package main

import (
	"fmt"
)

type Item struct {
	Value int
}

func alter(t *Item) {
	(*t).Value = 100
}

func main() {

	items := []Item{Item{0}, Item{1}}

	for i := range items {
		alter(&items[i])
	}

	fmt.Println(items) // Output is still [{0} {1}]

}