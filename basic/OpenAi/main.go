package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatalln("Missing API KEY")
	}

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	resp, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:    []string{"" +
					"You: Do you like play chess?\n" +
					"Friend:"},

		MaxTokens: gpt3.IntPtr(60),
		Stop:      []string{"You:"},
		Echo:      true,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Choices[0].Text)
}