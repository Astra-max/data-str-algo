package problems

func TrimSpace(s string) string {
	var start, end int

	runes := []rune(s)

	end  = len(runes)-1

	for i:=0; i<len(runes); i++ {
		if runes[i] != ' ' {
			start = i
			break
		}
	}
	
	for i:=len(runes)-1; i>=0; i-- {
		if runes[i] != ' ' {
			end = i
			break
		}
	}
	
	// for start < end && runes[start] == ' ' {
	// 	start++
	// }
	// for end  >= start && runes[end] == ' ' {
	// 			end--
	// }
	return string(runes[start:end+1])
}