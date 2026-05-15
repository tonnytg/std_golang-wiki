package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func addServedHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Served-Date", time.Now().String())
}

func makeRequestHandler(middleware http.HandlerFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		middleware(w, r)

		w.Write([]byte("OK"))
	}
}

func Handle(w http.ResponseWriter, r *http.Request) {
	var input []byte

	if r.Body != nil {
		defer r.Body.Close()

		body, _ := ioutil.ReadAll(r.Body)

		input = body
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello world, input was:\n %s", string(input))))
}

func main() {
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8080),
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: 1 << 20, // Max header of 1MB
	}

	next := addServedHeader
	http.HandleFunc("/", makeRequestHandler(next))
	http.HandleFunc("/v1", Handle)
	log.Fatal(s.ListenAndServe())
}