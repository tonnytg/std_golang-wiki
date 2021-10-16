package main

import "fmt"

type Person struct {
	Name string
	LastName string
	Age int
	PhoneNumber string
	Address string
}

func (p *Person) SetName(n, l string) {
	p.Name = n
	p.LastName = l
}

func (p *Person) SetAge(a int) {
	p.Age = a
}

func (p *Person) SetNumber(t string) {
	p.PhoneNumber = t
}

func (p *Person) SetAddress(a string) {
	p.Address = a
}


func main() {
	p := Person{}
	p.SetName("Test", "Silva")
	p.SetAge(10)
	p.SetNumber("5511999999999")
	p.SetAddress("Street Washington, USA")
	fmt.Println("Info of Person:", p)
}
