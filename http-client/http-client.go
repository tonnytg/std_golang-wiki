package http_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetSite(url string) error{
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	context, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return errors.New("403")
	}
	fmt.Printf("%s", context)
	return nil
}

func PostSite(url string) error{
	content := "application/json"
	values := map[string]string{
		"job": "cientist",
		"name": "antonio",
	}
	jsonValue, _ := json.Marshal(values)
	body := bytes.NewBuffer(jsonValue)

	resp, err := http.Post(url, content, body)
	if err != nil {
		log.Println(err)
		return errors.New("403")
	}
	fmt.Println(resp)
	return nil
}