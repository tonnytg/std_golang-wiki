package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	program := "git"
	path, err := exec.LookPath(program)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(path)

	args := []string{"config", "--list"}

	cmd := exec.Command(program, args...)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Print(string(stdout))
}
