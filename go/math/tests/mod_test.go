package tests

import (
	"testing"
	"math/libs"
)

func TestMod(t *testing.T) {
	expected := -1
	result := libs.Mod(-10, 3)
	
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}