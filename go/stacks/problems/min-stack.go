package problems

import (
	"stacks/stk"
) 


func MinStack(data interface{}) *stk.Stack {
	stk1 := new(stk.Stack)
	stk2 := new(stk.Stack)

	stk1.Push(1)
	stk2.Push(1)
	stk1.Push(10)
	stk2.Push(3)

	val, cond := stk1.Pop()
	val2, _:= stk2.Pop()

	if cond && data.(int) >= val.(int) {
		stk2.Push(val)
		stk1.Push(data)
	}
	if cond && data.(int) <= val.(int) {
		if val2.(int) > data.(int) {
			stk2.Push(data)
			stk2.Push(val2)
		} else {
			stk2.Push(val2)
		        stk2.Push(data)
	      }
	}
	return stk2
}
