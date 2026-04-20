package tree

import (
	"fmt"
	"github.com/Astra-max/data-str-algo/go/queue/libs"
)

type NodeLevel struct {
	node *Node
	level int
}


func (t *Tree) LeveLOrderingBst() {
	que := libs.NewQue()
	var curr *NodeLevel

	if !t.IsEmpty() {
		que.Push(&NodeLevel{t.Root, 0})
	}

	for !que.IsEmpty() {
		node, _ := que.Pop()
		curr = node.(*NodeLevel)
		fmt.Printf("Level value: %d node level %d\n", curr.node.Data, curr.level)

		if curr.node.Lchild != nil {
			que.Push(&NodeLevel{node: curr.node.Lchild, level: curr.level + 1})
		}
		if curr.node.Rchild != nil {
			que.Push(&NodeLevel{node: curr.node.Rchild, level: curr.level + 1})
		}
	}
}

func (t *Tree) TreeHeight(curr *Node) int {
	if curr == nil {
		return 0
	}
	
	lSubTree := t.TreeHeight(curr.Lchild)
	rSubTree := t.TreeHeight(curr.Rchild)

	if lSubTree > rSubTree {
		return lSubTree + 1
	}
	return rSubTree + 1
}
