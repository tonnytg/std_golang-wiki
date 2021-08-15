package server

import (
	"log"
	"net/http"
	"os"
	"time"
)

func Server() {

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
