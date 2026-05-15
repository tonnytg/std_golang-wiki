package data

import "github.com.br/tonnytg/std-golang-wiki/live-go-esquenta-fc/entity"

var Products entity.Products

func LoadData() {
	product1 := entity.NewProduct()
	product1.Name = "Carrinho"

	product2 := entity.NewProduct()
	product2.Name = "Boneca"
	Products.Add(*product1)
	Products.Add(*product2)
}