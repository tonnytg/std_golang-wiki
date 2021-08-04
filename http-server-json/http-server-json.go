package http_server_json

import (
	"log"
	"net/http"
	"os"
	"time"
)

func CallServer() {

	server := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      http.DefaultServeMux,
		ErrorLog:     log.New(os.Stderr, "server log:", log.Lshortfile),
	}

	http.Handle("/", http.FileServer(http.Dir("/tmp")))

	err := server.ListenAndServe()
	if err != nil {
		log.Println("listen error:", err)
	}
}
