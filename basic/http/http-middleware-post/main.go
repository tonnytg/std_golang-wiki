package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type RequestData struct {
	Data string `json:"data"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Lendo o corpo da requisição
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			log.Println("Error reading request body:", err)
			return
		}

		var requestData RequestData
		err = json.Unmarshal(body, &requestData)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			log.Println("Error unmarshaling JSON:", err)
			return
		}

		log.Println("Data received from POST:", requestData.Data)

		// Retornando a resposta como JSON
		response, err := json.Marshal(requestData)
		if err != nil {
			http.Error(w, "Failed to marshal response JSON", http.StatusInternalServerError)
			log.Println("Error marshaling response JSON:", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	})

	port := "8080"
	log.Printf("Server starting on port %s\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Could not listen on port %s: %v", port, err)
	}
}

