package problems

import (
	"errors"
	"github.com/Astra-max/data-str-algo/go/queue/libs"
)

/**
 * QueStack - implements a stack using quesues
 * Que2 - hold data temporarily
 */


 type QueStack struct {
	Que1 *libs.Que
	Que2 *libs.Que
 }

 func NewQueStack() *QueStack {
	que1 := new(libs.Que)
	que2 := new(libs.Que)

	qs := &QueStack{Que1: que1, Que2: que2}
	return qs
 }

 func (qs *QueStack) Push(data interface{}) {
	qs.Que1.Push(data)
 }

 func (qs *QueStack) Peek() (interface{}, error) {
	if qs.Que1.IsEmpty() {
		return nil, errors.New("Data not found")
	}

	queA := qs.Que1
	queB := qs.Que2

	// move data from que a to b
	for queA.Size() > 0 {
		data, found := queA.Pop()

		if found {
			queB.Push(data)
		}
	}
	data, found := queA.Peek()

	// restore queA to avoid corrupting other operations

	for queB.Size() > 0 {
		data,_ := queB.Pop()
		queA.Push(data)
	}
	if !found {
		return nil, errors.New("Data not found")
	}
	return data,nil
 }

 func (qs *QueStack) Pop() (interface{}, error) {
	if qs.Que1.IsEmpty() {
		return nil, errors.New("Data not found")
	}

	queA := qs.Que1
	queB := qs.Que2

	// move data from que a to b
	for queA.Size() > 0 {
		data, found := queA.Pop()

		if found {
			queB.Push(data)
		}
	}
	data, found := queA.Pop()
	// restore queA to avoid corrupting other operations

	for queB.Size() > 0 {
		data,_ := queB.Pop()
		queA.Push(data)
	}
	if !found {
		return nil, errors.New("Data not found")
	}
	return data,nil
 }