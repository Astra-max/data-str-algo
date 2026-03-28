package problems

func TrimSpace(s string) string {
	var start, end int

	runes := []rune(s)

	end  = len(runes)-1
	for start < end && runes[start] == ' ' {
		start++
	}
	for end  >= start && runes[end] == ' ' {
				end--
	}
	return string(runes[start:end+1])
}