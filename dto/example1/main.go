package main

import (
	"dto-example/entity/product"
	"dto-example/pkg/dto"
	"encoding/json"
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "make POST at http://localhost:8080/product\n")
}

func jsonError(msg string) []byte {
	error := struct {
		Message string `json:"message"`
	}{
		msg,
	}
	r, err := json.Marshal(error)
	if err != nil {
		return []byte(err.Error())
	}
	return r
}

func setProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var productDto dto.Product

		err := json.NewDecoder(r.Body).Decode(&productDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		fmt.Fprintf(w, "productDto: %+v\n", productDto)
		fmt.Printf("productDto: %+v\n", productDto)
	}
}

func main() {
	fmt.Println("Start test with DTO")

	p := product.Product{}
	fmt.Println(p)

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/project", setProduct)
	http.ListenAndServe(":8080", nil)


}