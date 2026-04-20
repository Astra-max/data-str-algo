package main

import (
	"b-tree/tree"
)

func main() {
	arr := []int{1,2,3,4,5,6,7,8}

	t := tree.NewTree(arr,0, len(arr)-1)
	t.PostOrder(t.Root)
	t.LeveLOrdering()
}
