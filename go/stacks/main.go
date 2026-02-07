package main


import (
	"stacks/stk"
	p "stacks/problems"
)

func main() {
	s := stk.NewStack()
	s.Push(20)
	s.Push(30)
	min := p.MinStack(9)
	min.Print()
	s.Print()
}
