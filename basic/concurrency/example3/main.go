package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func Writer(msg string, chw chan string) {
	if len(chw) == 0 {
		fmt.Println("Write: channel empty")
	}
	chw <- msg
}

func Reader(chr chan string) {
	fmt.Printf("You insert: %s", <- chr)
	if len(chr) == 0 {
		fmt.Println("Read: channel empty")
	}
}

func main() {
	c := make(chan string)
	fmt.Println("Chat init")

	// Chat
	msg := ""
	for msg != "exit"{
		input := bufio.NewReader(os.Stdin)
		time.Sleep(1 * time.Second)
		fmt.Printf("Enter message: ")
		msg, _ := input.ReadString('\n')
		go Writer(msg, c)
		if msg == "exit\n" {
			os.Exit(0)
		}

		//c <- msg

		go Reader(c)
	}
}
