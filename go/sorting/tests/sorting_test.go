package tests


import (
	"testing"
	"reflect"
	"sorting/sort"
)

func TestSelection(t *testing.T) {
	arr := []int{2, 11, 3, 6, 9, 1, 7}
	expected := []int{1, 2, 3, 6, 7, 9, 11}

	val := sort.SelectionSort(arr)

	if !reflect.DeepEqual(val, expected) {
		t.Fatalf("Test failed expected %v got %v\n", expected, val)
	}
}

func TestBubble(t *testing.T) {
	arr := []int{2, 11, 3, 6, 9, 1, 7}
	expected := []int{1, 2, 3, 6, 7, 9, 11}

	val := sort.BubbleSort(arr)

	if !reflect.DeepEqual(expected, val) {
		t.Fatalf("Test failed expected %v got %v\n", expected, val)
	}
}


func TestInsertort(t *testing.T) {
	arr := []int{2, 11, 3, 6, 9, 1, 7}
	expected := []int{1, 2, 3, 6, 7, 9, 11}

	val := sort.InsertionSort(arr)

	if !reflect.DeepEqual(expected, val) {
		t.Fatalf("Test failed expected %v got %v\n", expected, val)
	}
}

func TestMergeSort(t *testing.T) {
	arr := []int{2, 11, 3, 6, 9, 1, 7}
	expected := []int{1, 2, 3, 6, 7, 9, 11}

	val := sort.MergeSort(arr, 0, len(arr)-1)

	if !reflect.DeepEqual(expected, val) {
		t.Fatalf("Test failed expected %v got %v\n", expected, val)
	}
}