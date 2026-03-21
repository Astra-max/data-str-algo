package tests

import (
	"testing"
	"stacks/problems"
)

func Test_ValidParenthesis(t *testing.T) {
	got := problems.ValidParenthesis("([])")
	expected := true

	if expected != got {
		t.Fatalf("Expected %v got %v", expected, got)
	}
}