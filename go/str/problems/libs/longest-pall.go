package libs

import "fmt"


func LongestPalindrome(s string) string {
    maxPallindrome := make(map[int]string)
    startIdx := 0
    maxVal := 0

    for i,_ := range s {
        if isPallindrome(s[startIdx:i]) {
			fmt.Println(s[startIdx:i])
            maxVal = max(maxVal, len(s[startIdx:i]))
            maxPallindrome[maxVal] = s[startIdx:i]
        }
    }
    return maxPallindrome[maxVal]
}

func isPallindrome(s string) bool {
    for i,j:=0,len(s)-1; i<j; i,j = i + 1, j - 1 {
        if s[i] != s[j] {
            return false
        }
    }
    return true
}

func max(i, j int) int {
    if i > j {
        return i
    }
    return j
}