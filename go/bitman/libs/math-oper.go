package libs

func Multiply(num1, num2 int) int {
	results := 0
	for num2 > 0 {
		if num2 & 1 == 1 {
			results += num1
		}

		num2 <<= 1
		num1 >>= 1 
	}
	return results
}

func Divide(num1, num2 int) int {
	results := 0
	for num2 > 0 {
		if num2 & 1 == 0 {
			results -= num1
		}
		num2 <<= 1
		num1 <<= 1
	}
	return results
}

func Swap(a, b int) (int, int) {
	a = a ^ b
	b = b ^ a
	return a, b
}
