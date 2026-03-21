package tests

import (
	"testing"
	"que/problems"
)

func Test_QueStackPush(t *testing.T) {
	qs := problems.NewQueStack()

	qs.Push(9)
	qs.Push(2)

	expected := 2

	got, err := qs.Peek()

	if err != nil {
		t.Fatalf("Expected %v got %v", expected, err)
	}

	if expected != got {
		t.Fatalf("Expected %v but got %v", expected, got)
	}
}