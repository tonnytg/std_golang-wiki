package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var Fatia []string

func PersonHandler(w http.ResponseWriter, req *http.Request) {

	q := req.URL.Query()
	msg := q.Encode()

	if msg != "" {
		fmt.Println("Message:", msg)
		Fatia = append(Fatia, msg)
		fmt.Println(Fatia)
	}

	w.Write([]byte("hi"))
}

func WebServer() {

	server := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      http.DefaultServeMux,
		ErrorLog:     log.New(os.Stderr, "server log:", log.Lshortfile),
	}

	http.HandleFunc("/", PersonHandler)

	err := server.ListenAndServe()
	if err != nil {
		log.Println("listen error:", err)
	}
}

func main() {
	go Lifo()
	WebServer()
}

func Lifo() {
	for {
		if len(Fatia) > 0 {
			n := len(Fatia) - 1 // Top element
			fmt.Print("Fatiado:", Fatia[n], " ")

			Fatia = Fatia[:n] // Pop
		}
	}
}