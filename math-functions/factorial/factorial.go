package factorial

import "fmt"

func Factorial(number int) int {
	if number == 0 || number == 1 {
		return 1
	}

	if number < 0 {
		fmt.Println("Invalid number")
		return -1
	}
	return number * Factorial( number - 1 )
}
