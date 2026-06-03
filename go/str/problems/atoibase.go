package problems


func Atoibase(val, baseTo string, baseFrom int) string {
	digit := 0

	bases := "0123456789ABCDEF"

	for _, char := range val {
		num := ConvertToNum(char)

		if num >= baseFrom || num < 0 {
			return ""
		}
		digit = digit * baseFrom + num
	}

	baseToLen := len(baseTo)

	results := ""

	for digit > 0 {
		remainder := digit % baseToLen

		results = string(bases[remainder]) + results

		digit /= baseToLen
	}
	return results
}

func ConvertToNum(char rune) int {
	if char >= '0' && char <= '9' {
		return int(char - '0')
	} else if char >= 'a' && char <= 'z' {
		return  int((char - 'a') + 10)
	} else if char >= 'A' && char <= 'Z' {
		return int((char - 'A') + 10)
	}
	return -1
}

func valToChar(val int) byte {
	if val >= 0 && val <= 9 {
		return byte('0' + val)
	} else if val >= 10 && val <= 35 {
		return byte('A' + val - 10)
	}
	return '0'
}