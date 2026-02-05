package problems


import (
	"single-link/libs"
)


func Reverse(l *libs.List) {
	prev := new(libs.Node)
	curr := l.Head

	if curr == nil {
		return
	}

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
		curr = curr.Next
	}
	prev = l.Head
}
