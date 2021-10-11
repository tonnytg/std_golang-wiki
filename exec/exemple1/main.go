package main

import (
	"fmt"
	"os/exec"
)

func main() {

	prg := "ls"

	arg1 := "-l"

	cmd := exec.Command(prg, arg1)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}
