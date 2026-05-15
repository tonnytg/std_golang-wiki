package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/tasks", handleTasks)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getTasks(w, r)
	} else if r.Method == "POST" {
		createTask(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	allTasks := GetTasks()

	json.NewEncoder(w).Encode(allTasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task

	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	AddTask(newTask)

	w.WriteHeader(http.StatusCreated)
}
