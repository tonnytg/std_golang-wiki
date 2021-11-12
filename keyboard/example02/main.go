package main

import (
	"bufio"
	"fmt"
	"github.com/tonnytg/talkalone/pkg/httpRequest"
	"os"
	"strings"
)

// AnswerOpenAI it is a json response from OpenAI
type AnswerOpenAI struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
}

// QuestionOpenAI it is a question to send a Open AI
type QuestionOpenAI struct {
	Prompt           string  `json:"prompt"`
	Temperature      float64 `json:"temperature"`
	MaxTokens        int     `json:"max_tokens"`
	TopP             float64 `json:"top_p"`
	FrequencyPenalty float64 `json:"frequency_penalty"`
	PresencePenalty  float64 `json:"presence_penalty"`
	Stop             string  `json:"stop"`
}

func (q QuestionOpenAI) NewQuestion() {

	q.Prompt = `"You: What have you been up to?\nFriend: Watching old movies.\nYou: Did you watch anything interesting?\nFriend:",`
	q.FrequencyPenalty = 0.5
	q.PresencePenalty = 0.0
	q.TopP = 1.0
	q.MaxTokens = 60
	q.Temperature = 0.4
	q.Stop = `["You:"]`

}

func CallFriend(message string) []byte {
	m, _ := httpRequest.PostRequest("https://api.openai.com/v1/files", message)
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

		if strings.Compare("Hi", text) == 0 {
			answer := CallFriend("Hi")
			fmt.Println(string(answer))
		}

	}
}
