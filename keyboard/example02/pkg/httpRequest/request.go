package httpRequest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func PostRequest(url string, message string) ([]byte, error) {

	token := os.Getenv("OPENAI_API_KEY")
	if token == "" {
		fmt.Println("invalid token, you need export OPENAI_API_KEY")
		os.Exit(1)
	}

	bearer := "Bearer " + token

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(message)))
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

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
