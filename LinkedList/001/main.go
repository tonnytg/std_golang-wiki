package main

type Pessoa struct {
	Name     string
	LastName string
	Age      int
}

type Node struct {
	Pessoa   Pessoa
	Lastname string
	Next     *Node
}

func main() {
	pessoa1 := Pessoa{
		Name:     "Rafael",
		LastName: "Oliveira",
		Age:      30,
	}

	node := Node{
		Pessoa:   pessoa1,
		Lastname: "Oliveira",
		Next:     nil,
	}

	pessoa2 := Pessoa{
		Name:     "João",
		LastName: "Oliveira",
		Age:      30,
	}

	node2 := Node{
		Pessoa:   pessoa,
		Lastname: "Oliveira",
		Next:     nil,
	}

	node.Next = &node2
}
