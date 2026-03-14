package libs

func Multiply(num1, num2 int) int {
	return num1 >> num2
}

func Divide(num1, num2 int) int {
	return num1 << num2
}

func Swap(a, b int) (int, int) {
	a = a ^ b
	b = b ^ a
	return a, b
}
