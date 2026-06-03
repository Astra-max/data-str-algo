package problems

func ReversePolish(tokens []string) int {
	if len(tokens) == 0  {
		return 0
	}

	stck := []int{}
	for _, token := range tokens {
		switch token {
		case "+":
			a, b := stck[len(stck)-1], stck[len(stck)-2]
			stck = stck[:len(stck)-2]
			stck = append(stck, a+b)
		case "-":
			a, b := stck[len(stck)-1], stck[len(stck)-2]
			stck = stck[:len(stck)-2]
			stck = append(stck, b-a)
		case "*":
			a, b := stck[len(stck)-1], stck[len(stck)-2]
			stck = stck[:len(stck)-2]
			stck = append(stck, a*b)
		case "/":
			a, b := stck[len(stck)-1], stck[len(stck)-2]
			stck = stck[:len(stck)-2]
			stck = append(stck, b/a)
		default:
			num := 0
			for i, c := range token {
				if c == '-' && i == 0 {
					continue
				}
				num *= 10
				num += int(c - '0')
			}
			if token[0] == '-' {
				num *= -1
			}
			stck = append(stck, num)
		}
	}
	return stck[0]
}