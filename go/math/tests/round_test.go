package tests

import (
	"testing"
	"math/libs"
)

func TestRound(t *testing.T) {
	expected := 3.0
	result := libs.Round(2.5)
	
	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}