package main

import (
	"errors"
	"fmt"
	"log"
)

func MsgErrorLog(msg string) error {
	log.Println("Error: %s\n", msg)
	return msg
}

// doError emit fake error code 404
func doError() (string, error) {
	return "", errors.New("404")
}
// doNoError emit success code 200
func doNoError() (string, error) {
	return "200", nil
}

func doFmtError() error {
	errCode:= 401
	return fmt.Errorf("This is my error code: %d", errCode)
}

func main()  {
	fmt.Println("Let's Go!")
	resp, err := doError()
	if err != nil {
		log.Printf("There was an error: %v\n", err)
	} else {
		fmt.Println("My message:", resp)
	}

	resp, err = doNoError()
	if err != nil {
		log.Printf("This should not print")
	}
	log.Println("My response:", resp)

	err = doFmtError()
	if err != nil {
		log.Printf("There was an error: %v\n", err)
	}
}
