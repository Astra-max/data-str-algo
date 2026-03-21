package tests


import (
	"testing"
	"stacks/stk"
)

func TestPeek(t *testing.T) {
	s := stk.NewStack()
	s.Push(56)
	val, _ := s.Peek()
	expected := 56

	if val != expected {
		t.Fatal("test fail")
	}
}
