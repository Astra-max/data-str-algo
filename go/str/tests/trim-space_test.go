package tests

import (
	"testing"
	"str/problems"
)

func TestTrimSpace(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"   Hello, World!   "}, "Hello, World!"},
		{"test2", args{"NoSpaces"}, "NoSpaces"},
		{"test3", args{"    Leading spaces"}, "Leading spaces"},
		{"test4", args{"Trailing spaces    "}, "Trailing spaces"},
		{"test5", args{"    Both sides    "}, "Both sides"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problems.TrimSpace(tt.args.s); got != tt.want {
				t.Errorf("TrimSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}