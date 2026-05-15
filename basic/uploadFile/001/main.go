package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/upload", handleFileUpload)
	log.Fatal(http.ListenAndServe(":3005", nil))
}

func handleFileUpload(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Invalid request method")
		return
	}

	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // Set the maximum file size (in this case, 10MB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error parsing the form: %v", err)
		return
	}

	log.Printf("Request Method: %+v\n", r.Method)
	log.Printf("Request Headers: %+v\n", r.Header)
	log.Printf("Request Body: %+v\n", r.Body)
	log.Printf("Request GetBody: %+v\n", r.GetBody)
	log.Printf("Request Form: %+v\n", r.Form)
	log.Printf("Request PostForm: %+v\n", r.PostForm)
	log.Printf("Request MultipartForm: %+v\n", r.MultipartForm)
	log.Printf("Request RemoteAddr: %+v\n", r.RemoteAddr)
	log.Printf("Request RequestURI: %+v\n", r.RequestURI)
	log.Printf("Request URL: %+v\n", r.URL)
	log.Printf("Request Port: %+v\n", r.URL.Port())
	log.Printf("Request Scheme: %+v\n", r.URL.Scheme)
	log.Printf("Address: %+v\n", r.URL.String())
	log.Printf("Request Path: %+v\n", r.URL.Path)
	log.Printf("Request Query: %+v\n", r.URL.Query())
	log.Printf("Request Hostname: %+v\n", r.URL.Hostname())
	log.Printf("Request Address: %+v\n", r.URL.Host)
	log.Printf("RemoteAddr: %+v\n", r.RemoteAddr)

	// Get the file from the request
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error retrieving the file from the request: %s\n", err)
		log.Printf("Request: %+v\n", r)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error retrieving the file: %v", err)
		return
	}
	defer file.Close()

	// Create a new file on the server to store the uploaded file
	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error creating the file: %v", err)
		return
	}
	defer f.Close()

	// Copy the contents of the uploaded file to the newly created file on the server
	_, err = io.Copy(f, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error copying file contents: %v", err)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully!")
}
