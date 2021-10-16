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

// Veiculo interfaces sign who has Liga()
type Veiculo interface {
	Liga()
	SetMarca()
}

func (c *Carro) SetMarca() {
	c.Marca = "VW"
}

func (c *Carro) Liga() {
	c.Estado = true
}

func (m *Moto) Liga() {
	m.Estado = false
}

func main() {
	c := Carro{}
	c.Liga()
	c.SetMarca()
	fmt.Println(c.Estado)
	fmt.Printf("Format: %T\n", c)

	m := Moto{}
	m.Liga()
	fmt.Println(m.Estado)
	fmt.Printf("Format: %T\n", m)

	LigaVeiculo(&c)
}

func LigaVeiculo(v ...Veiculo) {
	for i, j := range v {
		fmt.Println(i, j)
	}
}
