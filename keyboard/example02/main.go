package main

import (
	"bufio"
	"fmt"
	"github.com/tonnytg/talkalone/pkg/httpSend"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("type Hi to start to talk")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		answer, _ := httpSender.SendWithArgs(text)
		fmt.Println(string(answer))
	}
}
