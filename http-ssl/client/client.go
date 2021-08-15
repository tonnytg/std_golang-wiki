package client

import (
	"crypto/tls"
	"errors"
	"log"
	"net/http"
)

func Client() {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	res, err := client.Get("https://tonnytg.com:8443")
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatal(errors.New("status code must be 200 and got"), res.StatusCode)
	}
}
