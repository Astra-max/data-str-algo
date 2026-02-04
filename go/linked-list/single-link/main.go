package main


import (
	"single-link/libs"
)

func main() {
	list := libs.NewList()
	list.AddNode(56)
	list.AddNode(12)
	list.AddNode(9)
	list.Print()
}
