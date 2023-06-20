package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

type FormData struct {
	Name       string                `json:"name"`
	Email      string                `json:"email"`a
	Phone      string                `json:"phone"`
	CheckTerms bool                  `json:"checkTerms"`
	CheckNews  bool                  `json:"checkNews"`
	File       *multipart.FileHeader `json:"file"`
}

func uploadFile2(w http.ResponseWriter, r *http.Request) {
	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)

	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
}

func uploadFileOld(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		// Responder com sucesso para as solicitações OPTIONS (pré-voo)
		return
	}

	// Verificar se a solicitação é um método POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido.", http.StatusMethodNotAllowed)
		return
	}

	// Ler o corpo da solicitação
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Falha ao ler o corpo da solicitação.", http.StatusBadRequest)
		return
	}

	// Converter o JSON para uma struct FormData
	var formData FormData
	err = json.Unmarshal(body, &formData)
	if err != nil {
		http.Error(w, "Falha ao converter o JSON para FormData.", http.StatusBadRequest)
		return
	}

	// Abrir o arquivo enviado
	file, err := formData.File.Open()
	if err != nil {
		http.Error(w, "Falha ao abrir o arquivo enviado.", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Criar um arquivo no servidor para salvar o arquivo enviado
	dst, err := os.Create(formData.File.Filename)
	if err != nil {
		http.Error(w, "Falha ao criar o arquivo no servidor.", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copiar o conteúdo do arquivo enviado para o arquivo criado no servidor
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Falha ao salvar o arquivo no servidor.", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Arquivo recebido com sucesso!")
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	AllowCord(&w)
	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)

	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create file
	dst, err := os.Create(handler.Filename)
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	uploadFile(w, r)
}

func AllowCord(w *http.ResponseWriter) {
	// Permitir solicitações de origem cruzada (CORS)
	ww := *w
	ww.Header().Set("Access-Control-Allow-Origin", "*")
	ww.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	ww.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":3005", nil)
}
