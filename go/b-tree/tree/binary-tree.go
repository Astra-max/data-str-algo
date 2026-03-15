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

func NewTree(arr []int, low,high int) *Node {
	t := CreateTree(arr,low,high)
	return t
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

	if t.IsEmpty() {
		return
	}

	fmt.Println(curr.Data)
	t.PreOrder(curr.Lchild)
	t.PreOrder(curr.Rchild)
}

func Inorder(curr *Node) {
	if curr == nil {
		return
	}
	Inorder(curr.Lchild)
	fmt.Println(curr.Data)
	Inorder(curr.Rchild)
}

func (t *Tree) IsEmpty() bool {
	return t.Root == nil 
}
