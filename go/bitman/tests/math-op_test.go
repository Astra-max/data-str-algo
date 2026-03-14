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
