package tests

import (
	"bitman/libs"
	"testing"
)

func Test_MathOpt(t *testing.T) {
	for _,tt := range MathCases {
		t.Run(tt.TestName, func(t *testing.T){
			results := libs.Multiply(tt.Value1, tt.Value2)
			if results != tt.Expected {
				t.Fatalf("Expected %v but got %v", tt.Expected, results)
			}
		})
	}
}

func Test_AddNum(t *testing.T) {
	results := libs.AddNum(5, 9)
	expected := 14

	if results != expected {
		t.Fatalf("Expected %v got %v.", expected, results)
	}
}
