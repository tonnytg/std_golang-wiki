package roman_to_int

func romanToInt(s string) int {
	R := map[int32]int{73: 1, 86: 5, 88: 10, 76: 50, 67: 100, 68: 500, 77: 1000}

	var sum, n int

	for i := len(s) - 1; i >= 0; i-- {
		if n <= R[int32(s[i])] {
			sum = sum + R[int32(s[i])]
		} else {
			sum = sum - R[int32(s[i])]
		}
		n = R[int32(s[i])]
	}
	return sum
}
