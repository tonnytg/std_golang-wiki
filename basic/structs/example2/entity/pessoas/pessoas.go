package pessoas

type Pessoa struct {
	Name string
	Idade int
}

type Pessoas struct {
	Pessoas []Pessoa
}

func (pp *Pessoas) AddPessoa(p Pessoa) *Pessoas {
	pp.Pessoas = append(pp.Pessoas, p)
	return pp
}

func NewPessoa() Pessoa {
	p := Pessoa{}
	p.Name = "Alguem"
	p.Idade = 99
	return p
}