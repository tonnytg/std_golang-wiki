package main

import "fmt"



func main() {
	fmt.Println("Start example 3 of Interface")

	// Accept Generic values
	var values  []interface{}
	values = append(values, "Antonio")
	values = append(values, 30)
	fmt.Println(Join(values))
}

func Join(x []interface{}) (result string){

	// loop in slice of interface
	for _, v := range x {
		// concat interface values in string result
		result += fmt.Sprintf("%v ", v)
	}
	// return string
	return result
}
