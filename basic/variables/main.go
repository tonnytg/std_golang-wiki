package main

import (
	"errors"
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

func MyName( x string ) (bool, error) {

	if x == "Antonio" {
		return true, nil
	}
	return false, errors.New("you inserted wrong value")
}

// main it is where the magic happens
func main()  {
	fmt.Printf("Hello friend %v! \n", student1.Name )
	fmt.Printf("Friend you have %v years old! \n", student1.Age )

	var name, err = MyName("Maria")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(name)
}



