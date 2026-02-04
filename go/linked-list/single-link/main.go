package main


import (
	"fmt"
	"single-link/libs"
)

func main() {
	list := libs.NewList()
	list.AddNode(56)
	list.AddNode(12)
	list.AddNode(9)
	//fmt.Println(list.Delete())
	fmt.Println(list.Insert(3, 89))
	list.Print()
}
