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

func Test_Atoibase(t *testing.T) {
	expected := "10"
	got := problems.Atoibase("1010","0123456789", 2)

	if expected != got {
		t.Fatalf("Expected %v but got %v", expected, got)
	}
}