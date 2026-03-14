package main

import (
	"fmt"
	"b-tree/tree"
)

func main() {
	arr := []int{1,2,3,4,5,6,7,8}

	t := tree.NewTree(arr,0, len(arr))
	//t.PreOrder(t.Root)
	fmt.Printf("%#v\n",t.Root.Lchild.Data)
}
