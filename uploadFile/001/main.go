package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type FormData struct {
	Name       string                `json:"name"`
	Email      string                `json:"email"`
	Phone      string                `json:"phone"`
	CheckTerms bool                  `json:"checkTerms"`
	CheckNews  bool                  `json:"checkNews"`
	File       *multipart.FileHeader `json:"file"`
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Permitir solicitações de origem cruzada (CORS)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

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

func main() {
	// Rota para upload de arquivo
	http.HandleFunc("/upload", uploadFile)

	// Iniciar o servidor na porta 3005
	log.Fatal(http.ListenAndServe(":3005", nil))
}
