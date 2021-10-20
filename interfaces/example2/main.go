package main

import "log"

type Database interface {
	GetUser() string
}

type DefaultDatabases struct {}

func (db DefaultDatabases) GetUser() string {
	return "test"
}

type Greeter interface {
	Greet(userName string)
}

type NiceGreeter struct {}

func (ng NiceGreeter) Greet(userName string) {
	log.Println("Greet", userName)
}

type Program struct {
	Db Database
	Greeter Greeter
}

func (p Program) Execute() {
	user := p.Db.GetUser()
	p.Greeter.Greet(user)
}

func main() {
	db := DefaultDatabases{}
	greeter := NiceGreeter{}
	p := Program{
		Db: db,
		Greeter: greeter,
	}
	p.Execute()
}
