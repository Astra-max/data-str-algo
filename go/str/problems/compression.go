package problems


import (
	"strconv"
)

func StrCompression(str string) string {
	results := ""

	occurance := make(map[rune]int)

	for _, char := range str {
		occurance[char]++
	}

	for _, char := range str {
		count, ok := occurance[char]

		if ok {
			val := strconv.Itoa(count)
			results += string(char) + val
			delete(occurance, char)
		}
	}
	return results
}