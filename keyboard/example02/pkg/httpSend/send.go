package httpSender

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func SendWithArgs(msg string) ([]byte, error) {

	token := os.Getenv("OPENAI_API_KEY")
	if token == "" {
		fmt.Println("invalid token, you need export GCP_TOKEN")
		os.Exit(1)
	}

	url := "https://api.openai.com/v1/engines/davinci-instruct-beta/completions"
	body := io.Reader(strings.NewReader(`{
  "prompt": "How old are you?:\n\n1.",
  "max_tokens": 64,
  "temperature": 0.8,
  "frequency_penalty": 0.0,
  "presence_penalty": 0.0,
  "top_p": 1.0,
  "stop": ["\n\n"]
}
`))
	req, err := http.NewRequest("POST", url, body)

	// Header with Authorization
	bearer := "Bearer " + token
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	r := regexp.MustCompile(`20([0-9])`)
	if !r.Match([]byte(string(resp.StatusCode))) {
		fmt.Println("statusCode:", resp.StatusCode)
	}

	data, _ := ioutil.ReadAll(resp.Body)
	return data, nil
}
