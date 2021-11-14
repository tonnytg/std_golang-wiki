package main

import (
	"struct2/entity/pessoas"
	"fmt"
	"github.com/jaswdr/faker"
)

func main() {

	faker := faker.New()

	p := pessoas.NewPessoa()

	var gente pessoas.Pessoas

	for i := 0; i < 10; i++ {
		p.Name = faker.Person().Name()
		p.Idade = faker.IntBetween(18, 60)
		gente.AddPessoa(p)
	}

	for _, pessoa := range gente.Pessoas {
		fmt.Println(pessoa)
	}
}
