package http_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Jobs struct {
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	Job       string    `json:"job"`
	ID        string    `json:"id"`
}

func GetSite(url string) error{
	res, err := http.Get(url)
	if err != nil || res.StatusCode == 500 {
		return errors.New("invalid url or return status code 500")
	}
	context, err := io.ReadAll(res.Body)
	var j []Jobs
	err = json.Unmarshal(context, &j)
	if err != nil {
		return errors.New("error to unmarshalling")
	}

	for i := range j {
		fmt.Printf("Job[%v]: %s\tName:[%v]\n",i ,j[i].Job, j[i].Name)
	}

	return nil
}

func PostSite(url string) error{
	content := "application/json"
	values := map[string]string{
		"job": "scientist",
		"name": "antonio",
	}
	jsonValue, _ := json.Marshal(values)
	body := bytes.NewBuffer(jsonValue)

	resp, err := http.Post(url, content, body)
	if err != nil {
		log.Println(err)
		return errors.New("error to posting")
	}

	fmt.Println(resp.Status)
	return nil
}