package problems


func IsValid(s string) bool {
	match := map[rune]rune{'(':')','[':']','{':'}'}

	stack := []rune{}

	for _, char := range s {
		switch char {
		case '(','[','{':
			stack = append(stack, char)
			break
		case ')',']','}':
			if len(stack) == 0 || match[stack[len(stack)-1]] != char {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
			break
		}
	}
	return len(stack) == 0
}
