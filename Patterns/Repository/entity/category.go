package entity

type Category struct {
	Id   int
	Name string
}

type CategoryRepository interface {
	Get(id int) (Category, error)
	Create(name string) (Category, error)
}

func NewCategory(name string) Category {
	return Category{}
}
