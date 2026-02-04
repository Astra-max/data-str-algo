package libs

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type List struct {
	head *Node
}

func NewList() *List {
	return new(List)
}

func (l *List) AddNode(data interface{}) {
	newNode := &Node{data: data, next: nil}

	if l.head == nil {
		l.head = newNode
		return
	}
	newNode.next = l.head
	l.head = newNode
}

func (l *List) Print() {
	if l.head == nil {
		return
	}
	curr := l.head

	for curr != nil {
		fmt.Printf("[%d] ->", curr.data)
		curr = curr.next
	}
	fmt.Println()
}
