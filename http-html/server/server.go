package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type Info struct {
	ID	int
	Title string
	Content string
}

func PersonHandler(w http.ResponseWriter, req *http.Request) {
	news := []Info{
		{1, "Vaga", "Test"},
		{2, "Vaga2", "Test2"},
		{3, "Vaga3", "Test3"},
	}
	t, err := template.ParseFiles("./server/templates/home.html")
	if err != nil {
		log.Panic(err)
	}
	t.Execute(w, news)
	fmt.Println(news[0].ID)

	w.WriteHeader(http.StatusOK)
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
