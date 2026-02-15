package problems


import (
	"stacks/stk"
)

type Min struct {
	stk1 stk.Stack
	stk2 stk.Stack
}

func MinConst() *Min {
	return new(Min)
}

func (m *Min) Push(data interface{}) {
	m.stk1.Push(data)
	top,_ := m.stk2.Peek()

	if m.stk2.IsEmpty() || data.(int) <= top.(int) {
		m.stk2.Push(data)
	}
}

func (m *Min) Pop() {
	val1,_ := m.stk1.Pop()
	top,_ := m.stk2.Peek()
	if val1 == top {
		_,_ = m.stk2.Pop()
	}
}

func (m *Min) Min() interface{} {
	top,_ := m.stk2.Peek()
	return top
}
