package libs

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

type Que struct {
	Head *Node
}

func NewQue() *Que {
	return new(Que)
}

func (q *Que) Push(data interface{}) {
	newNode := &Node{Data:data}

	if q.IsEmpty() {
		q.Head = newNode
		return
	}
	curr := q.Head

	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
}

func (q *Que) Pop() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	}
	firstVal := q.Head.Data
	q.Head = q.Head.Next
	return firstVal, true
}

func (q *Que) Peek() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	}
	return q.Head.Data, true
}

func (q *Que) IsEmpty() bool {
	return q.Head == nil
}

func (q *Que) Print() {
	if q.IsEmpty() {
		fmt.Println("Empty queue!!")
		return
	}

	curr := q.Head

	for curr != nil {
		fmt.Println(curr.Data)
		curr = curr.Next
	}
}
