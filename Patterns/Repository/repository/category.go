package repository

import (
	"fmt"
	"github.com/tonnytg/std-go-repository/entity"
)

type CategoryRepositoryInterface struct {
	Repository entity.CategoryRepository
}

func (c *CategoryRepositoryInterface) Get(id int) (entity.Category, error) {
	return c.Repository.Get(id)
}

func (c *CategoryRepositoryInterface) Create(name string) (entity.Category, error) {
	fmt.Println("Crated")
	return c.Repository.Create(name)
}

func NewCategoryRepositoryInterface(repository entity.CategoryRepository) *CategoryRepositoryInterface {
	return &CategoryRepositoryInterface{Repository: repository}
}
