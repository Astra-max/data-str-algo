package tests

import (
	p "bitman/problems"
	"testing"
)

func Test_ActiveBits(t *testing.T) {
	expected := 3
	got := p.ActiveBits(7)

	if got != expected {
		t.Fatalf("Expected %v got %v\n", expected, got)
	}
}
