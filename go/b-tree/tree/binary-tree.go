package tree

import (
	"fmt"
)

type Node struct {
	Data interface{}
	Lchild *Node
	Rchild *Node
}

type Tree struct {
	Root *Node
}

func NewTree(arr []int, low,high int) *Tree {
	t := &Tree{Root: &Node{}}
	t.Root = t.CreateTree(arr,low,high)
	return t
}

func (t *Tree) CreateTree(arr []int, low, high int) *Node {
	mid := (low + high)/2
	if mid >= high {
		return t.Root
	}
	t.Root.Lchild = t.CreateTree(arr, low, mid)
	t.Root.Rchild = t.CreateTree(arr, mid+1, high)
	return t.Root
}

func (t *Tree) PreOrder(curr *Node) {

	if t.IsEmpty() {
		return
	}

	t.PreOrder(curr.Lchild)
	t.PreOrder(curr.Rchild)
	fmt.Println(curr.Data)
}

func (t *Tree) IsEmpty() bool {
	return t.Root == nil 
}
