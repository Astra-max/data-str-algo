package main


import (
	"stacks/stk"
)

func main() {
	s := stk.NewStack()
	s.Push(20)
	s.Push(30)
	s.Print()
}
