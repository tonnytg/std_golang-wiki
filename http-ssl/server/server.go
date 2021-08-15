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
		Addr:         ":8443",
		Handler:      http.DefaultServeMux,
		ErrorLog:     log.New(os.Stderr, "server log:", log.Lshortfile),
	}

	http.HandleFunc("/", PersonHandler)

	err := server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		log.Println("listen error:", err)
	}
}
