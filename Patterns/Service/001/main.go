package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Aluno struct {
	ID   string
	Nome string
}

type AlunoRepository interface {
	GetByID(id string) (*Aluno, error)
	Save(aluno *Aluno) error
}

type MemoryAlunoRepository struct {
	alunos []*Aluno
}

func (r *MemoryAlunoRepository) GetByID(id string) (*Aluno, error) {
	for _, aluno := range r.alunos {
		if aluno.ID == id {
			return aluno, nil
		}
	}
	return nil, fmt.Errorf("aluno não encontrado")
}

func (r *MemoryAlunoRepository) Save(aluno *Aluno) error {
	r.alunos = append(r.alunos, aluno)
	return nil
}

type AlunoService struct {
	repo AlunoRepository
}

func (s *AlunoService) GetAlunoByID(id string) (*Aluno, error) {
	aluno, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return aluno, nil
}

func (s *AlunoService) CreateAluno(aluno *Aluno) error {
	return s.repo.Save(aluno)
}

func main() {
	repo := &MemoryAlunoRepository{}
	service := &AlunoService{repo: repo}

	// Exemplo de uso
	aluno := &Aluno{ID: "1", Nome: "João"}
	err := service.CreateAluno(aluno)
	if err != nil {
		log.Fatal(err)
	}

	alunoRetornado, err := service.GetAlunoByID("1")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Aluno: %+v\n", alunoRetornado)

	// Configuração das rotas HTTP
	http.HandleFunc("/alunos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			id := r.URL.Query().Get("id")
			if id == "" {
				http.Error(w, "ID não fornecido", http.StatusBadRequest)
				return
			}

			aluno, err := service.GetAlunoByID(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			json.NewEncoder(w).Encode(aluno)
		case http.MethodPost:
			var novoAluno Aluno
			err := json.NewDecoder(r.Body).Decode(&novoAluno)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			err = service.CreateAluno(&novoAluno)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
		default:
			http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		}
	})

	// Inicia o servidor HTTP
	log.Fatal(http.ListenAndServe(":8080", nil))
}
