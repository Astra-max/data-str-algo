package libs

import "fmt"

type Node struct {
	data interface{}
	Next *Node
}

type List struct {
	Head *Node
}

func NewList() *List {
	return new(List)
}

func (l *List) AddNode(data interface{}) {
	newNode := &Node{data: data, Next: nil}

	if l.Head == nil {
		l.Head = newNode
		return
	}
	newNode.Next = l.Head
	l.Head = newNode
}

func (l *List) Print() {
	if l.Head == nil {
		return
	}
	curr := l.Head

	for curr != nil {
		fmt.Printf("[%d] ->", curr.data)
		curr = curr.Next
	}
	fmt.Println("nil")
}

func (l *List) isEmpty() bool {
	return l.Head == nil
}

func (l *List) Delete() bool {
	dummy := &Node{data: 0, Next: nil}
	if l.Head == nil {
		return false
	}

	if l.Head.Next == nil {
		if l.Head.data == 0 {
			return false
		}
		l.Head = dummy
		return true
	}

	l.Head = l.Head.Next
	return true
}

func (l *List) Insert(pstn int, data interface{}) bool {
	size := l.Len()
	newNode := &Node{data, nil}
	curr := l.Head

	if pstn > size {
		return false
	} else if pstn == 0 {
		l.AddNode(data)
		return true
	} else if pstn == size {
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = newNode
		return true
	} else {
		for curr != nil {
			next := curr.Next

			if pstn-1 == size {
				curr.Next = newNode
				newNode.Next = next
				return true
			}
			curr = curr.Next
			size--
		}
	}
	return false
}

func (l *List) Len() int {
	curr := l.Head
	count := 0

	for curr != nil {
		count++
		curr = curr.Next
	}
	return count
}
