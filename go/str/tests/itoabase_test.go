package tests

import (
	"testing"
	"str/problems"
)

func Test_Itoabase(t *testing.T) {
	expected := "FF"
	got := problems.Itoabase(255, 16)

	if expected != got {
		t.Fatalf("Expected %v expected but got %v", expected, got)
	}
}