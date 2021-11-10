package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	r, err := ioutil.ReadFile("/etc/passwd")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%s", r)
	for i, v := range strings.Split(string(r), ":") {
		fmt.Printf("%d: %s", i, v)
	}

}
