package problems

import (
	"stacks/stk"
)

func ValidParenthesis(s string) bool {
	perfectMatch := map[rune]rune{
		')':'(',
		']':'[',
		'}':'{',
	}

	stack := stk.NewStack()

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack.Push(char)
			break
		case ']', '}', ')':
			pair, ok := stack.Pop()

			perfectPair, exist := perfectMatch[char]

			if ok && exist && pair == perfectPair {
				_,_ = stack.Pop()
			}
		}
	}
	return stack.IsEmpty()
}