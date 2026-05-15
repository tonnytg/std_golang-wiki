package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// get url and convert to base64
	url := "https://www.google.com"
	base64Url := base64.StdEncoding.EncodeToString([]byte(url))
	fmt.Println(base64Url)

	// TODO: split base64Url 4 slices and after decode each slice


}
