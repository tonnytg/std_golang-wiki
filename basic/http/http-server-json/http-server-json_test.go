package http_server_json_test

import (
	"fmt"
	http_server_json "github.com/tonnytg/std-golang-wiki/http-server-json"
	"io"
	"log"
	"net/http"
	"testing"
)

func TestServeMyJson(t *testing.T) {
	go http_server_json.CallServer()

	// if you need test browser
	//time.Sleep(10 * time.Second)
	url := "http://localhost:8081"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	context, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", context)
}
