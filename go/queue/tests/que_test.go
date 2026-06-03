package tests


import (
	"github.com/Astra-max/data-str-algo/go/queue/libs"
	"testing"
)

func Test_PeekQue(t *testing.T) {
	que := libs.NewQue()
	que.Push(3)
	que.Push(7)
	val, ok := que.Peek()
	expected := 3

	if ok && val != expected {
		t.Fatalf("Expected %v but got %v!!.", expected, val)
	}
}
