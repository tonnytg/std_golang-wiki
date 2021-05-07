package main

import (
	"fmt"
)

// Student is the format to build a student card
type Student struct {
	Name string
	Age int32
}

// student1 is variable static of first student of class
var student1 = Student {
	"Antonio",
	32,
}

// main it is where the magic happens
func main()  {
	fmt.Printf("Hello friend %v! \n", student1.Name )
	fmt.Printf("Friend you have %v years old! \n", student1.Age )
}