package main

import (
	"fmt"
	"log"
)

func main()  {

	var name string = "Antonio"

	fmt.Println("Name:", name)
	fmt.Printf("Repita o nome acima:")
	fmt.Scanf("%s", &name)

	if name == "Antonio" {
		log.Printf("Name is Antonio")
	} else {
		log.Printf("Status Code: 422, name is diferent: %v", name)
	}

}
