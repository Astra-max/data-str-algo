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

func AddNum(n1, n2 int) int {
	for n2 != 0 {
		add := n1 ^ n2
		carry := (n1 & n2) << 1

		n1 = add
		n2 = carry
	}
	return n1
}

func Substract(n1, n2 int) int {
	for n2 != 0 {
		sub := n1 ^ n2
		borrow := (^n1 & n2)<<1
		n1 = sub
		n2 =  borrow
	}
	return n1
}

func Swap(a, b int) (int, int) {
	a = a ^ b
	b = b ^ a
	return a, b
}
