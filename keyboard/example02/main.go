package main

import (
	"bufio"
	"fmt"
	"github.com/tonnytg/talkalone/pkg/httpRequest"
	"os"
	"strings"
)

// doc
//curl https://api.openai.com/v1/files \
//-H "Authorization: Bearer $OPENAI_API_KEY" \
//-F purpose="answers" \
//-F file="@myfile.jsonl"

func CallFriend(message string) []byte {
	m, _ := httpRequest.GetRequest("https://api.openai.com/v1/files", []byte(message))
	return m
}

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

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}

	}
}
