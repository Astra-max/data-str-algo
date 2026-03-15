package tests


import (
	"testing"
	"que/libs"
)

func SQuePop_Test(t *testing.T) {
	que := libs.NewSQue()
	que.Push(9)
	que.Push(3)
	val, ok := que.Pop()
	expected := 9

	if ok && expected != val {
		t.Fatalf("Expected %v but got %v.", expected,val)
	}
}
