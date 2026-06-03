package problems


import (
	"single-link/libs"
)


func Reverse(l *libs.List) {
	var prev *libs.Node = nil
	curr := l.Head

	if curr == nil {
		return
	}

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head = prev
}
