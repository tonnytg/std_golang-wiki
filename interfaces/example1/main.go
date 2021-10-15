package main

import "fmt"

type Carro struct {
	Marca string
	Modelo string
	Ano int
	Estado bool
}

type Moto struct {
	Marca string
	Modelo string
	Ano int
	Estado bool
}

func (c *Carro) Liga() {
	c.Estado = true
}

func (m *Moto) Liga() {
	m.Estado = true
}

// Veiculo interfaces sign who has Liga()
type Veiculo interface {
	Liga()
}

func main() {
	c := Carro{}
	c.Liga()
	fmt.Println(c.Estado)
	fmt.Printf("Format: %T\n", c)

	m := Moto{}
	m.Liga()
	fmt.Println(m.Estado)
	fmt.Printf("Format: %T\n", m)
}
