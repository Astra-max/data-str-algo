package problems


func Atoibase(val, baseTo string, baseFrom int) {}

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