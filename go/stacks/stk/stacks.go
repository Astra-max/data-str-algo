package stk

import (
	"fmt"
	"errors"
)


type Node struct {
	data interface{}
	next *Node
}

type Stack struct {
	top *Node
}

func NewStack() *Stack {
	return new(Stack)
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Push(data interface{}) {
	s.top = &Node{data: data, next: s.top}
}

func (s *Stack) Pop() bool {
	if s.IsEmpty() {
		return false
	}
	
	if s.top.next == nil {
		return false
	}
	s.top = s.top.next
	return true
}


func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return s.top.data, nil
}


func (s *Stack) Print() {
	curr := s.top

	for curr != nil {
		fmt.Printf("[ %d  ]-> ", s.top.data)
		curr = curr.next
	}
	fmt.Println("nil")
}
