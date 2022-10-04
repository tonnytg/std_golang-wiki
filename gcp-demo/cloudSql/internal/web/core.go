package web

import (
	"context"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Core is a function that makes a request to a URL and returns the response body
// Method is GET, and the timeout is 10 seconds (Default)
func Core(Method string, url string, timeOut int, data []string) ([]byte, error) {

	// client run everything
	client := &http.Client{
		CheckRedirect: nil,
	}

	// context help us to control timeout for request
	ctx := context.Background()
	if timeOut < 10 {
		timeOut = 10
	}

	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeOut)*time.Second)
	defer cancel()

	// convert data to io.Reader
	var dataIo io.Reader
	for _, v := range data {
		dataIo = strings.NewReader(v)
	}

	// request with NewRequest permit add headers
	request, err := http.NewRequestWithContext(ctx, Method, url, dataIo)
	if err != nil {
		return nil, err
	}

	token := os.Getenv("GCP_TOKEN")
	bearer := "Bearer " + token

	request.Header.Add("Content-Type", "application/json")
	// Headr Authorization
	request.Header.Add("Authorization", bearer)
	request.Method = Method

	// execute call and return *http.Response type
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	// clone request after process
	defer response.Body.Close()

	// convert information to slice of byte
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
