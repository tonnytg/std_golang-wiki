package main

import (
	"fmt"
)

type Agents struct {
	Name string
	Surname string
	Age int
	Country string
}

func main()  {

	agent007 := Agents{
		"Bond",
		"James",
		42,
		"United Kingdom",
	}

	agent006 := Agents{
		"Alec",
		"Trevelyan",
		48,
		"United Kingdom",
	}

	fmt.Println(agent007)
	fmt.Printf("Name: %s\t Age: %d\n", agent007.Name, agent007.Age)
	fmt.Printf("Name: %s\t Age: %d\n", agent006.Name, agent006.Age)
}