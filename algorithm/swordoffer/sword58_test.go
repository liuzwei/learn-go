package swordoffer

import "testing"

func TestReverseLeftWords(t *testing.T) {
	tests := []struct {
		input  string
		num    int
		output string
	}{
		{"abc", 1, "bca"},
		{"abc", 2, "cab"},
	}

	for _, tt := range tests {
		result := ReverseLeftWords(tt.input, tt.num)
		if result != tt.output {
			t.Errorf("ReverseLeftWords error ! got=%v, want=%v", result, tt.output)
		}
	}
}
