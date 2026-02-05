package main


import (
	"fmt"
	"single-link/libs"
	"single-link/problems"
)

func main() {
	list := libs.NewList()
	list.AddNode(56)
	list.AddNode(12)
	list.AddNode(9)
	problems.Reverse(list)
	fmt.Println(list.Insert(3, 89))
	list.Print()
}
