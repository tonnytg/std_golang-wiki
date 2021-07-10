package main

import "testing"

func TestCalcTax(t *testing.T) {
	x := Products{"Malboro", 10, 0}
	y := Products{"Beer", 5, 0 }
	z := Products{"Watter", 2, 0}

	x.Tax = CalcTax(&x.Price)
	if x.Tax != 1 {
		t.Errorf("CalcTax(&x.Price) = %v; want = 1", x.Tax)
	}

	y.Tax = CalcTax(&y.Price)
	if y.Tax != 0.5 {
		t.Errorf("CalcTax(&y.Price) = %v; want = 0.5", y.Tax)
	}

	z.Tax = CalcTax(&z.Price)
	if z.Tax != 0.2 {
		t.Errorf("CalcTax(&z.Price) = %v; want = 0.2", z.Tax)
	}
}

