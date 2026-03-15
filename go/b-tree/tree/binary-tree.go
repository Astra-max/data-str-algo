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

func CreateTree(arr []int, low, high int) *Node {
	mid := (low + high)/2

	if low > high {
		return nil
	}
	root := &Node{Data: arr[mid]}
	root.Lchild = CreateTree(arr, low, mid-1)
	root.Rchild = CreateTree(arr, mid+1, high)
	return root
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
