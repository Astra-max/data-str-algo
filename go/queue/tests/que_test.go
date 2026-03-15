package tests


import (
	"que/libs"
	"testing"
)

func PeekQue_Test(t *testing.T) {
	que := libs.NewQue()
	que.Push(3)
	que.Push(7)
	val, ok := que.Peek()
	expected := 8

	if ok && val != expected {
		t.Fatalf("Expected %v but got %v!!.", expected, val)
	}
}
