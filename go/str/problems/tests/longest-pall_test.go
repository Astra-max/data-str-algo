package tests

import (
	"testing"
	"problems/libs"
)

func Test_LongestPall(t *testing.T) {
	expected := "bab"
	results := libs.LongestPalindrome("babad")

	if expected != results {
		t.Fatalf("Expected %#v got %#v", expected, results)
	}
}