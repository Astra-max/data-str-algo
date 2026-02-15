package main


import (
	"fmt"
	"stacks/stk"
	p "stacks/problems"
)

func main() {
	s := stk.NewStack()
	s.Push(20)
	s.Push(30)
	min := p.MinConst()
	min.Push(3)
	min.Push(2)
	min.Push(5)
	fmt.Println(min.Min())
	s.Print()
}
