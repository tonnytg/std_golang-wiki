package main

import "encoding/base64"

func main() {
	// get url and convert to base64
	url := "https://www.google.com"
	base64Url := base64.StdEncoding.EncodeToString([]byte(url))
}
