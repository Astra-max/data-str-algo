package tests

import (
	"testing"
	"dp/libs"
)

func Test_DpFibn(t *testing.T) {
	got := libs.DpFibnacci(5)
	expected := 5

	if expected != got {
		t.Fatalf("Expected %v got %v.", expected, got)
	}
}
