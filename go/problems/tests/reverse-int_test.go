package tests

import (
	"testing"
	"problems/libs"
)

func Test_ReverseInt(t *testing.T) {
	results := libs.ReverseInt(95)
	expected := 59

	if results != expected {
		t.Fatalf("Expected %v got %v", expected, results)
	}
}