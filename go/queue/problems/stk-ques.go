package problems

import (
	"que/libs"
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
	que := new(libs.Que)

	qs := &QueStack{Que1: que, Que2: que}
	return qs
 }

 func (qs *QueStack) Push(data interface{}) {
	qs.Que1.Push(data)
 }

 func (qs *QueStack) Peek() (interface{}, error) {
	return nil,nil
 }

 func (qs *QueStack) Pop() (interface{}, error) {
	return nil,nil
 }