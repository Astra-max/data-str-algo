package tests

import (
	"testing"
	"str/problems"
)

func Test_CompressStr(t *testing.T) {
	expected := "a1b5c1"
	got := problems.StrCompression("abbbbbc")

	if expected != got {
		t.Fatalf("Expected %v but got %v", expected, got)
	}
}