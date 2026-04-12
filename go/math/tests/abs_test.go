package tests

import (
	"testing"
	"math/libs"
)

func TestAbs(t *testing.T) {
	expected := 5.0
	result := libs.Abs(-5.0)

	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}
}