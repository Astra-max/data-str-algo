package tree


type Node struct {
	Data interface{}
}

type Tree struct {
	Lchild *Node
	Rchild *Node
}

