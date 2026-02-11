package main

import (
	"fmt"
	"sorting/sort"
)

func main() {
	arr := []int{2, 11, 3, 6, 9, 1, 7}
	val := sort.MergeSort(arr, 0, len(arr)-1)
	fmt.Println(val)
}
