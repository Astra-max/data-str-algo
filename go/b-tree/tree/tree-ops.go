package tree

import (
	"fmt"
	"github.com/Astra-max/data-str-algo/go/queue/libs"
)


func (t *Tree) LeveLOrdering() {
	que := libs.NewQue()
	var curr *Node

	if !t.IsEmpty() {
		que.Push(t.Root)
	}

	for !que.IsEmpty() {
		node, _ := que.Pop()
		curr = node.(*Node)
		fmt.Println("Level value: ", curr.Data)

		if curr.Lchild != nil {
			que.Push(curr.Lchild)
		}
		if curr.Rchild != nil {
			que.Push(curr.Rchild)
		}
	}
}