package libs

import (
	"math"
)

func ReverseInt(num int) int {
	if num % 10 == num {
		return num
	}

	results := 0

	for num > 0 {
		lastVal := num % 10

		switch {
		case results*10 > math.MaxInt32:
			return 0
		case results * 10 + lastVal > math.MaxInt32:
			return 0
			
		}

		results = results * 10 + lastVal
		num /= 10
	}
	return results
}