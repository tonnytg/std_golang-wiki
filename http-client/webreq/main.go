package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tonnytg/webreq"
)

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func main() {
	url := "https://610aa52552d56400176afebe.mockapi.io/api/v1/friendlist"
	timeOut := 20
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	body, err := webreq.Get(url, headers, timeOut)
	if err != nil {
		fmt.Println(err)
	}
	bodyString := string(body)
	if bodyString == "" {
		fmt.Println("body is empty")
	}
	pretty, err := PrettyString(bodyString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pretty)
}
