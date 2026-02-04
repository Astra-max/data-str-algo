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
	fmt.Println("nil")
}

func (l *List) isEmpty() bool {
	return l.head == nil
}

func (l *List) Delete() bool {
	dummy := &Node{data: 0, nil}
	if l.head == nil {
		return true
	}

	if l.head.next == z

	if l.head.next != nil {
		l.head = l.head.next
		return true
	}
	return false
}

