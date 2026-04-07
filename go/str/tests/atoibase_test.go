package tests

import (
	"testing"
	"str/problems"
)

func Test_ContvNum(t *testing.T) {
	expected := 1

	got := problems.ConvertToNum('1')

	if expected != got {
		t.Fatalf("Expected %v but got %v", expected, got)
	}
}