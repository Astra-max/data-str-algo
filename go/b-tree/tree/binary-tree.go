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
	tree := new(Tree)
	t := CreateTree(arr,low,high)
	tree.Root = t
	return tree
}

func CreateTree(arr []int, start, size int) *Node {
	curr := &Node{Data: arr[start]}
	left := 2 * start + 1
	right := 2 * start + 2

	if start >= size {
		return nil
	}
	
	curr.Lchild = CreateTree(arr, left, size)
	curr.Rchild = CreateTree(arr, right, size)
	return curr
}

func (t *Tree) PreOrder(curr *Node) {

	if t.IsEmpty() || curr == nil {
		return
	}

	fmt.Println(curr.Data)
	t.PreOrder(curr.Lchild)
	t.PreOrder(curr.Rchild)
}

func (t *Tree) Inorder(curr *Node) {
	if curr == nil {
		return
	}
	t.Inorder(curr.Lchild)
	fmt.Println(curr.Data)
	t.Inorder(curr.Rchild)
}

func (t *Tree) IsEmpty() bool {
	return t.Root == nil 
}
