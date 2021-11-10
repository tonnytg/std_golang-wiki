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
	count := 0
	for _, v := range strings.Split(string(r), ":") {
		switch count {
		case 0:
            fmt.Printf("User: %s ", v)
			count++
		case 1:
			fmt.Printf("UID: %s ", v)
            count++
		case 2:
			fmt.Printf("GID: %s ", v)
            count++
		case 3:
            fmt.Printf("Groups: %s ", v)
            count++
		case 4:
            fmt.Printf("Home: %s ", v)
            count++
		case 5:
            fmt.Printf("Shell: %s ", v)
			count = 0
		//case 6:
        //    fmt.Printf("\n")
        //    count = 0
		}
	}
	fmt.Println()
}
