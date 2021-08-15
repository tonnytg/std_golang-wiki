package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string
	Age int
	Job string
}

func PersonHandler(w http.ResponseWriter, req *http.Request) {

	person := Person { "Antonio", 32, "Analyst"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
	fmt.Println(w.Header())
}