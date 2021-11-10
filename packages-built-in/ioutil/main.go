package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	r, err := ioutil.ReadFile("/etc/passwd")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%s", r)
}
