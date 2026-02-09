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


func (t *Tree) PreOrder(curr *Node) {
	if t.IsEmpty() {
		return
	}
	t.PreOrder(curr.Lchild)
	t.PreOrder(curr.Lchild)
	fmt.Println(curr.Data)
}

func (t *Tree) IsEmpty() bool {
	return t.Root == nil 
}