package tests

import (
	"sorting/sort"
	"reflect"
	"testing"
)


func TestRadixSort(t *testing.T) {
	arr := []int{2, 11, 3, 6, 9, 1, 7}
	expected := []int{1, 2, 3, 6, 7, 9, 11}

	val := sort.RadixSort(arr)

	if !reflect.DeepEqual(expected, val) {
		t.Fatalf("Test failed expected %v got %v\n", expected, val)
	}
}