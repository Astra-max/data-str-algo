package tests

import (
	p "bitman/problems"
	"testing"
)

func Test_ActiveBits(t *testing.T) {
	err := LoadTests("../tests.json")

	if err != nil {
		t.Fatal(err)
		return
	}

	for _, tt := range Cases.Cases {
		t.Run(Cases.Test, func(t *testing.T) {
			results := p.ActiveBits(tt.Value)
			if results != tt.Expected {
				t.Fatalf("Exepected %v got %v.", tt.Expected, results)
			}
		})
	}
}
