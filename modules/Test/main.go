package main

import "fmt"

type Products struct {
	Name string
	Price float64
	Tax float64
}

func CalcTax(x *float64) float64 {
	return *x * 0.1
}

func main() {
	x := Products{"Malboro", 10, 0}
	y := Products{"Beer", 5, 0}
	z := Products{"Watter", 2, 0 }

	x.Tax = CalcTax(&x.Price)
	y.Tax = CalcTax(&y.Price)
	z.Tax = CalcTax(&z.Price)

	fmt.Printf("%v\n%v\n%v\n", x, y, z )
}
