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
	dummy := &Node{data: 0, next: nil}
	if l.head == nil {
		return false
	}

	if l.head.next == nil {
		if l.head.data == 0 {
			return false
		}
		l.head = dummy
		return true
	}

	l.head = l.head.next
	return true
}

func (l *List) Insert(pstn int, data interface{}) bool {
	size := l.Len()
	newNode := &Node{data, nil}
	curr := l.head

	if pstn > size {
		return false
	} else if pstn == 0 {
		l.AddNode(data)
		return true
	} else if pstn == size {
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = newNode
		return true
	} else {
		for curr != nil {
			next := curr.next

			if pstn-1 == size {
				curr.next = newNode
				newNode.next = next
				return true
			}
			curr = curr.next
			size--
		}
	}
	return false
}

func (l *List) Len() int {
	curr := l.head
	count := 0

	for curr != nil {
		count++
		curr = curr.next
	}
	return count
}
