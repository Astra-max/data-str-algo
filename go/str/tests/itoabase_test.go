package tests

import (
	"testing"
	"str/problems"
)

func Test_Itoabase(t *testing.T) {
	expected := "1010"
	got := problems.Itoabase(10, 2)

	if expected != got {
		t.Fatalf("Expected %v expected but got %v", expected, got)
	}
}