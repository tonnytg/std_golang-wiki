//package main
//
//import (
//	"errors"
//	"fmt"
//	"github.com.br/tonnytg/std-golang-wiki/live-go-esquenta-fc/entity"
//)
//
//func main() {
//	fmt.Println("Ola")
//	resultado, err := soma( 10, 10)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	fmt.Println(resultado)
//
//	products := entity.Products{}
//	product := entity.Product{}
//	product.ID = "abc"
//	product.Name = "Fusca"
//
//	product2 := entity.NewProduct()
//	product2.Name = "BMW"
//
//	products.Add(product)
//	products.Add(*product2)
//
//	fmt.Println(products)
//
//}
//
//func soma(a int, b int) (int, error) {
//	if a < 0 {
//		return 0, errors.New("a < 0")
//	}
//	return a + b, nil
//}
package main

import (
	"github.com.br/tonnytg/std-golang-wiki/live-go-esquenta-fc/data"
	"github.com.br/tonnytg/std-golang-wiki/live-go-esquenta-fc/webserver"
)

func main() {
	data.LoadData()
	webserver2 := webserver.NewWebserver()
	webserver2.Serve()
}





